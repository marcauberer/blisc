package encoder

import (
	"clientlib-encoder/internal"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
)

// Encoder represents a binary compression encoder
type Encoder struct {
	Config *EncoderConfig
}

// GetEncoder creates and returns a new binary compression encoder
func GetEncoder() *Encoder {
	return &Encoder{
		Config: &EncoderConfig{
			Version: DefaultEncoderConfigVersion,
		},
	}
}

// LoadConfigFromJSON parses a configuration from a JSON string and attaches it to the encoder
func (e Encoder) LoadConfigFromJSON(configBytes []byte) error {
	return json.Unmarshal(configBytes, &e.Config)
}

// LoadConfigFromBinary parses a configuration from a binary and attaches it to the encoder
func (e Encoder) LoadConfigFromBinary(configBytes []byte) error {
	byteCounter := 0
	// Major version
	e.Config.Version.Major = int(binary.LittleEndian.Uint32(configBytes[0:4])) // 4 bytes
	byteCounter += 4
	// Minor version
	e.Config.Version.Minor = int(binary.LittleEndian.Uint32(configBytes[4:8])) // 4 bytes
	byteCounter += 4
	// Field count
	fieldCount := binary.LittleEndian.Uint32(configBytes[8:12]) // 4 bytes
	byteCounter += 4
	for i := 0; i < int(fieldCount); i++ {
		var field EncoderConfigField

		// Field name
		for configBytes[byteCounter] != '\x00' {
			field.Name += string(configBytes[byteCounter])
			byteCounter++
		}
		byteCounter++

		// Field type
		fieldType := int(configBytes[byteCounter]) // 1 byte
		field.setTypeString(fieldType)
		byteCounter += 1

		// Field pos
		field.Pos = uint(binary.LittleEndian.Uint32(configBytes[byteCounter : byteCounter+4])) // 4 bytes
		byteCounter += 4

		// Field length
		field.Len = uint(binary.LittleEndian.Uint32(configBytes[byteCounter : byteCounter+4])) // 4 bytes
		byteCounter += 4

		// Field bias
		field.Bias = int(binary.LittleEndian.Uint32(configBytes[byteCounter : byteCounter+4])) // 4 bytes
		byteCounter += 4

		// Field mul
		field.Mul = math.Float64frombits(binary.LittleEndian.Uint64(configBytes[byteCounter : byteCounter+8])) // 8 bytes
		byteCounter += 8

		// Append field to the list of fields
		e.Config.Fields = append(e.Config.Fields, field)
	}
	return nil
}

func (e Encoder) LoadConfigFromFile(configPath string) error {
	// Read file at file path
	// #nosec G304
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	// If it is valid JSON, parse the JSON. Otherwise try to read it as binary
	if json.Valid(content) {
		return e.LoadConfigFromJSON(content)
	}
	return e.LoadConfigFromBinary(content)
}

// LoadConfigFromUrl parses a configuration from a URL and attaches it to the encoder
func (e Encoder) LoadConfigFromUrl(configUrl string) error {
	// Execute web request to get JSON from url
	// #nosec G107
	response, err := http.Get(configUrl)
	if err != nil {
		return err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Parse JSON to config
	return e.LoadConfigFromJSON(body)
}

// Encode executes the compression process and returns the output as byte array
func (e Encoder) Encode(input interface{}) ([]byte, error) {
	// Abort if no config is set
	if e.Config == nil {
		return []byte{}, errors.New("no encoding configuration was set")
	}
	// Calculate total length of output
	totalLengthBits := e.Config.GetTotalLength()
	totalLengthBytes := uint64(totalLengthBits / 8)
	if totalLengthBits%8 > 0 {
		totalLengthBytes++
	}
	// Encode input based on the instructions from the attached config
	output := internal.CreateOutput(totalLengthBytes)
	v := reflect.ValueOf(input).Elem()
	for _, field := range e.Config.Fields {
		value := v.FieldByName(field.Name)
		switch value.Type().Name() { // Switch by field type name
		case reflect.Int.String():
			// Check type of field config matches the actual type
			if field.Type != "int" {
				return []byte{}, fmt.Errorf("expected int, but got '%s' for field '%s'", field.Type, field.Name)
			}
			// Apply bias and mul
			intValue := value.Int()
			intValue += int64(field.Bias)
			intValue = int64(math.Round(float64(intValue) * field.Mul))
			fmt.Println("Int:", intValue)
			// Push to output byte array
			err := output.PushUInt64(uint64(intValue), uint64(field.Len))
			if err != nil {
				return []byte{}, err
			}
			fmt.Println("Stringified:", output.ToString())
		case reflect.Float32.String(): // Double field
			// Check type of field config matches the actual type
			if field.Type != "double" {
				return []byte{}, fmt.Errorf("expected double, but got '%s' for field '%s'", field.Type, field.Name)
			}
			// Apply bias and mul
			doubleValue := value.Float()
			doubleValue += float64(field.Bias)
			doubleValue *= field.Mul
			intValue := uint64(doubleValue)
			fmt.Println("Double:", intValue)
			// Push to output byte array
			err := output.PushUInt64(intValue, uint64(field.Len))
			if err != nil {
				return []byte{}, err
			}
			fmt.Println("Stringified:", output.ToString())
		case reflect.String.String(): // String field
			// Check type of field config matches the actual type
			if field.Type != "string" {
				return []byte{}, fmt.Errorf("expected string, but got '%s' for field '%s'", field.Type, field.Name)
			}
		case reflect.Bool.String(): // Boolean field
			// Check type of field config matches the actual type
			if field.Type != "bool" {
				return []byte{}, fmt.Errorf("expected bool, but got '%s' for field '%s'", field.Type, field.Name)
			}
			// Apply bias and mul
			boolValue := value.Bool()
			intValue := uint64(0)
			if boolValue {
				intValue = 1
			}
			fmt.Println("Bool:", intValue)
			// Push to output byte array
			err := output.PushUInt64(intValue, uint64(field.Len))
			if err != nil {
				return []byte{}, err
			}
			fmt.Println("Stringified:", output.ToString())
		}
	}
	// Flush the content of the last buffer into the stream
	output.Conclude()
	// Return the output
	return output.ToByteArray(), nil
}
