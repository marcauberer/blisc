#include <stdio.h>

#include "config/config.h"
#include "encoder/encoder.h"

struct TestData {
    double pm10;
    double pm2_5;
    double temperature;
    double humidity;
    double pressure;
};

int main(int argc, char const *argv[]) {
    // Encode test payload
    struct TestData testData = { 12.43, 6.14, 25.124, 78.01, 100001.9 };

    // Create encoder and load config
    struct Encoder encoder;
    struct EncoderConfig config;
    loadConfig(&config, "../../config/client-config.bin");
    encoder.encoderConfig = &config;
    int encodedSize = initializeEncoder(&encoder);

    // Create encoding output
    char byteArray[encodedSize];
    struct EncodingOutput output;
    output.bytes = byteArray;
    output.cursorPos = 0;
    encoder.output = &output;

    // Encode tets data
    encodeDouble(&encoder, testData.pm10, "pm10");
    encodeDouble(&encoder, testData.pm2_5, "pm2_5");
    encodeDouble(&encoder, testData.temperature, "temperature");
    encodeDouble(&encoder, testData.humidity, "humidity");
    encodeDouble(&encoder, testData.pressure, "pressure");

    char result[encodedSize * 9]; // Allocate enough for the spaces between the blocks and the NULL terminator
    outputToString(encoder.output, result, encodedSize);
    printf("Stringified: %s\n", result);

    return 0;
}
