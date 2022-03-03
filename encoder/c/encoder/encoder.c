#include "encoder.h"

const int ENCODER_ERR_CONFIG_NOT_SET = -1;

int initializeEncoder(struct Encoder* e) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;
    // Calculate total length of output
    int totalLengthBits = getTotalConfigLength(e->encoderConfig);
    int totalLengthBytes = totalLengthBits / 8;
    if (totalLengthBits % 8 > 0) totalLengthBytes++;
    // Create output
    struct EncodingOutput output;
    char byteArray[totalLengthBytes];
    output.bytes = byteArray;
    output.cursorPos = 0;
    e->output = &output;
}

int encodeInt(struct Encoder* e, int value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;

    return 0;
}

int encodeDouble(struct Encoder* e, double value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;

    return 0;
}

int encodeString(struct Encoder* e, char* value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;

    return 0;
}

int encodeBool(struct Encoder* e, bool value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;

    return 0;
}