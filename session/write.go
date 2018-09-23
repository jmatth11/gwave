package session

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/iCurlmyster/wave/format"
	"github.com/iCurlmyster/wave/notes"
)

// addNote pushes the note data for the note length onto the buffer
func addNote(i int, n *notes.Note, wWriter format.WaveWriter) (int, error) {
	if n.Length < 1 {
		return 0, errors.New("length of note is too small. must be greater than or equal to 1")
	}

	// I haven't tested things that aren't just a whole second duration yet, to see if this will screw up
	duration := int(time.Duration(wWriter.FileHeader().BytesPerSecond) * (n.Length / time.Second))
	err := writeNote(i, duration, n, wWriter)
	return duration, err
}

// TODO Did addNoteParallel because working on small scale. Maybe should rework to handle groups of notes instead of individual notes for bigger sound samples

// addNoteParallel Adds notes to file in parallel.
// Accepts a WaitGroup to increment and call Done when writing this note is complete
func addNoteParallel(i int, n *notes.Note, wg *sync.WaitGroup, wWriter format.WaveWriter) int {
	if n.Length < 1 {
		return 0 //, errors.New("length of note is too small. must be greater than or equal to 1")
	}

	// TODO I haven't tested things that aren't just a whole second duration yet, to see if this will screw up
	duration := int(time.Duration(wWriter.FileHeader().BytesPerSecond) * (n.Length / time.Second))
	wg.Add(1)
	go func() {
		err := writeNote(i, duration, n, wWriter)
		// TODO maybe i need to except a channel to send the error message back to
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		wg.Done()
	}()
	return duration
}

func handleNoteBySample(index int, data []byte, wWriter format.WaveWriter) (int, error) {
	n, err := wWriter.WriteAt(data, int64(index))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func writeNote(i, d int, n *notes.Note, wWriter format.WaveWriter) error {
	bc := wWriter.FileHeader().GetByteCount()
	nc := int(wWriter.FileHeader().NumChannels)
	jumpc := bc * nc
	// TODO maybe find a better way to handle phase. notes still have jumping sound between them
	phase := 0
	for j := 0; j < d; j += jumpc {
		val := n.ToData(phase + i)
		data := convertToData(val, wWriter.FileHeader().BitsPerSample, wWriter.FileHeader().FileByteOrder())
		dataLen := 0
		for index := 0; index < nc; index++ {
			step, err := handleNoteBySample(i+j+dataLen, data, wWriter)
			if err != nil {
				return err
			}
			dataLen += step
		}
		phase++
	}
	return nil
}

func convertToData(d float64, bitsPerSample int16, bo binary.ByteOrder) []byte {
	buf := bytes.NewBuffer([]byte{})
	switch bitsPerSample {
	case 8:
		{
			// correct range offset with lower signed value
			binary.Write(buf, bo, uint8(d+128))
		}
	case 16:
		{
			binary.Write(buf, bo, int16(d))
		}
	default:
		{
			binary.Write(buf, bo, int32(d))
		}
	}
	return buf.Bytes()
}
