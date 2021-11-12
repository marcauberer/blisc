package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------- CreateBitmask64FromRange ----------------------------------------------------------

func TestCreateBitmask64FromRange1(t *testing.T) { // Happy path
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
		acualResult, actualError := createBitmask64ForRange(item.posHigh, item.posLow)
		assert.Nil(t, actualError)
		assert.Equal(t, item.expectedResult, acualResult)
	}
}

func TestCreateBitmask64FromRange2(t *testing.T) { // Error
	// Test data
	posHigh := uint64(65)
	posLow := uint64(42)
	// Execute test
	actualResult, actualError := createBitmask64ForRange(posHigh, posLow)
	assert.Equal(t, uint64(0), actualResult)
	assert.NotNil(t, actualError)
	assert.Equal(t, "the stated bitmask range [42,65] was invalid", actualError.Error())
}

// ------------------------------------------------------------- CreateBitmask32FromRange ----------------------------------------------------------

func TestCreateBitmask32FromRange1(t *testing.T) { // Happy path
	// Test data
	testData := []struct {
		posHigh        uint64
		posLow         uint64
		expectedResult uint32
	}{
		{
			posHigh:        13,
			posLow:         6,
			expectedResult: 0x1fc0,
		},
		{
			posHigh:        30,
			posLow:         19,
			expectedResult: 0x3ff80000,
		},
		{
			posHigh:        32,
			posLow:         0,
			expectedResult: 0xffffffff,
		},
		{
			posHigh:        5,
			posLow:         12,
			expectedResult: 0xfe0,
		},
	}
	// Execute test
	for _, item := range testData {
		acualResult, actualError := createBitmask32ForRange(item.posHigh, item.posLow)
		assert.Nil(t, actualError)
		assert.Equal(t, item.expectedResult, acualResult)
	}
}

func TestCreateBitmask32FromRange2(t *testing.T) { // Error
	// Test data
	posHigh := uint64(33)
	posLow := uint64(22)
	// Execute test
	actualResult, actualError := createBitmask32ForRange(posHigh, posLow)
	assert.Equal(t, uint32(0), actualResult)
	assert.NotNil(t, actualError)
	assert.Equal(t, "the stated bitmask range [22,33] was invalid", actualError.Error())
}

// ------------------------------------------------------------- CreateBitmask16FromRange ----------------------------------------------------------

func TestCreateBitmask16FromRange1(t *testing.T) { // Happy path
	// Test data
	testData := []struct {
		posHigh        uint64
		posLow         uint64
		expectedResult uint16
	}{
		{
			posHigh:        13,
			posLow:         6,
			expectedResult: 0x1fc0,
		},
		{
			posHigh:        14,
			posLow:         9,
			expectedResult: 0x3e00,
		},
		{
			posHigh:        16,
			posLow:         0,
			expectedResult: 0xffff,
		},
		{
			posHigh:        5,
			posLow:         12,
			expectedResult: 0xfe0,
		},
	}
	// Execute test
	for _, item := range testData {
		acualResult, actualError := createBitmask16ForRange(item.posHigh, item.posLow)
		assert.Nil(t, actualError)
		assert.Equal(t, item.expectedResult, acualResult)
	}
}

func TestCreateBitmask16FromRange2(t *testing.T) { // Error
	// Test data
	posHigh := uint64(17)
	posLow := uint64(12)
	// Execute test
	actualResult, actualError := createBitmask16ForRange(posHigh, posLow)
	assert.Equal(t, uint16(0), actualResult)
	assert.NotNil(t, actualError)
	assert.Equal(t, "the stated bitmask range [12,17] was invalid", actualError.Error())
}

// ------------------------------------------------------------- CreateBitmask8FromRange -----------------------------------------------------------

func TestCreateBitmask8FromRange1(t *testing.T) { // Happy path
	// Test data
	testData := []struct {
		posHigh        uint64
		posLow         uint64
		expectedResult uint8
	}{
		{
			posHigh:        4,
			posLow:         3,
			expectedResult: 0x8,
		},
		{
			posHigh:        5,
			posLow:         1,
			expectedResult: 0x1e,
		},
		{
			posHigh:        8,
			posLow:         0,
			expectedResult: 0xff,
		},
		{
			posHigh:        4,
			posLow:         7,
			expectedResult: 0x70,
		},
	}
	// Execute test
	for _, item := range testData {
		acualResult, actualError := createBitmask8ForRange(item.posHigh, item.posLow)
		assert.Nil(t, actualError)
		assert.Equal(t, item.expectedResult, acualResult)
	}
}

func TestCreateBitmask8FromRange2(t *testing.T) { // Error
	// Test data
	posHigh := uint64(9)
	posLow := uint64(6)
	// Execute test
	actualResult, actualError := createBitmask8ForRange(posHigh, posLow)
	assert.Equal(t, uint8(0), actualResult)
	assert.NotNil(t, actualError)
	assert.Equal(t, "the stated bitmask range [6,9] was invalid", actualError.Error())
}
