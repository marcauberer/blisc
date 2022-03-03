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

    // Encode tets data
    encodeDouble(&encoder, testData.pm10, "pm10");

    // Print result to the console
    printf("%f\n", testData.pm10);
    return 0;
}
