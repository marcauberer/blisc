#include "config.h"

#include <stdio.h>
#include <string.h>
#include <stdint.h>

int loadConfig(struct EncoderConfig* config, const char* configPath) {
    // Cancel if the file cannot be opened
    FILE* binaryFile = fopen(configPath, "rb");
    if (!binaryFile) {
        printf("Could not open file");
        return -1;
    }

    // Read major version
    (void)!fread(&config->version.major, 4, 1, binaryFile);

    // Read minor version
    (void)!fread(&config->version.minor, 4, 1, binaryFile);

    // Read field count
    int fieldCount;
    (void)!fread(&fieldCount, 4, 1, binaryFile);
    struct EncoderConfigField fields[fieldCount];

    // Read the fields
    for (int i = 0; i < fieldCount; i++) {
        // Read field name char by char
        char fieldName[100];
        int j = 0;
        while ((fieldName[j] = fgetc(binaryFile)) != '\0') j++;
        fields[j].name = fieldName;
        printf("Field name: %s\n", fieldName);

        // Read type
        (void)!fread(&fields[i].type, 1, 1, binaryFile);

        // Read pos
        (void)!fread(&fields[i].pos, 4, 1, binaryFile);

        // Read type
        (void)!fread(&fields[i].len, 4, 1, binaryFile);

        // Read bias
        (void)!fread(&fields[i].bias, 4, 1, binaryFile);

        // Read mul
        (void)!fread(&fields[i].mul, 8, 1, binaryFile);
    }
    config->fields = fields;

    // Close binary file
    fclose(binaryFile);

    return 0;
}