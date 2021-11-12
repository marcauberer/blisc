package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) { // Fully functional test
	// Execute test
	output := CreateOutput(8)
	assert.Nil(t, output.PushUInt64(0, 1))
	output.Conclude()
	// Assert
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, output.ToByteArray())
	assert.Equal(t, "00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000", output.ToString())
}
