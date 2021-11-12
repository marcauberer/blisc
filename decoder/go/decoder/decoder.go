package decoder

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
func (d Decoder) Decode(encoded []byte, output interface{}) {

}
