#include "encoder.h"
#include <string.h>
#include <stdio.h>

const int ENCODER_ERR_CONFIG_NOT_SET = -1;

int initializeEncoder(struct Encoder* e) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;
    // Calculate total length of output
    int totalLengthBits = getTotalConfigLength(e->encoderConfig);
    int totalLengthBytes = totalLengthBits / 8;
    if (totalLengthBits % 8 > 0) totalLengthBytes++;
    // Create output
    char byteArray[totalLengthBytes];
    struct EncodingOutput output;
    output.bytes = byteArray;
    output.cursorPos = 0;
    e->output = &output;
}

struct EncoderConfigField* findConfigField(struct EncoderConfig* c, char* name) {
    for (int i = 0; i < c->fieldCount; i++) {
        if (strcmp(c->fields[i].name, name) == 0)
            return &c->fields[i];
    }
    return 0;
}

int encodeInt(struct Encoder* e, int value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;
    // Get config field
    struct EncoderConfigField* configField = findConfigField(e->encoderConfig, name);
    // Apply bias and mul
    value += configField->bias;
    value *= configField->mul;
    printf("Int: %d\n", value);

    return 0;
}

int encodeDouble(struct Encoder* e, double value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;
    // Get config field
    struct EncoderConfigField* configField = findConfigField(e->encoderConfig, name);
    // Apply bias and mul
    value += configField->bias;
    value *= configField->mul;
    pushUInt64(e->output, value, configField->len);
    printf("Double: %f\n", value);

    return 0;
}

int encodeString(struct Encoder* e, char* value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;
    // Get config field
    struct EncoderConfigField* configField = findConfigField(e->encoderConfig, name);

    return 0;
}

int encodeBool(struct Encoder* e, bool value, char* name) {
    // Abort if config is not set
    if (!e->encoderConfig) return ENCODER_ERR_CONFIG_NOT_SET;
    // Get config field
    struct EncoderConfigField* configField = findConfigField(e->encoderConfig, name);

    return 0;
}