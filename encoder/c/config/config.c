#include "config.h"

#include <stdio.h>

int loadConfig(struct EncoderConfig* config, const char* configPath) {
    FILE* binaryFile;

    // Cancel if the file cannot be opened
    if ((binaryFile = fopen(configPath, "r")) == NULL) return -1;

    fread((void*) &config->version.major, 4, 1, binaryFile);
    fread((void*) &config->version.minor, 4, 1, binaryFile);

    return 0;
}