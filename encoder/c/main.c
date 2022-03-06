#include <stdio.h>

#include "config/config.h"
#include "encoder/encoder.h"

#define BYTE_TO_BINARY(b)  \
  (b & 0x80 ? '1' : '0'), \
  (b & 0x40 ? '1' : '0'), \
  (b & 0x20 ? '1' : '0'), \
  (b & 0x10 ? '1' : '0'), \
  (b & 0x08 ? '1' : '0'), \
  (b & 0x04 ? '1' : '0'), \
  (b & 0x02 ? '1' : '0'), \
  (b & 0x01 ? '1' : '0') 

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

    for (int i = 0; i < encodedSize; i++)
        printf("Output %d: %c%c%c%c%c%c%c%c\n", i, BYTE_TO_BINARY(encoder.output->bytes[i]));

    // Print result to the console
    printf("%f\n", testData.pm10);
    return 0;
}
