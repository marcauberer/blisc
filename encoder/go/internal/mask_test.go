package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBitmaskFromRange1(t *testing.T) { // Happy path
	// Test data
	testData := []struct {
		posHigh        uint64
		posLow         uint64
		expectedResult uint64
	}{
		{
			posHigh:        13,
			posLow:         6,
			expectedResult: 0x1fc0,
		},
		{
			posHigh:        62,
			posLow:         39,
			expectedResult: 0x3fffff8000000000,
		},
		{
			posHigh:        64,
			posLow:         0,
			expectedResult: 0xffffffffffffffff,
		},
		{
			posHigh:        5,
			posLow:         12,
			expectedResult: 0xfe0,
		},
	}
	// Execute test
	for _, item := range testData {
		acualResult, actualError := createBitmaskForRange(item.posHigh, item.posLow)
		assert.Nil(t, actualError)
		assert.Equal(t, item.expectedResult, acualResult)
	}
}

func TestCreateBitmaskFromRange2(t *testing.T) { // Error
	// Test data
	posHigh := uint64(65)
	posLow := uint64(42)
	// Execute test
	actualResult, actualError := createBitmaskForRange(posHigh, posLow)
	assert.Equal(t, uint64(0), actualResult)
	assert.NotNil(t, actualError)
	assert.Equal(t, "the stated bitmask range [42,65] was invalid", actualError.Error())
}
