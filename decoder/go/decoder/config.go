package decoder

// DecoderConfig represents the client encoding instruction configuration
type DecoderConfig struct {
	Version DecoderConfigVersion
	Fields  []DecoderConfigField
}

// EncoderConfigVersion represents a specific config protocol version
type DecoderConfigVersion struct {
	Major int
	Minor int
}

// DefaultEncoderConfigVersion marks the maximum supported config version
var DefaultEncoderConfigVersion = DecoderConfigVersion{Major: 1, Minor: 0}

// EncoderConfigField represents a instruction set for one particular data field
type DecoderConfigField struct {
	Name string
	Type string
	Pos  uint
	Len  uint
	Bias int
	Mul  float64
}

const (
	FieldTypeInteger = "int"
	FieldTypeDouble  = "double"
	FieldTypeString  = "string"
	FieldTypeBool    = "bool"
)

// GetTotalLength returns the total length of the encoding output in bits
func (c DecoderConfig) GetTotalLength() uint {
	if len(c.Fields) == 0 {
		return 0
	}
	lastField := c.Fields[len(c.Fields)-1]
	return lastField.Pos + lastField.Len
}
