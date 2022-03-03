#include "encoder.h"

const int ENCODER_ERR_CONFIG_NOT_SET = -1;

int encodeInt(struct Encoder* e, int value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;
    // Calculate total length of output
    int totalLengthBits = getTotalConfigLength(e->encoderConfig);

    return 0;
}

int encodeDouble(struct Encoder* e, double value, char* name) {

    return 0;
}

int encodeString(struct Encoder* e, char* value, char* name) {

    return 0;
}

int encodeBool(struct Encoder* e, bool value, char* name) {

    return 0;
}