package internal

import (
	"fmt"
)

func CreateBitmask64ForRange(posHigh, posLow uint64) (uint64, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh = posHigh ^ posLow
		posLow = posLow ^ posHigh
		posHigh = posHigh ^ posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 64 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint64((1 << posHigh) - 1)
	mask = mask ^ ((1 << posLow) - 1)
	return mask, nil
}

func CreateBitmask32ForRange(posHigh, posLow uint64) (uint32, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh = posHigh ^ posLow
		posLow = posLow ^ posHigh
		posHigh = posHigh ^ posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 32 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint32((1 << posHigh) - 1)
	mask = mask ^ ((1 << posLow) - 1)
	return mask, nil
}

func CreateBitmask16ForRange(posHigh, posLow uint64) (uint16, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh = posHigh ^ posLow
		posLow = posLow ^ posHigh
		posHigh = posHigh ^ posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 16 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint16((1 << posHigh) - 1)
	mask = mask ^ ((1 << posLow) - 1)
	return mask, nil
}

func CreateBitmask8ForRange(posHigh, posLow uint64) (uint8, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh = posHigh ^ posLow
		posLow = posLow ^ posHigh
		posHigh = posHigh ^ posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 8 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint8((1 << posHigh) - 1) // e.g. for posHigh = 6: 0011 1111
	mask = mask ^ ((1 << posLow) - 1) // e.g. for losLow = 2:  0011 1100
	return mask, nil
}
