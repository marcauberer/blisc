package main

import (
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"os"
)

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

func (f EncoderConfigField) typeToInt() uint8 {
	if f.Type == "int" {
		return 0
	}
	if f.Type == "double" {
		return 1
	}
	if f.Type == "string" {
		return 2
	}
	return 3
}

const (
	FieldTypeInteger = "int"
	FieldTypeDouble  = "double"
	FieldTypeString  = "string"
	FieldTypeBool    = "bool"
)

func loadConfigFromFile(config *EncoderConfig, configPath string) error {
	// Read file at file path
	// #nosec G304
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	// Parse JSON to config
	return json.Unmarshal(content, &config)
}

func writeConfigToBinary(config *EncoderConfig, binaryFile *os.File) {
	// Write version spec major
	binary.Write(binaryFile, binary.BigEndian, uint32(config.Version.Major))
	// Write version spec minor
	binary.Write(binaryFile, binary.BigEndian, uint32(config.Version.Minor))
	// Write field count
	binary.Write(binaryFile, binary.BigEndian, uint32(len(config.Fields)))
	// Write fields
	for _, field := range config.Fields {
		// Write field name
		nameNullTerminated := field.Name + string(rune(0))
		binary.Write(binaryFile, binary.BigEndian, []byte(nameNullTerminated))
		// Write field type
		binary.Write(binaryFile, binary.BigEndian, field.typeToInt())
		// Write position
		binary.Write(binaryFile, binary.BigEndian, uint32(field.Pos))
		// Write length
		binary.Write(binaryFile, binary.BigEndian, uint32(field.Len))
		// Write bias
		binary.Write(binaryFile, binary.BigEndian, uint32(field.Bias))
		// Write multiplicator
		binary.Write(binaryFile, binary.BigEndian, field.Mul)
	}
}

func main() {
	config := EncoderConfig{}

	// Parse JSON file
	loadConfigFromFile(&config, "../../config/client-config.json")

	// Write binary file
	binary, _ := os.Create("../../config/client-config.bin")
	writeConfigToBinary(&config, binary)
	binary.Close()
}
