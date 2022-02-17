struct EncoderConfigVersion {
    int major;
    int minor;
};

struct EncoderConfigField {
    char* name;
    int type;
    unsigned int pos;
    unsigned int len;
    int bias;
    double mul;
};

struct EncoderConfig {
    struct EncoderConfigVersion version;
    struct EncoderConfigField* fields;
};

// Functions
int loadConfig(struct EncoderConfig* config, const char* configPath);