#include <stdio.h>

#include "config/config.h"

struct TestData {
    float pm10;
    float pm2_5;
    float temperature;
    float humidity;
    float pressure;
};

int main(int argc, char const *argv[]) {
    // Encode test payload
    struct TestData testData = { 12.43, 6.14, 25.124, 78.01, 100001.9 };

    struct EncoderConfig config;
    loadConfig(&config, "../../config/client-config.bin");

    //encoder.push(data.pm10, "pm10");

    // Print result to the console
    printf("%f\n", testData.pm10);
    return 0;
}
