#include "output.h"
#include <stdio.h>
#include "mask.h"

int pushUInt64(struct EncodingOutput* o, long long value, unsigned int len) {
	printf("Int64: %llu\n", value);
    // Get only the least n bytes of the input (n = len)
    value = value & ((1 << len) - 1);
    // Pre-calculate some important numbers
	int thisBufferBitsAlloc = o->cursorPos % 8;
	int thisBufferBitsFree = 8 - thisBufferBitsAlloc;
	int nextBufferBitsAlloc = (len - thisBufferBitsFree) % 8;
	int nextBufferBitsFree = 8 - nextBufferBitsAlloc;
    // Step 1: Fill the rest of the old buffer
	int posLow = thisBufferBitsFree > len ? 0 : len - thisBufferBitsFree;
	mask64 bitMask = createBitmask64ForRange(len, posLow);
	int leftShift = thisBufferBitsFree - len;
	o.buffer |= (value & bitMask) << leftShift >> posLow;
	int inputCursorPos = thisBufferBitsFree;
	if (thisBufferBitsAlloc + len < 8) {
		// Use same buffer for the next round
		o.cursorPos += len;
		inputCursorPos = len;
	} else {
		// Write buffer to output array
		o.bytes[getCurrentOutputIndex(o)] = o.buffer;
		o.buffer = 0;
		o.cursorPos += thisBufferBitsFree;
	}


    return 0;
}

int pushUInt32(struct EncodingOutput* o, long value, unsigned int len) {

	return 0;
}

int pushUInt16(struct EncodingOutput* o, int value, unsigned int len) {

	return 0;
}

int pushUInt8(struct EncodingOutput* o, short value, unsigned int len) {

	return 0;
}

void conclude(struct EncodingOutput* o) {
	if (o.buffer > 0) {
		o.bytes[getCurrentOutputIndex(output)] = o.buffer;
		o.buffer = 0;
	}
}

unsigned long long getCurrentOutputIndex(struct EncodingOutput* o) {
	return o.cursorPos / 8;
}