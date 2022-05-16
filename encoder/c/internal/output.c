#include "output.h"
#include <stdio.h>
#include <stdlib.h>
#include "mask.h"

#define BYTE_TO_BINARY(b)  \
  (b & 0x80 ? '1' : '0'), \
  (b & 0x40 ? '1' : '0'), \
  (b & 0x20 ? '1' : '0'), \
  (b & 0x10 ? '1' : '0'), \
  (b & 0x08 ? '1' : '0'), \
  (b & 0x04 ? '1' : '0'), \
  (b & 0x02 ? '1' : '0'), \
  (b & 0x01 ? '1' : '0') 

int pushUInt64(struct EncodingOutput* o, long long value, unsigned int len) {
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
	o->buffer |= (value & bitMask) << leftShift >> posLow;
	int inputCursorPos = thisBufferBitsFree;
	if (thisBufferBitsAlloc + len < 8) {
		// Use same buffer for the next round
		o->cursorPos += len;
		inputCursorPos = len;
	} else {
		// Write buffer to output array
		o->bytes[getCurrentOutputIndex(o)] = o->buffer;
		o->buffer = 0;
		o->cursorPos += thisBufferBitsFree;
	}
	// Step 2: Do insert steps for middle parts
	if (len > 8) {
		while (inputCursorPos < len - 8) {
			bitMask = createBitmask64ForRange(len - inputCursorPos - 8, len - inputCursorPos);
			inputCursorPos += 8;
			o->bytes[getCurrentOutputIndex(o)] = (value & bitMask) >> (len - inputCursorPos);
			o->cursorPos += 8;
		}
	}
	// Step 4: Fill the new buffer
	if (inputCursorPos < len) {
		bitMask = createBitmask64ForRange(nextBufferBitsAlloc, 0);
		o->buffer = (value & bitMask) << nextBufferBitsFree;
		o->cursorPos += nextBufferBitsAlloc;
	}

    return 0;
}

int pushUInt32(struct EncodingOutput* o, long value, unsigned int len) {
    // Get only the least n bytes of the input (n = len)
    value = value & ((1 << len) - 1);
    // Pre-calculate some important numbers
	int thisBufferBitsAlloc = o->cursorPos % 8;
	int thisBufferBitsFree = 8 - thisBufferBitsAlloc;
	int nextBufferBitsAlloc = (len - thisBufferBitsFree) % 8;
	int nextBufferBitsFree = 8 - nextBufferBitsAlloc;
    // Step 1: Fill the rest of the old buffer
	int posLow = thisBufferBitsFree > len ? 0 : len - thisBufferBitsFree;
	mask32 bitMask = createBitmask32ForRange(len, posLow);
	int leftShift = thisBufferBitsFree - len;
	o->buffer |= (value & bitMask) << leftShift >> posLow;
	int inputCursorPos = thisBufferBitsFree;
	if (thisBufferBitsAlloc + len < 8) {
		// Use same buffer for the next round
		o->cursorPos += len;
		inputCursorPos = len;
	} else {
		// Write buffer to output array
		o->bytes[getCurrentOutputIndex(o)] = o->buffer;
		o->buffer = 0;
		o->cursorPos += thisBufferBitsFree;
	}
	// Step 2: Do insert steps for middle parts
	if (len > 8) {
		while (inputCursorPos < len - 8) {
			bitMask = createBitmask32ForRange(len - inputCursorPos - 8, len - inputCursorPos);
			inputCursorPos += 8;
			o->bytes[getCurrentOutputIndex(o)] = (value & bitMask) >> (len - inputCursorPos);
			o->cursorPos += 8;
		}
	}
	// Step 4: Fill the new buffer
	if (inputCursorPos < len) {
		bitMask = createBitmask32ForRange(nextBufferBitsAlloc, 0);
		o->buffer = (value & bitMask) << nextBufferBitsFree;
		o->cursorPos += nextBufferBitsAlloc;
	}
    return 0;
}

int pushUInt16(struct EncodingOutput* o, int value, unsigned int len) {
    // Get only the least n bytes of the input (n = len)
    value = value & ((1 << len) - 1);
    // Pre-calculate some important numbers
	int thisBufferBitsAlloc = o->cursorPos % 8;
	int thisBufferBitsFree = 8 - thisBufferBitsAlloc;
	int nextBufferBitsAlloc = (len - thisBufferBitsFree) % 8;
	int nextBufferBitsFree = 8 - nextBufferBitsAlloc;
    // Step 1: Fill the rest of the old buffer
	int posLow = thisBufferBitsFree > len ? 0 : len - thisBufferBitsFree;
	mask16 bitMask = createBitmask16ForRange(len, posLow);
	int leftShift = thisBufferBitsFree - len;
	o->buffer |= (value & bitMask) << leftShift >> posLow;
	int inputCursorPos = thisBufferBitsFree;
	if (thisBufferBitsAlloc + len < 8) {
		// Use same buffer for the next round
		o->cursorPos += len;
		inputCursorPos = len;
	} else {
		// Write buffer to output array
		o->bytes[getCurrentOutputIndex(o)] = o->buffer;
		o->buffer = 0;
		o->cursorPos += thisBufferBitsFree;
	}
	// Step 2: Do insert steps for middle parts
	if (len > 8) {
		while (inputCursorPos < len - 8) {
			bitMask = createBitmask16ForRange(len - inputCursorPos - 8, len - inputCursorPos);
			inputCursorPos += 8;
			o->bytes[getCurrentOutputIndex(o)] = (value & bitMask) >> (len - inputCursorPos);
			o->cursorPos += 8;
		}
	}
	// Step 4: Fill the new buffer
	if (inputCursorPos < len) {
		bitMask = createBitmask16ForRange(nextBufferBitsAlloc, 0);
		o->buffer = (value & bitMask) << nextBufferBitsFree;
		o->cursorPos += nextBufferBitsAlloc;
	}
    return 0;
}

