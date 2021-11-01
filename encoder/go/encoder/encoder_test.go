package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode1(t *testing.T) { // Happy path
	// Test data
	pmData := struct {
		pm10        float32
		pm2_5       float32
		temperature float32
		humidity    float32
		pressure    float32
	}{
		pm10:        12.43,
		pm2_5:       6.14,
		temperature: 25.124,
		humidity:    78.01,
		pressure:    1001.923,
	}
	// Execute test
	var encoder Encoder
	encoder.Config = &EncoderConfig{
		Version: DefaultEncoderConfigVersion,
		Fields: []EncoderConfigField{
			{
				Name: "pm10",
				Type: FieldTypeDouble,
				Pos:  0,
				Len:  14,
				Bias: 0,
				Mul:  1,
			},
			{
				Name: "pm2_5",
				Type: FieldTypeDouble,
				Pos:  1,
				Len:  14,
				Bias: 0,
				Mul:  1,
			},
			{
				Name: "temperature",
				Type: FieldTypeDouble,
				Pos:  27,
				Len:  10,
				Bias: 0,
				Mul:  1,
			},
			{
				Name: "humidity",
				Type: FieldTypeDouble,
				Pos:  37,
				Len:  10,
				Bias: 0,
				Mul:  1,
			},
			{
				Name: "pressure",
				Type: FieldTypeDouble,
				Pos:  47,
				Len:  11,
				Bias: -87000,
				Mul:  1,
			},
		},
	}
	result, err := encoder.Encode(&pmData)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, []byte{}, result)
}

func TestEncode2(t *testing.T) { // Missing config
	// Test data
	pmData := struct{}{}
	// Execute test
	var encoder Encoder
	result, err := encoder.Encode(&pmData)
	assert.NotNil(t, err)
	assert.Equal(t, "no encoding configuration was set", err.Error())
	assert.Equal(t, []byte{}, result)
}
