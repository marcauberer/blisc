#include "config.h"

#include <stdio.h>
#include <string.h>
#include <stdint.h>

const int CONFIG_ERR_LOAD_FAILED = -1;

int loadConfig(struct EncoderConfig* config, const char* configPath) {
    // Cancel if the file cannot be opened
    FILE* binaryFile = fopen(configPath, "rb");
    if (!binaryFile) return CONFIG_ERR_LOAD_FAILED;

    // Read major version
    (void)!fread(&config->version.major, 4, 1, binaryFile);

    // Read minor version
    (void)!fread(&config->version.minor, 4, 1, binaryFile);

    // Read field count
    (void)!fread(&config->fieldCount, 4, 1, binaryFile);

    // Read the fields
    for (int i = 0; i < config->fieldCount; i++) {
        // Read field name char by char
        int j = 0;
        while ((config->fields[i].name[j] = fgetc(binaryFile)) != '\0') j++;

        // Read type
        (void)!fread(&config->fields[i].type, 1, 1, binaryFile);

        // Read pos
        (void)!fread(&config->fields[i].pos, 4, 1, binaryFile);

        // Read len
        (void)!fread(&config->fields[i].len, 4, 1, binaryFile);

        // Read bias
        (void)!fread(&config->fields[i].bias, 4, 1, binaryFile);

        // Read mul
        (void)!fread(&config->fields[i].mul, 8, 1, binaryFile);
    }

    // Close binary file
    fclose(binaryFile);

    return 0;
}

int getTotalConfigLength(struct EncoderConfig* config) {
    if (config->fieldCount == 0) return 0;
    struct EncoderConfigField lastField = config->fields[config->fieldCount -1];
    return lastField.pos + lastField.len;
}