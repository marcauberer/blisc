#pragma once

// Constants
extern const int CONFIG_ERR_LOAD_FAILED;

// Structs
struct EncoderConfigVersion {
    int major;
    int minor;
};

struct EncoderConfigField {
    char name[100];
    int type;
    unsigned int pos;
    unsigned int len;
    int bias;
    double mul;
};

struct EncoderConfig {
    struct EncoderConfigVersion version;
    int fieldCount;
    struct EncoderConfigField fields[100];
};

// Functions
int loadConfig(struct EncoderConfig* config, const char* configPath);
int getTotalConfigLength(struct EncoderConfig* config);