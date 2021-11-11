package internal

import (
	"fmt"
	"strings"
)

type Output interface {
	PushUInt64(uint64, uint64) error
	PushUInt32(uint32, uint64) error
	ToByteArray() []byte
	ToString() string
}

type output struct {
	bytes     []byte
	buffer    byte
	cursorPos uint64
}

func (o *output) PushUInt64(value uint64, len uint64) error {
	fmt.Printf("Pushing %d with length %d\n", value, len)
	// Get only the least n bytes of the input (n = len)
	value = value & ((1 << len) - 1)
	// Pre-calculate some important numbers
	thisBufferBitsAlloc := o.getAllocBufferBits()
	//fmt.Printf("This buffer bits alloc: %d\n", thisBufferBitsAlloc)
	thisBufferBitsFree := 8 - thisBufferBitsAlloc
	//fmt.Printf("This buffer bits free: %d\n", thisBufferBitsFree)
	nextBufferBitsAlloc := (len - thisBufferBitsFree) % 8
	//fmt.Printf("Next buffer bits free: %d\n", nextBufferBitsFree)
	nextBufferBitsFree := 8 - nextBufferBitsAlloc
	//fmt.Printf("Next buffer bits alloc: %d\n", nextBufferBitsAlloc)
	// Fill the rest of the old buffer
	bitMask, err := createBitmaskForRange(len, len-thisBufferBitsFree)
	if err != nil {
		return err
	}
	fmt.Printf("Value:           %064b\n", value)
	fmt.Printf("Bit mask buffer: %064b\n", bitMask)
	//fmt.Printf("Buffer before: %08b\n", o.buffer)
	//fmt.Printf("Buffer result: %d\n", uint8(value&bitMask))
	o.buffer |= byte((value & bitMask) >> (len - thisBufferBitsFree))
	o.bytes[o.getCurrentIndex()] = o.buffer
	o.buffer = byte(0)
	o.cursorPos += 8
	//fmt.Printf("Buffer after: %08b\n", o.buffer)
	inputCurserPos := thisBufferBitsFree
	fmt.Printf("Input curser pos free: %d\n", inputCurserPos)
	// Do insert steps for middle parts
	for inputCurserPos < len-8 {
		bitMask, err := createBitmaskForRange(len-inputCurserPos-8, len-inputCurserPos)
		if err != nil {
			return err
		}
		fmt.Printf("Value:           %064b\n", value)
		fmt.Printf("Bit mask middle: %064b\n", bitMask)
		inputCurserPos += 8
		o.bytes[o.getCurrentIndex()] = byte((value & bitMask) >> (len - inputCurserPos))
		o.cursorPos += 8
	}
	// Fill the new buffer
	bitMask, err = createBitmaskForRange(nextBufferBitsAlloc, 0)
	if err != nil {
		return err
	}
	fmt.Printf("Value:        %064b\n", value)
	fmt.Printf("Bit mask end: %064b\n", bitMask)
	o.buffer = byte((value & bitMask) << nextBufferBitsFree)
	fmt.Printf("Buffer after: %08b\n", o.buffer)
	o.cursorPos += nextBufferBitsAlloc
	return nil
}

func (o *output) PushUInt32(value uint32, len uint64) error {

	o.cursorPos += len
	return nil
}

func (o *output) ToByteArray() []byte {
	return o.bytes
}

func (o *output) ToString() string {
	var byteStrings = make([]string, len(o.bytes))
	for _, n := range o.bytes {
		byteString := fmt.Sprintf("%08b", n)
		byteStrings = append(byteStrings, byteString)
	}
	return strings.Join(byteStrings, " ")
}

func (o *output) getCurrentIndex() uint64 {
	fmt.Println("Cursorpos:", o.cursorPos)
	return o.cursorPos / 8
}

func (o *output) getAllocBufferBits() uint64 {
	return o.cursorPos % 8
}

func CreateOutput(size uint64) Output {
	return &output{
		make([]byte, size), // byte array
		byte(0),            // buffer byte
		uint64(0),          // current pos
	}
}
