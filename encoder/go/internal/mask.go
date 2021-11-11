package internal

import (
	"fmt"
)

func createBitmaskForRange(posHigh, posLow uint64) (uint64, error) {
	//fmt.Printf("[%d,%d]\n", posLow, posHigh)
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
