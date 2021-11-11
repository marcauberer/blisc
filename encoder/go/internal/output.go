package internal

import (
	"fmt"
	"strings"
)

// Output represents the encoding output in form of a byte array
type Output interface {
	PushUInt64(uint64, uint64) error
	PushUInt32(uint32, uint64) error
	Conclude()
	ToByteArray() []byte
	ToString() string
}

type output struct {
	bytes     []byte
	buffer    byte
	cursorPos uint64
}

// CreateOutput returns the reference to a new output object
func CreateOutput(size uint64) Output {
	return &output{
		make([]byte, size), // byte array
		0,                  // buffer byte
		0,                  // current pos
	}
}

func (o *output) PushUInt64(value uint64, len uint64) error {
	fmt.Printf("Pushing %d with length %d\n", value, len)
	// Get only the least n bytes of the input (n = len)
	value = value & ((1 << len) - 1)
	// Pre-calculate some important numbers
	thisBufferBitsAlloc := o.cursorPos % 8
	thisBufferBitsFree := 8 - thisBufferBitsAlloc
	nextBufferBitsAlloc := (len - thisBufferBitsFree) % 8
	nextBufferBitsFree := 8 - nextBufferBitsAlloc
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
	o.cursorPos += thisBufferBitsFree
	//fmt.Printf("Buffer after: %08b\n", o.buffer)
	inputCurserPos := thisBufferBitsFree
	fmt.Printf("Input curser pos: %d\n", inputCurserPos)
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
	fmt.Printf("Final cursor pos: %d\n", o.cursorPos)
	return nil
}

func (o *output) PushUInt32(value uint32, len uint64) error {

	o.cursorPos += len
	return nil
}

// Conclude writes the buffer content to the stream. This method is ment to be called after sending the last value
func (o *output) Conclude() {
	if o.buffer > 0 {
		o.bytes[o.getCurrentIndex()] = o.buffer
		o.buffer = 0
	}
}

// ToByteArray returns the output as a byte array
func (o *output) ToByteArray() []byte {
	return o.bytes
}

// ToString returns the curent output as a binary string
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
