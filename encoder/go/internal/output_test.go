package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) { // Fully functional test
	output := CreateOutput(8)
	assert.Nil(t, output.PushUInt64(0, 0))
}
