#include "output.h"

int pushUInt64(struct EncodingOutput* o, long long value, unsigned int len) {
    // Get only the least n bytes of the input (n = len)
    value = value & ((1 << len) - 1);
    // Pre-calculate some important numbers
	int thisBufferBitsAlloc = o->cursorPos % 8;
	int thisBufferBitsFree = 8 - thisBufferBitsAlloc;
	int nextBufferBitsAlloc = (len - thisBufferBitsFree) % 8;
	int nextBufferBitsFree = 8 - nextBufferBitsAlloc;
    // Step 1: Fill the rest of the old buffer
	/*posLow := len - thisBufferBitsFree;
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
	inputCurserPos := thisBufferBitsFree
	if thisBufferBitsAlloc+len < 8 {
		// Use same buffer for the next round
		o.cursorPos += len
		inputCurserPos = len
	} else {
		// Write buffer to output array
		o.bytes[o.getCurrentIndex()] = o.buffer
		o.buffer = 0
		o.cursorPos += thisBufferBitsFree
	}*/



    return 0;
}

int pushUInt32(struct EncodingOutput* o, long value, unsigned int len) {

}

int pushUInt16(struct EncodingOutput* o, int value, unsigned int len) {

}

int pushUInt8(struct EncodingOutput* o, short value, unsigned int len) {

}