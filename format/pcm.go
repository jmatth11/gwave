package format

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/iCurlmyster/wave/notes"
)

// PCM represents a PCM wave file
type PCM struct {
	Header
	Data []byte
}

// DefaultCDPCM creates a generic 16 bit Riff CD quality wave file
func DefaultCDPCM() *PCM {
	var numChannels int16 = 2
	var bitsPerSample int16 = 16
	return &PCM{
		Header: Header{
			ByteType:       Riff,
			HeaderType:     Wav,
			FmtMarker:      FmtMarker,
			FmtSize:        16,
			FmtType:        PcmType,
			NumChannels:    numChannels,
			SampleRate:     CdSampleRate,
			BytesPerSecond: (CdSampleRate * int32(bitsPerSample) * int32(numChannels)) / 8,
			BlockAlign:     (numChannels * bitsPerSample) / 8,
			BitsPerSample:  bitsPerSample,
			DataMarker:     DataMarker,
		},
		Data: make([]byte, 0),
	}
}

// ReadPCM reads in a WAVE file with PCM header type
func ReadPCM(fileName string) (*PCM, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	header, offset, err := grabHeader(f)
	data := make([]byte, header.DataSize)
	offset, err = readOffset(f, data, offset)
	if err != nil {
		return nil, err
	}
	return &PCM{
		Data:   data,
		Header: header,
	}, nil
}

func grabHeader(f *os.File) (Header, int64, error) {
	header := Header{}
	b := make([]byte, 4)
	var byte4 [4]byte
	var offset int64
	// byte order
	offset, err := readOffset(f, b, offset)
	if err != nil {
		return header, 0, err
	}
	copy(header.ByteType[:], b)
	// grab byte order of file
	order := header.FileByteOrder()
	// file size
	offset, err = readOffset(f, b, offset)
	copy(byte4[:], b)
	header.Size = BytesToInt32(byte4, order)
	b = make([]byte, 12)
	offset, err = readOffset(f, b, offset)
	if err != nil {
		return header, 0, err
	}
	// Header Type
	copy(header.HeaderType[:], b[:4])
	// Format marker
	copy(header.FmtMarker[:], b[4:8])
	// Format Size
	copy(byte4[:], b[8:12])
	header.FmtSize = BytesToInt32(byte4, order)
	var byte2 [2]byte
	b = make([]byte, header.FmtSize)
	offset, err = readOffset(f, b, offset)
	if err != nil {
		return header, 0, err
	}
	// Format Type
	copy(byte2[:], b[:2])
	header.FmtType = BytesToInt16(byte2, order)
	// Number of Channels
	copy(byte2[:], b[2:4])
	header.NumChannels = BytesToInt16(byte2, order)
	// Hz
	copy(byte4[:], b[4:8])
	header.SampleRate = BytesToInt32(byte4, order)
	// Bytes per second
	copy(byte4[:], b[8:12])
	header.BytesPerSecond = BytesToInt32(byte4, order)
	// Block Align
	copy(byte2[:], b[12:14])
	header.BlockAlign = BytesToInt16(byte2, order)
	// Bits per sample
	copy(byte2[:], b[14:16])
	header.BitsPerSample = BytesToInt16(byte2, order)
	b = make([]byte, 8)
	offset, err = readOffset(f, b, offset)
	if err != nil {
		return header, 0, err
	}
	// Data Marker
	copy(header.DataMarker[:], b[:4])
	// Data Size
	copy(byte4[:], b[4:8])
	header.DataSize = BytesToInt32(byte4, order)
	return header, offset, nil
}

func readOffset(f *os.File, b []byte, offset int64) (int64, error) {
	n, err := f.ReadAt(b, offset)
	return (offset + int64(n)), err
}

// WriteToFile writes out PCM data to file
func (pcm *PCM) WriteToFile(fileName string) error {
	return WritePCM(pcm, fileName)
}

