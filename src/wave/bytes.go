package wave

// Int16ToBytes converts int16 to 2 byte array
func Int16ToBytes(n int16, order ByteOrder) [2]byte {
	if (order & RiffByteOrder) == RiffByteOrder {
		return [2]byte{byte(n), byte(n >> 8)}
	}
	return [2]byte{byte(n >> 8), byte(n)}
}

// BytesToInt16 converts 2 byte array to int16
func BytesToInt16(b [2]byte, order ByteOrder) int16 {
	if (order & RiffByteOrder) == RiffByteOrder {
		return int16(b[0]) | int16(b[1])>>8
	}
	return int16(b[0])<<8 | int16(b[1])
}

// Int32ToBytes converts int32 to 4 byte array
func Int32ToBytes(n int32, order ByteOrder) [4]byte {
	if (order & RiffByteOrder) == RiffByteOrder {
		return [4]byte{byte(n), byte(n >> 8), byte(n >> 16), byte(n >> 24)}
	}
	return [4]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)}
}

// BytesToInt32 converts 4 byte array into int32
func BytesToInt32(b [4]byte, order ByteOrder) int32 {
	if (order & RiffByteOrder) == RiffByteOrder {
		return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
	}
	return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
}