int pushUInt8(struct EncodingOutput* o, short value, unsigned int len) {
    // Get only the least n bytes of the input (n = len)
    value = value & ((1 << len) - 1);
    // Pre-calculate some important numbers
	int thisBufferBitsAlloc = o->cursorPos % 8;
	int thisBufferBitsFree = 8 - thisBufferBitsAlloc;
	int nextBufferBitsAlloc = (len - thisBufferBitsFree) % 8;
	int nextBufferBitsFree = 8 - nextBufferBitsAlloc;
    // Step 1: Fill the rest of the old buffer
	int posLow = thisBufferBitsFree > len ? 0 : len - thisBufferBitsFree;
	mask8 bitMask = createBitmask8ForRange(len, posLow);
	int leftShift = thisBufferBitsFree - len;
	o->buffer |= (value & bitMask) << leftShift >> posLow;
	int inputCursorPos = thisBufferBitsFree;
	if (thisBufferBitsAlloc + len < 8) {
		// Use same buffer for the next round
		o->cursorPos += len;
		inputCursorPos = len;
	} else {
		// Write buffer to output array
		o->bytes[getCurrentOutputIndex(o)] = o->buffer;
		o->buffer = 0;
		o->cursorPos += thisBufferBitsFree;
	}
	// Step 2: Do insert steps for middle parts
	if (len > 8) {
		while (inputCursorPos < len - 8) {
			bitMask = createBitmask8ForRange(len - inputCursorPos - 8, len - inputCursorPos);
			inputCursorPos += 8;
			o->bytes[getCurrentOutputIndex(o)] = (value & bitMask) >> (len - inputCursorPos);
			o->cursorPos += 8;
		}
	}
	// Step 4: Fill the new buffer
	if (inputCursorPos < len) {
		bitMask = createBitmask8ForRange(nextBufferBitsAlloc, 0);
		o->buffer = (value & bitMask) << nextBufferBitsFree;
		o->cursorPos += nextBufferBitsAlloc;
	}
    return 0;
}

int pushUInt1(struct EncodingOutput* o, short value, unsigned int len) {
	// Get only the least n bytes of the input (n = len)
    value = value & ((1 << len) - 1);
    // Pre-calculate some important numbers
	int thisBufferBitsAlloc = o->cursorPos % 8;
	int thisBufferBitsFree = 8 - thisBufferBitsAlloc;
	int nextBufferBitsAlloc = (len - thisBufferBitsFree) % 8;
	int nextBufferBitsFree = 8 - nextBufferBitsAlloc;
    // Step 1: Fill the rest of the old buffer
	int posLow = thisBufferBitsFree > len ? 0 : len - thisBufferBitsFree;
	mask1 bitMask = createBitmask1ForRange(len, posLow);
	int leftShift = thisBufferBitsFree - len;
	o->buffer |= (value & bitMask) << leftShift >> posLow;
	int inputCursorPos = thisBufferBitsFree;
	if (thisBufferBitsAlloc + len < 8) {
		// Use same buffer for the next round
		o->cursorPos += len;
		inputCursorPos = len;
	} else {
		// Write buffer to output array
		o->bytes[getCurrentOutputIndex(o)] = o->buffer;
		o->buffer = 0;
		o->cursorPos += thisBufferBitsFree;
	}
	// Step 2: Do insert steps for middle parts
	if (len > 8) {
		while (inputCursorPos < len - 8) {
			bitMask = createBitmask1ForRange(len - inputCursorPos - 8, len - inputCursorPos);
			inputCursorPos += 8;
			o->bytes[getCurrentOutputIndex(o)] = (value & bitMask) >> (len - inputCursorPos);
			o->cursorPos += 8;
		}
	}
	// Step 4: Fill the new buffer
	if (inputCursorPos < len) {
		bitMask = createBitmask1ForRange(nextBufferBitsAlloc, 0);
		o->buffer = (value & bitMask) << nextBufferBitsFree;
		o->cursorPos += nextBufferBitsAlloc;
	}
	return 0;
}

void conclude(struct EncodingOutput* o) {
	if (o->buffer > 0) {
		o->bytes[getCurrentOutputIndex(o)] = o->buffer;
		o->buffer = 0;
	}
}

void outputToString(struct EncodingOutput* o, char* result, int size) {
	for (int i = 0; i < size; i++) {
		result += sprintf(result, "%c%c%c%c%c%c%c%c ", BYTE_TO_BINARY(o->bytes[i]));
	}
	result += '\0';
}

unsigned long long getCurrentOutputIndex(struct EncodingOutput* o) {
	return o->cursorPos / 8;
}