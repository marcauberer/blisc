#pragma once

struct EncodingOutput {
    char* bytes;
    char buffer;
    long long cursorPos;
};

int pushLongLong(struct EncodingOutput* output, long long value, long long len);