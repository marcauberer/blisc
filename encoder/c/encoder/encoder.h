#pragma once

#include <stdbool.h>
#include "../config/config.h"
#include "../internal/output.h"

// Constants
extern const int ENCODER_ERR_CONFIG_NOT_SET;

// Structs
struct Encoder {
    struct EncoderConfig* encoderConfig;
    struct EncodingOutput* output;
};

// Functions
int initializeEncoder(struct Encoder* encoder);
int encodeInt(struct Encoder* encoder, int value, char* name);
int encodeDouble(struct Encoder* encoder, double value, char* name);
int encodeString(struct Encoder* encoder, char* value, char* name);
int encodeBool(struct Encoder* encoder, bool value, char* name);