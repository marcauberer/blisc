#include "mask.h"

mask64 createBitmask64ForRange(mask_pos posHigh, mask_pos posLow) {
    // If posLow > posHigh, swap the values with the XOR swap algorithm
	if (posLow > posHigh) {
		posHigh ^= posLow;
		posLow ^= posHigh;
		posHigh ^= posLow;
	}
    // Check if posHigh exceeds range
    if (posHigh > 64) return -1;
    // Create bitmask
    mask64 m = (1 << posHigh) - 1;
    m ^= (1 << posLow) - 1;
    return m;
}

mask32 createBitmask32ForRange(mask_pos posHigh, mask_pos posLow) {
    // If posLow > posHigh, swap the values with the XOR swap algorithm
	if (posLow > posHigh) {
		posHigh ^= posLow;
		posLow ^= posHigh;
		posHigh ^= posLow;
	}
    // Check if posHigh exceeds range
    if (posHigh > 32) return -1;
    // Create bitmask
    mask32 m = (1 << posHigh) - 1;
    m ^= (1 << posLow) - 1;
    return m;
}

mask16 createBitmask16ForRange(mask_pos posHigh, mask_pos posLow) {
    // If posLow > posHigh, swap the values with the XOR swap algorithm
	if (posLow > posHigh) {
		posHigh ^= posLow;
		posLow ^= posHigh;
		posHigh ^= posLow;
	}
    // Check if posHigh exceeds range
    if (posHigh > 16) return -1;
    // Create bitmask
    mask16 m = (1 << posHigh) - 1;
    m ^= (1 << posLow) - 1;
    return m;
}

mask8 createBitmask8ForRange(mask_pos posHigh, mask_pos posLow) {
    // If posLow > posHigh, swap the values with the XOR swap algorithm
	if (posLow > posHigh) {
		posHigh ^= posLow;
		posLow ^= posHigh;
		posHigh ^= posLow;
	}
    // Check if posHigh exceeds range
    if (posHigh > 8) return -1;
    // Create bitmask
    mask8 m = (1 << posHigh) - 1;
    m ^= (1 << posLow) - 1;
    return m;
}