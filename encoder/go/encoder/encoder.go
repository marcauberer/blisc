package encoder

import (
	"clientlib-go/internal"
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

func (e Encoder) LoadConfigFromFile(configPath string) error {
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
	// ToDo: Sort fields after position
	for _, field := range e.Config.Fields {
		value := v.FieldByName(field.Name)
		switch value.Type().Name() { // Switch by field type name
		case reflect.Int.String():
			intValue := value.Int()
			intValue += int64(field.Bias)
			intValue = int64(math.Round(float64(intValue) * field.Mul))
			fmt.Println("Int:", intValue)
			err := output.PushUInt64(uint64(intValue), uint64(field.Len))
			if err != nil {
				return []byte{}, err
			}
			fmt.Println("Stringified:", output.ToString())
		case reflect.Float32.String(): // Double field
			doubleValue := value.Float()
			doubleValue += float64(field.Bias)
			doubleValue *= field.Mul
			intValue := uint64(doubleValue)
			fmt.Println("Double:", intValue)
			err := output.PushUInt64(intValue, uint64(field.Len))
			if err != nil {
				return []byte{}, err
			}
			fmt.Println("Stringified:", output.ToString())
		case reflect.String.String(): // String field

		case reflect.Bool.String(): // Boolean field
			boolValue := value.Bool()
			intValue := uint64(0)
			if boolValue {
				intValue = 1
			}
			fmt.Println("Bool:", intValue)
			err := output.PushUInt64(intValue, uint64(field.Len))
			if err != nil {
				return []byte{}, err
			}
			fmt.Println("Stringified:", output.ToString())
		}
	}
	// Write last buffer contents to stream
	output.Conclude()
	// Return the output
	return output.ToByteArray(), nil
}
