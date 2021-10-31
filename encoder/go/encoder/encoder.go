package encoder

import (
	"errors"
	"math"
)

// Encoder represents a binary compression encoder
type Encoder struct {
	Config *EncoderConfig
}

// GetEncoder creates and returns a new binary compression encoder
func GetEncoder() *Encoder {
	return &Encoder{
		Config: &EncoderConfig{
			Version: DefaultEncoderConfigVersion,
		},
	}
}

// Encode executes the compression process and returns the output as byte array
func (e Encoder) Encode(input interface{}) ([]byte, error) {
	// Abort if no config is set
	if e.Config == nil {
		return []byte{}, errors.New("no encoding configuration was set")
	}
	// Calculate total length of output
	totalLengthBits := e.Config.GetTotalLength()
	totalLengthBytes := int(math.Ceil(float64(totalLengthBits) / 8.0))
	// Encode input based on the instructions from the attached config
	output := make([]byte, totalLengthBytes)

	// Return the output
	return output, nil
}
