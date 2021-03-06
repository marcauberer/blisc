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
func (c EncoderConfig) GetTotalLength() uint {
	if len(c.Fields) == 0 {
		return 0
	}
	lastField := c.Fields[len(c.Fields)-1]
	return lastField.Pos + lastField.Len
}

func (f *EncoderConfigField) setTypeString(typeInt int) {
	switch typeInt {
	case 0:
		f.Type = "int"
	case 1:
		f.Type = "double"
	case 2:
		f.Type = "string"
	case 3:
		f.Type = "bool"
	}
}
