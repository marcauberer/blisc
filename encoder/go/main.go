package main

import (
	"blisc-encoder/encoder"
	"fmt"
)

type data struct {
	pm10        float32
	pm2_5       float32
	temperature float32
	humidity    float32
	pressure    float32
}

func main() {
	data := data{
		pm10:        12.43,
		pm2_5:       6.14,
		temperature: 25.124,
		humidity:    78.01,
		pressure:    100001.9,
	}
	// Encode test payload
	e := encoder.GetEncoder()
	err := e.LoadConfigFromFile("../../config/client-config.bin")
	if err != nil {
		panic(err)
	}
	result, err := e.Encode(&data)
	if err != nil {
		panic(err)
	}
	// Print result to the console
	for _, n := range result {
		fmt.Printf("%08b ", n)
	}
}
