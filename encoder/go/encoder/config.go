package encoder

// EncoderConfig represents the client encoding instruction configuration
type EncoderConfig struct {
	Version EncoderConfigVersion
	Fields  []EncoderConfigField
}

// EncoderConfigVersion represents a specific config protocol version
type EncoderConfigVersion struct {
	Major int
	Minor int
}

// DefaultEncoderConfigVersion marks the maximum supported config version
var DefaultEncoderConfigVersion = EncoderConfigVersion{Major: 1, Minor: 0}

// EncoderConfigField represents a instruction set for one particular data field
type EncoderConfigField struct {
	Name string
	Pos  int
	Len  int
	Bias int
	Mul  int
}

// GetTotalLength returns the total length of the encoding output in bits
func (c EncoderConfig) GetTotalLength() int {
	if len(c.Fields) == 0 {
		return 0
	}
	lastField := c.Fields[len(c.Fields)-1]
	return lastField.Pos + lastField.Len
}
