package bitarray

import "strings"

type bitArray struct {
	blocks []block
}

func (ba *bitArray) InsertUInt64At(pos, value uint64) {
	other := newBitArray(uint64(len(ba.blocks)) * s)
	other.blocks[len(ba.blocks)-8] = block(value & 0xFF00000000000000 >> 56)
	other.blocks[len(ba.blocks)-7] = block(value & 0x00FF000000000000 >> 48)
	other.blocks[len(ba.blocks)-6] = block(value & 0x0000FF0000000000 >> 40)
	other.blocks[len(ba.blocks)-5] = block(value & 0x000000FF00000000 >> 32)
	other.blocks[len(ba.blocks)-4] = block(value & 0x00000000FF000000 >> 24)
	other.blocks[len(ba.blocks)-3] = block(value & 0x0000000000FF0000 >> 16)
	other.blocks[len(ba.blocks)-2] = block(value & 0x000000000000FF00 >> 8)
	other.blocks[len(ba.blocks)-1] = block(value & 0x00000000000000FF)
	other.Shl(uint64(len(ba.blocks))*s - 64 - pos)
	*ba = ba.Or(*other)
}

func (ba *bitArray) Shl(num uint64) {
	for i := uint64(0); i < num; i++ {
		carry := false
		// Loop through blocks in reverse order
		for j := len(ba.blocks) - 1; j >= 0; j-- {
			tmp := ba.blocks[j].getBit(0) // Get value of leftest bit
			ba.blocks[j].shl()
			if carry {
				ba.blocks[j].setBit(s - 1)
			}
			carry = tmp
		}
		// If the last carry is true, there is an additional block required
		if carry {
			ba.blocks = append([]block{1}, ba.blocks...)
		}
	}
}

func (ba *bitArray) Shr(num uint64) {

}

func (ba *bitArray) Or(other bitArray) bitArray {
	if len(ba.blocks) >= len(other.blocks) {
		newArray := ba.copy()
		for i := 1; i <= len(other.blocks); i++ {
			newArray.blocks[len(newArray.blocks)-i].or(other.blocks[len(other.blocks)-i])
		}
		return newArray
	}
	return other.Or(*ba)
}

func (ba *bitArray) copy() bitArray {
	blocks := make([]block, len(ba.blocks))
	copy(blocks, ba.blocks)
	return bitArray{blocks: blocks}
}

func (ba *bitArray) ToByteArray() []byte {
	// ToDo: Implement
	return []byte{}
}

func (ba *bitArray) String() string {
	var stringifiedBlocks = make([]string, len(ba.blocks))
	for i, block := range ba.blocks {
		stringifiedBlocks[i] = block.String()
	}
	return strings.Join(stringifiedBlocks, " ")
}

func getIndexAndRemainder(k uint64) (uint64, uint64) {
	return k / s, k % s
}

func newBitArray(size uint64) *bitArray {
	i, r := getIndexAndRemainder(size)
	if r > 0 {
		i++
	}
	return &bitArray{
		blocks: make([]block, i),
	}
}

func NewBitArray(size uint64) BitArray {
	return newBitArray(size)
}
