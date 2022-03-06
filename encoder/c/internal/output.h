#pragma once

struct EncodingOutput {
    char* bytes;
    unsigned char buffer;
    unsigned long long cursorPos;
};

int pushUInt64(struct EncodingOutput* output, long long value, unsigned int len);
int pushUInt32(struct EncodingOutput* output, long value, unsigned int len);
int pushUInt16(struct EncodingOutput* output, int value, unsigned int len);
int pushUInt8(struct EncodingOutput* output, short value, unsigned int len);
void conclude(struct EncodingOutput* output);
void outputToString(struct EncodingOutput* output, char* result, int size);
unsigned long long getCurrentOutputIndex(struct EncodingOutput* output);