// WritePCM takes a PCM object and writes out contents to file specified by fileName
func WritePCM(pcm *PCM, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println("opened filed")
	header := pcm.Header
	order := header.FileByteOrder()
	var offset int64
	byte4 := make([]byte, 4)
	byte2 := make([]byte, 2)
	// header type
	bHead := bytes.NewBuffer(header.HeaderType[:])
	// fmt marker
	bHead.Write(header.FmtMarker[:])
	// format size.
	write4Byte(bHead, byte4, Int32ToBytes(header.FmtSize, order))
	// pcm type
	write2Byte(bHead, byte2, Int16ToBytes(header.FmtType, order))
	// num channels
	write2Byte(bHead, byte2, Int16ToBytes(header.NumChannels, order))
	// sample rate
	write4Byte(bHead, byte4, Int32ToBytes(header.SampleRate, order))
	// bytes per second
	write4Byte(bHead, byte4, Int32ToBytes(header.BytesPerSecond, order))
	// block align
	write2Byte(bHead, byte2, Int16ToBytes(header.BlockAlign, order))
	// bits per sample
	write2Byte(bHead, byte2, Int16ToBytes(header.BitsPerSample, order))
	// data marker
	bHead.Write(header.DataMarker[:])
	// data size
	header.DataSize = int32(len(pcm.Data))
	write4Byte(bHead, byte4, Int32ToBytes(header.DataSize, order))
	// get file size
	fileSize := int32(bHead.Len()) + header.DataSize
	beginData := bytes.NewBuffer(header.ByteType[:])
	write4Byte(beginData, byte4, Int32ToBytes(fileSize, order))
	// byte type
	offset, err = writeOffset(f, beginData.Bytes(), offset)
	if err != nil {
		return err
	}
	fmt.Println("first offset", offset)
	// rest of header
	offset, err = writeOffset(f, bHead.Bytes(), offset)
	if err != nil {
		return err
	}
	tmpD := bHead.Bytes()
	fmt.Println("bhead bytes end", tmpD[len(tmpD)-11:])
	fmt.Println("10 bytes of data", pcm.Data[:20])
	fmt.Println("second offset", offset)
	fmt.Println("finished writing header")
	// data
	offset, err = writeOffset(f, pcm.Data, offset)
	fmt.Println("Last offset", offset)
	return err
}

func writeOffset(f *os.File, b []byte, offset int64) (int64, error) {
	fmt.Println()
	n, err := f.WriteAt(b, offset)
	return (offset + int64(n)), err
}

func write4Byte(b *bytes.Buffer, placeholder []byte, data [4]byte) (int, error) {
	copy(placeholder, data[:])
	return b.Write(placeholder)
}

func write2Byte(b *bytes.Buffer, placeholder []byte, data [2]byte) (int, error) {
	copy(placeholder, data[:])
	return b.Write(placeholder)
}

// AddNote pushes the note data for the note length onto the buffer
func (pcm *PCM) AddNote(i int, n *notes.Note) (int, error) {
	if n.Length < 1 {
		return 0, errors.New("length of note is too small. must be greater than or equal to 1")
	}
	duration := int(time.Duration(pcm.Header.BytesPerSecond) * (n.Length / time.Second))
	//fmt.Println("duration", duration)
	pcm.writeNote(i, duration, n)
	return duration, nil
}

func (pcm *PCM) AddNoteParallel(i int, n *notes.Note, wg *sync.WaitGroup) (int, error) {
	if n.Length < 1 {
		return 0, errors.New("length of note is too small. must be greater than or equal to 1")
	}
	duration := int(time.Duration(pcm.Header.BytesPerSecond) * (n.Length / time.Second))
	fmt.Println("duration", duration)
	wg.Add(1)
	go func() {
		pcm.writeNote(i, duration, n)
		wg.Done()
	}()
	return duration, nil
}

func (pcm *PCM) handleNoteBySample(index int, data byte) int {
	nc := int(pcm.Header.NumChannels)
	for c := 0; c < nc; c++ {
		pcm.Data[index+c] = data
	}
	return nc
}

func (pcm *PCM) writeNote(i, d int, n *notes.Note) {
	bc := pcm.GetByteCount()
	nc := int(pcm.Header.NumChannels)
	jumpc := bc * nc
	j := 0
	phase := 0
	for ; j < d; j += jumpc {
		val := n.ToData(pcm.Header.BytesPerSecond, phase+i)
		data := pcm.convertToData(val)
		channels := 0
		for index := 0; index < len(data); index++ {
			channels += pcm.handleNoteBySample(i+j+channels, data[index])
		}
		phase++
	}
	fmt.Println("j val", j, "jumpc:", jumpc, "nc", nc, "i+j", i+j)
}

func (pcm *PCM) convertToData(d float64) []byte {
	buf := bytes.NewBuffer([]byte{})
	switch pcm.Header.BitsPerSample {
	case 8:
		{
			// correct range offset with lower signed value
			binary.Write(buf, pcm.Header.FileByteOrder(), uint8(d+128))
		}
	case 16:
		{
			binary.Write(buf, pcm.Header.FileByteOrder(), int16(d))
		}
	default:
		{
			binary.Write(buf, pcm.Header.FileByteOrder(), int32(d))
		}
	}
	return buf.Bytes()
}

func (pcm *PCM) GetByteCount() int {
	switch pcm.Header.BitsPerSample {
	case 8:
		{
			return 1
		}
	case 16:
		{
			return 2
		}
	default:
		{
			return 4
		}
	}
}

// AllocateDataSize sets up the PCM Data field to the size given.
func (pcm *PCM) AllocateDataSize(size int32) {
	pcm.DataSize = size
	pcm.Data = make([]byte, size)
}
