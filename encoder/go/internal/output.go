package internal

import (
	"fmt"
	"strings"
)

// Output represents the encoding output in form of a byte array
type Output interface {
	PushUInt64(uint64, uint64) error
	PushUInt32(uint32, uint64) error
	PushUInt16(uint16, uint64) error
	PushUInt8(uint8, uint64) error
	PushUInt1(uint8, uint64) error
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

// PushUInt64 appends an uint64 to the output
func (o *output) PushUInt64(value uint64, len uint64) error {
	// Get only the least n bytes of the input (n = len)
	value = value & ((1 << len) - 1)
	// Pre-calculate some important numbers
	thisBufferBitsAlloc := o.cursorPos % 8
	thisBufferBitsFree := 8 - thisBufferBitsAlloc
	nextBufferBitsAlloc := (len - thisBufferBitsFree) % 8
	nextBufferBitsFree := 8 - nextBufferBitsAlloc
	// Step 1: Fill the rest of the old buffer
	posLow := len - thisBufferBitsFree
	if thisBufferBitsFree > len {
		posLow = 0
	}
	bitMask, err := createBitmask64ForRange(len, posLow)
	if err != nil {
		return err
	}
	leftShift := thisBufferBitsFree - len
	if len >= 8 {
		leftShift = 0
	}
	o.buffer |= byte((value & bitMask) << leftShift >> posLow)
	inputCursorPos := thisBufferBitsFree
	if thisBufferBitsAlloc+len < 8 {
		// Use same buffer for the next round
		o.cursorPos += len
		inputCursorPos = len
	} else {
		// Write buffer to output array
		o.bytes[o.getCurrentIndex()] = o.buffer
		o.buffer = 0
		o.cursorPos += thisBufferBitsFree
	}
	// Step 2: Do insert steps for middle parts
	if len > 8 {
		for inputCursorPos < len-8 {
			bitMask, err := createBitmask64ForRange(len-inputCursorPos-8, len-inputCursorPos)
			if err != nil {
				return err
			}
			inputCursorPos += 8
			o.bytes[o.getCurrentIndex()] = byte((value & bitMask) >> (len - inputCursorPos))
			o.cursorPos += 8
		}
	}
	// Step 3: Fill the new buffer
	if inputCursorPos < len {
		bitMask, err = createBitmask64ForRange(nextBufferBitsAlloc, 0)
		if err != nil {
			return err
		}
		o.buffer = byte((value & bitMask) << nextBufferBitsFree)
		o.cursorPos += nextBufferBitsAlloc
	}
	return nil
}

// PushUInt32 appends an uint32 to the output
func (o *output) PushUInt32(value uint32, len uint64) error {
	// Get only the least n bytes of the input (n = len)
	value = value & ((1 << len) - 1)
	// Pre-calculate some important numbers
	thisBufferBitsAlloc := o.cursorPos % 8
	thisBufferBitsFree := 8 - thisBufferBitsAlloc
	nextBufferBitsAlloc := (len - thisBufferBitsFree) % 8
	nextBufferBitsFree := 8 - nextBufferBitsAlloc
	// Fill the rest of the old buffer
	bitMask, err := createBitmask32ForRange(len, len-thisBufferBitsFree)
	if err != nil {
		return err
	}
	o.buffer |= byte((value & bitMask) >> (len - thisBufferBitsFree))
	o.bytes[o.getCurrentIndex()] = o.buffer
	o.cursorPos += thisBufferBitsFree
	inputCursorPos := thisBufferBitsFree
	// Do insert steps for middle parts
	for inputCursorPos < len-8 {
		bitMask, err := createBitmask32ForRange(len-inputCursorPos-8, len-inputCursorPos)
		if err != nil {
			return err
		}
		inputCursorPos += 8
		o.bytes[o.getCurrentIndex()] = byte((value & bitMask) >> (len - inputCursorPos))
		o.cursorPos += 8
	}
	// Fill the new buffer
	bitMask, err = createBitmask32ForRange(nextBufferBitsAlloc, 0)
	if err != nil {
		return err
	}
	o.buffer = byte((value & bitMask) << nextBufferBitsFree)
	o.cursorPos += nextBufferBitsAlloc
	return nil
}

// PushUInt16 appends an uint16 to the output
func (o *output) PushUInt16(value uint16, len uint64) error {
	// Get only the least n bytes of the input (n = len)
	value = value & ((1 << len) - 1)
	// Pre-calculate some important numbers
	thisBufferBitsAlloc := o.cursorPos % 8
	thisBufferBitsFree := 8 - thisBufferBitsAlloc
	nextBufferBitsAlloc := (len - thisBufferBitsFree) % 8
	nextBufferBitsFree := 8 - nextBufferBitsAlloc
	// Fill the rest of the old buffer
	bitMask, err := createBitmask16ForRange(len, len-thisBufferBitsFree)
	if err != nil {
		return err
	}
	o.buffer |= byte((value & bitMask) >> (len - thisBufferBitsFree))
	o.bytes[o.getCurrentIndex()] = o.buffer
	o.cursorPos += thisBufferBitsFree
	inputCursorPos := thisBufferBitsFree
	// Do insert steps for middle parts
	for inputCursorPos < len-8 {
		bitMask, err := createBitmask16ForRange(len-inputCursorPos-8, len-inputCursorPos)
		if err != nil {
			return err
		}
		inputCursorPos += 8
		o.bytes[o.getCurrentIndex()] = byte((value & bitMask) >> (len - inputCursorPos))
		o.cursorPos += 8
	}
	// Fill the new buffer
	bitMask, err = createBitmask16ForRange(nextBufferBitsAlloc, 0)
	if err != nil {
		return err
	}
	o.buffer = byte((value & bitMask) << nextBufferBitsFree)
	o.cursorPos += nextBufferBitsAlloc
	return nil
}

// PushUInt8 appends an uint8 to the output
func (o *output) PushUInt8(value uint8, len uint64) error {
	// Get only the least n bytes of the input (n = len)
	value = value & ((1 << len) - 1)
	// Pre-calculate some important numbers
	thisBufferBitsAlloc := o.cursorPos % 8
	thisBufferBitsFree := 8 - thisBufferBitsAlloc
	nextBufferBitsAlloc := (len - thisBufferBitsFree) % 8
	nextBufferBitsFree := 8 - nextBufferBitsAlloc
	// Fill the rest of the old buffer
	bitMask, err := createBitmask8ForRange(len, len-thisBufferBitsFree)
	if err != nil {
		return err
	}
	o.buffer |= byte((value & bitMask) >> (len - thisBufferBitsFree))
	o.bytes[o.getCurrentIndex()] = o.buffer
	o.cursorPos += thisBufferBitsFree
	inputCursorPos := thisBufferBitsFree
	// Do insert steps for middle parts
	for inputCursorPos < len-8 {
		bitMask, err := createBitmask8ForRange(len-inputCursorPos-8, len-inputCursorPos)
		if err != nil {
			return err
		}
		inputCursorPos += 8
		o.bytes[o.getCurrentIndex()] = byte((value & bitMask) >> (len - inputCursorPos))
		o.cursorPos += 8
	}
	// Fill the new buffer
	bitMask, err = createBitmask8ForRange(nextBufferBitsAlloc, 0)
	if err != nil {
		return err
	}
	o.buffer = byte((value & bitMask) << nextBufferBitsFree)
	o.cursorPos += nextBufferBitsAlloc
	return nil
}

// PushUInt1 appends the least significant bit of an uint8 to the output
func (o *output) PushUInt1(value uint8, len uint64) error {
	// Get only the least n bytes of the input (n = len)
	value = value & ((1 << len) - 1)
	// Pre-calculate some important numbers
	thisBufferBitsAlloc := o.cursorPos % 8
	thisBufferBitsFree := 8 - thisBufferBitsAlloc
	nextBufferBitsAlloc := (len - thisBufferBitsFree) % 8
	nextBufferBitsFree := 8 - nextBufferBitsAlloc
	// Fill the rest of the old buffer
	bitMask, err := createBitmask1ForRange(len, len-thisBufferBitsFree)
	if err != nil {
		return err
	}
	o.buffer |= byte((value & bitMask) >> (len - thisBufferBitsFree))
	o.bytes[o.getCurrentIndex()] = o.buffer
	o.cursorPos += thisBufferBitsFree
	inputCursorPos := thisBufferBitsFree
	// Do insert steps for middle parts
	for inputCursorPos < len-8 {
		bitMask, err := createBitmask1ForRange(len-inputCursorPos-8, len-inputCursorPos)
		if err != nil {
			return err
		}
		inputCursorPos += 8
		o.bytes[o.getCurrentIndex()] = byte((value & bitMask) >> (len - inputCursorPos))
		o.cursorPos += 8
	}
	// Fill the new buffer
	bitMask, err = createBitmask1ForRange(nextBufferBitsAlloc, 0)
	if err != nil {
		return err
	}
	o.buffer = byte((value & bitMask) << nextBufferBitsFree)
	o.cursorPos += nextBufferBitsAlloc
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
	for i, item := range o.bytes {
		byteStrings[i] = fmt.Sprintf("%08b", item)
	}
	return strings.Join(byteStrings, " ")
}

func (o *output) getCurrentIndex() uint64 {
	return o.cursorPos / 8
}
