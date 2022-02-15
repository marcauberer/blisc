#include "config.h"

#include <stdio.h>

int loadConfig(struct EncoderConfig* config, const char* configPath) {
    FILE* binaryFile;
    int intBuffer;
    

    // Cancel if the file cannot be opened
    if ((binaryFile = fopen(configPath, "r")) == NULL) return -1;

    fread((void*) &intBuffer, 4, 1, binaryFile);
    config->version.major = (int) intBuffer;
    fread((void*) &intBuffer, 4, 1, binaryFile);
    config->version.minor = (int) intBuffer;

    return 0;
}