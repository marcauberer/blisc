package encoder

import (
	"clientlib-go/internal/bitarray"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func (e Encoder) LoadConfigFromFile(configPath string) error {
	// Read file at file path
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	// Parse JSON to config
	e.LoadConfigFromJSON(content)
	return nil
}

// LoadConfigFromUrl parses a configuration from a URL and attaches it to the encoder
func (e Encoder) LoadConfigFromUrl(configUrl string) error {
	// Execute web request to get JSON from url
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
	e.LoadConfigFromJSON(body)
	return nil
}

// Encode executes the compression process and returns the output as byte array
func (e Encoder) Encode(input interface{}) ([]byte, error) {
	// Abort if no config is set
	if e.Config == nil {
		return []byte{}, errors.New("no encoding configuration was set")
	}
	// Calculate total length of output
	totalLengthBits := e.Config.GetTotalLength()
	// Encode input based on the instructions from the attached config
	output := bitarray.NewBitArray(uint64(totalLengthBits))
	v := reflect.ValueOf(input).Elem()
	for _, field := range e.Config.Fields {
		value := v.FieldByName(field.Name)
		switch value.Type().Name() { // Switch by field type name
		case "int":
			intValue := value.Int()
			intValue += int64(field.Bias)
			intValue = int64(float64(intValue) * field.Mul)
			fmt.Println("Int:", intValue)
			output.InsertUInt64At(uint64(field.Pos), uint64(intValue))
			fmt.Println("Stringified:", output.String())
		case "float32": // Double field
			doubleValue := value.Float()
			doubleValue += float64(field.Bias)
			doubleValue *= field.Mul
			intValue := int64(doubleValue)
			fmt.Println("Double:", intValue)
			output.InsertUInt64At(uint64(field.Pos), uint64(intValue))
			fmt.Println("Stringified:", output.String())
		case "string": // String field

		case "bool": // Boolean field

		}
	}
	// Return the output
	return output.ToByteArray(), nil
}
