#pragma once

#include <stdbool.h>
#include "../config/config.h"

// Constants
extern const int ENCODER_ERR_CONFIG_NOT_SET;

// Structs
struct Encoder {
    struct EncoderConfig* encoderConfig;
};

// Functions
int encodeInt(struct Encoder* encoder, int value, char* name);
int encodeDouble(struct Encoder* encoder, double value, char* name);
int encodeString(struct Encoder* encoder, char* value, char* name);
int encodeBool(struct Encoder* encoder, bool value, char* name);