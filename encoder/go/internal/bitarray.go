package internal

import (
	"errors"
	"strconv"
	"strings"
)

type BitArray struct {
	size   uint
	values []uint8
}

// PutInt64 places the passed int64 value to the stated position in the bit array
func (a BitArray) PutInt64(pos uint, len uint, value int64) {
	if len > 64 {
		panic(errors.New("length and data type not matching"))
	}
	// Create bitmask for packing the value
	mask := int64(2 ^ len - 1) // e.g. length of 4: 2^4-1 = 15d = 00001111b
	// Pack value
	packedValue := value & mask
	// Insert the value into the bit array
	posInByte := pos % 8
	index := pos / 8

}

// ToByteArray converts the BitArray to a ByteArray and returns it
func (a BitArray) ToByteArray() []byte {
	byteArray := make([]byte, a.size)

	return byteArray
}

// ToString returns a string representation of the bit array
func (a BitArray) ToString() string {
	byteStrings := make([]string, len(a.values))
	for i, value := range a.values {
		byteStrings[i] = strconv.FormatUint(uint64(value), 2)
	}
	return strings.Join(byteStrings, " ")
}

// NewBitArray creates a new BitArray instance
func NewBitArray(size uint) BitArray {
	return BitArray{
		size:   size,
		values: make([]uint8, (size+7)/8),
	}
}
