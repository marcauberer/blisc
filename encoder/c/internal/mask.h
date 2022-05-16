#pragma once

typedef unsigned long long mask64;
typedef unsigned long mask32;
typedef unsigned int mask16;
typedef unsigned short mask8;
typedef unsigned short mask1;
typedef unsigned long long mask_pos;

mask64 createBitmask64ForRange(mask_pos posHigh, mask_pos posLow);
mask32 createBitmask32ForRange(mask_pos posHigh, mask_pos posLow);
mask16 createBitmask16ForRange(mask_pos posHigh, mask_pos posLow);
mask8 createBitmask8ForRange(mask_pos posHigh, mask_pos posLow);
mask1 createBitmask1ForRange(mask_pos posHigh, mask_pos posLow);