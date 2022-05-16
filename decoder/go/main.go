package main

import (
	"blisc-decoder/decoder"
	"fmt"
)

type data struct {
	Pm10        float32
	Pm2_5       float32
	Temperature float32
	Humidity    float32
	Pressure    float32
}

func main() {
	encoded := []byte{0b00000001, 0b11110000, 0b00000111, 0b10110100, 0b01011110, 0b00011000, 0b11001011, 0b00100000}
	var decoded data
	d := decoder.GetDecoder()
	err := d.LoadConfigFromFile("../../config/client-config.json")
	if err != nil {
		panic(err)
	}
	err = d.Decode(encoded, &decoded)
	if err != nil {
		panic(err)
	}
	// Print result to the console
	fmt.Println("PM10:", decoded.Pm10)
	fmt.Println("PM2.5:", decoded.Pm2_5)
	fmt.Println("Temperature:", decoded.Temperature)
	fmt.Println("Humidity:", decoded.Humidity)
	fmt.Println("Pressure:", decoded.Pressure)
}
