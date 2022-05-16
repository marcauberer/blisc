package internal

import "fmt"

func createBitmask64ForRange(posHigh, posLow uint64) (uint64, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh ^= posLow
		posLow ^= posHigh
		posHigh ^= posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 64 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint64((1 << posHigh) - 1)
	mask ^= (1 << posLow) - 1
	return mask, nil
}

func createBitmask32ForRange(posHigh, posLow uint64) (uint32, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh ^= posLow
		posLow ^= posHigh
		posHigh ^= posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 32 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint32((1 << posHigh) - 1)
	mask ^= (1 << posLow) - 1
	return mask, nil
}

func createBitmask16ForRange(posHigh, posLow uint64) (uint16, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh ^= posLow
		posLow ^= posHigh
		posHigh ^= posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 16 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint16((1 << posHigh) - 1)
	mask ^= (1 << posLow) - 1
	return mask, nil
}

func createBitmask8ForRange(posHigh, posLow uint64) (uint8, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh ^= posLow
		posLow ^= posHigh
		posHigh ^= posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 8 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint8((1 << posHigh) - 1)
	mask ^= (1 << posLow) - 1
	return mask, nil
}

func createBitmask1ForRange(posHigh, posLow uint64) (uint8, error) {
	// If posLow > posHigh, swap the values with the XOR swap algorithm
	if posLow > posHigh {
		posHigh ^= posLow
		posLow ^= posHigh
		posHigh ^= posLow
	}
	// Check if posHigh exceeds range
	if posHigh > 1 {
		return 0, fmt.Errorf("the stated bitmask range [%d,%d] was invalid", posLow, posHigh)
	}
	// Create bitmask
	mask := uint8((1 << posHigh) - 1)
	mask ^= (1 << posLow) - 1
	return mask, nil
}
