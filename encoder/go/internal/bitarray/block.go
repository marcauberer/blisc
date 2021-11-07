package bitarray

import (
	"fmt"
	"unsafe"
)

type block uint8

const s = uint64(unsafe.Sizeof(block(0)) * 8)

func (b block) setBit(position uint64) {
	b = b | block(1<<position)
}

func (b block) clearBit(position uint64) {
	b = b & ^block(1<<position)
}

func (b block) getBit(position uint64) bool {
	return b&block(1<<position) == 1
}

func (b block) or(other block) {
	b = b | other
}

func (b block) and(other block) {
	b = b & other
}

func (b block) equals(other block) bool {
	return b == other
}

func (b block) shl() {
	b = b << 1
}

func (b block) shr() {
	b = b >> 1
}

func (b block) String() string {
	return fmt.Sprintf(fmt.Sprintf("%%0%db", s), b)
}
