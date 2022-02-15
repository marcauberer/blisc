#include "config.h"

#include <stdio.h>

int loadConfig(struct EncoderConfig* config, const char* configPath) {
    FILE* binaryFile;
    char buffer;

    // Cancel if the file cannot be opened
    if ((binaryFile = fopen(configPath, "r")) == NULL) return -1;

    //fread((void*) &buffer, 1, 1, binaryFile);
    //config->version.major = (int) buffer;

    return 0;
}