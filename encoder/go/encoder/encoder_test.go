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
	assert.Equal(t, "no encoding config attached", err.Error())
	assert.Equal(t, []byte{}, result)
}
