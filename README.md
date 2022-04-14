# ClientLib

Idee von Julien: Wurzelfunktion

## ToDo
- Go Benchmark für Größe pro DS
- Encoder in C schreiben
- Encoding-Zeit vergleichen gegen Zip / zlib / brotli (Led blinken lassen und Zeit messen)

## Limitations
### C Implementation
- Only up to 100 fields
- Field names only up to 99 chars

## Comparison with ZIP
- Encoded record sizes are always the same with the Client Lib
- For the pm example data, this is a size of 8 bytes

Client Lib with 1000 Records:     8000 bytes ≈ 7.8 kiB
JSON ZIP with 1000 Records:       28009 bytes ≈ 27.4 kiB
BSON ZIP with 1000 Records:       -

© Marc Auberer 2021