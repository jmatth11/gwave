package format

import (
	"bytes"
	"encoding/binary"
)

// Int16ToBytes converts int16 to 2 byte array
func Int16ToBytes(n int16, order binary.ByteOrder) (b [2]byte) {
	buf := new(bytes.Buffer)
	binary.Write(buf, order, n)
	copy(b[:], buf.Bytes())
	return
}

// BytesToInt16 converts 2 byte array to int16
func BytesToInt16(b [2]byte, order binary.ByteOrder) (n int16) {
	buf := bytes.NewReader(b[:])
	binary.Read(buf, order, &n)
	return
}

// Int32ToBytes converts int32 to 4 byte array
func Int32ToBytes(n int32, order binary.ByteOrder) (b [4]byte) {
	buf := new(bytes.Buffer)
	binary.Write(buf, order, n)
	copy(b[:], buf.Bytes())
	return
}

// BytesToInt32 converts 4 byte array into int32
func BytesToInt32(b [4]byte, order binary.ByteOrder) (n int32) {
	buf := bytes.NewReader(b[:])
	binary.Read(buf, order, &n)
	return
}
