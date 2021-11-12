package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) { // Fully functional test
	// Execute test
	output := CreateOutput(7)
	assert.Nil(t, output.PushUInt64(5, 3))      // Value 5, length 3
	assert.Nil(t, output.PushUInt64(12, 5))     // Value 12, length 5
	assert.Nil(t, output.PushUInt64(3502, 13))  // Value 3502, length 13
	assert.Nil(t, output.PushUInt64(1, 1))      // Value 1, length 1
	assert.Nil(t, output.PushUInt64(10001, 16)) // Value 10001, length 16
	assert.Nil(t, output.PushUInt64(127, 7))    // Value 127, length 7
	output.Conclude()
	// Assert
	assert.Equal(t, []byte{0xac, 0x6d, 0x74, 0x9c, 0x44, 0xf8, 0x00}, output.ToByteArray())
	assert.Equal(t, "10101100 01101101 01110100 10011100 01000100 11111000 00000000", output.ToString())
}
