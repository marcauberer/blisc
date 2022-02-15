# Client Config Binary Spec

## Structure
- Version spec major (int / 4 bytes)
- Version spec minor (int / 4 bytes)
- Field spec count (int / 4 bytes)
- Field 1 name (terminated with \00) (arbitrary length)
- Field 1 type (1 byte)
- Field 1 position (int / 4 bytes)
- Field 1 length (int / 4 bytes)
- Field 1 bias (int / 4 bytes)
- Field 1 mul (double / 8 bytes)
- ... (more fields)

## Types
- `int`: 0
- `double`: 1
- `string`: 2
- `bool`: 3