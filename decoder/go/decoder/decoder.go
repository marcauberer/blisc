package decoder

import (
	"clientlib-decoder/internal"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// Decoder represents a binary compression decoder
type Decoder struct {
	Config *DecoderConfig
}

// GetDecoder creates and returns a new binary compression encoder
func GetDecoder() *Decoder {
	return &Decoder{
		Config: &DecoderConfig{
			Version: DefaultEncoderConfigVersion,
		},
	}
}

// LoadConfigFromJSON parses a configuration from a JSON string and attaches it to the encoder
func (d Decoder) LoadConfigFromJSON(configBytes []byte) error {
	return json.Unmarshal(configBytes, &d.Config)
}

func (e Decoder) LoadConfigFromFile(configPath string) error {
	// Read file at file path
	// #nosec G304
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	// Parse JSON to config
	return e.LoadConfigFromJSON(content)
}

// LoadConfigFromUrl parses a configuration from a URL and attaches it to the encoder
func (d Decoder) LoadConfigFromUrl(configUrl string) error {
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
	return d.LoadConfigFromJSON(body)
}

// Decode executes the decompression process and enriches the output interface with the decoded input data
func (d Decoder) Decode(encoded []byte, output interface{}) error {
	// Abort if no config is set
	if d.Config == nil {
		return errors.New("no encoding configuration was set")
	}
	// Abort if encoded input is not long enough
	if uint(len(encoded))*8 < d.Config.GetTotalLength() {
		return errors.New("encoded byte array length does not match the configuration")
	}
	// Abort if the passed output interface is no struct
	pointToStruct := reflect.ValueOf(output).Elem()
	if pointToStruct.Kind() != reflect.Struct {
		return errors.New("the passed interface is no struct")
	}

	for _, b := range encoded {
		fmt.Printf("%08b ", b)
	}
	fmt.Println()

	// Fill struct
	for _, field := range d.Config.Fields {
		pointToField := pointToStruct.FieldByName(strings.Title(field.Name))
		// Abort when actual field type does not match config field type
		switch field.Type {
		case "int":
			if pointToField.Type().Name() != "int" {
				return fmt.Errorf("expected int, but got '%s' for field '%s'", field.Type, field.Name)
			}
		case "double":
			if pointToField.Type().Name() != "float32" {
				return fmt.Errorf("expected float32, but got '%s' for field '%s'", field.Type, field.Name)
			}
		case "string":
			if pointToField.Type().Name() != "string" {
				return fmt.Errorf("expected string, but got '%s' for field '%s'", field.Type, field.Name)
			}
		case "bool":
			if pointToField.Type().Name() != "bool" {
				return fmt.Errorf("expected bool, but got '%s' for field '%s'", field.Type, field.Name)
			}
		}
		// Calculate left and right border of considered bitframe
		leftPos := field.Pos
		rightPos := field.Pos + field.Len
		firstByte := leftPos / 8
		lastByte := rightPos / 8
		if rightPos%8 != 0 {
			lastByte++
		}
		// Get affected bytes
		affectedBytes := encoded[firstByte:lastByte]
		// Extract number from affected bytes
		value := uint64(0)
		startPosInFirstByte := leftPos % 8
		endPosInLastByte := rightPos % 8
		for i, b := range affectedBytes {
			// Determine relevant bits in the current byte
			startPosInByte := 0
			endPosInByte := 8
			if i == 0 { // First byte
				startPosInByte = int(startPosInFirstByte)
			}
			if i == len(affectedBytes)-1 { // Last byte
				endPosInByte = int(endPosInLastByte)
			}
			// Create bitmask
			bitmask, err := internal.CreateBitmask8ForRange(uint64(8-startPosInByte), uint64(8-endPosInByte))
			if err != nil {
				panic(err)
			}
			// Mask the current byte and add it to the output value
			value |= uint64(b) & uint64(bitmask)
			if endPosInByte < 8 {
				value = value >> (8 - endPosInByte)
			} else {
				value = value << 8
			}
		}
		// Apply mul and bias in reverse
		switch field.Type {
		case "int":

		case "double":
			doubleValue := float64(value)
			doubleValue /= float64(field.Mul)
			doubleValue -= float64(field.Bias)
			pointToField.SetFloat(doubleValue)
		case "string":

		case "bool":

		}
	}
	return nil
}
