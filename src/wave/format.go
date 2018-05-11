package wave

// currently not supporting Non-PCM and Extensible fmt types

// Header is a simple layout of the format for a Wave file
type Header struct {
	// "RIFF"
	ByteType [4]byte
	// File Size
	Size int32
	// "WAVE"
	HeaderType [4]byte
	// "fmt "
	FmtMarker [4]byte
	// Format size
	FmtSize int32
	// 1 is PCM type
	FmtType int16
	// number of channels
	NumChannels int16
	// Number of samples per second (Hz) 44100 is CD quality
	SampleRate int32
	// Sample Rate * BitsPerSample * NumChannels / 8. Constant Bit Rate (CBR)
	BytesPerSecond int32
	// NumChannels * BitsPerSample / 8. Number of bytes for one sample including all channels
	BlockAlign int16
	// 16 would be two 8 bytes to sample from. usually 8, 16, or 32
	BitsPerSample int16
	// "data"
	DataMarker [4]byte
	// Size of data
	DataSize int32
}

// ByteOrder type for comparison in converting bytes
type ByteOrder int8

const (
	// PcmType is to represent the PCM type
	PcmType = 1
	// CdSampleRate is the sample rate for a wave file that a CD would have
	CdSampleRate = 44100
	// RiffByteOrder value to compare when converting bytes
	RiffByteOrder = ByteOrder(1)
	// RifxByteOrder value to compare when converting bytes
	RifxByteOrder = ByteOrder(2)
)

var (
	// Wav represents WAVE type
	Wav = [4]byte{'W', 'A', 'V', 'E'}
	// FmtMarker represents the fmt marker in the header
	FmtMarker = [4]byte{'f', 'm', 't', ' '}
	// Riff represents the Byte order type for the file in the header
	Riff = [4]byte{'R', 'I', 'F', 'F'}
	// DataMarker represents the DATA marker in the header.
	DataMarker = [4]byte{'d', 'a', 't', 'a'}
)

// FileByteOrder returns the byte order the file is
func (h Header) FileByteOrder() ByteOrder {
	if h.ByteType[3] == 'F' {
		return RiffByteOrder
	}
	return RifxByteOrder
}
