package jsonx

import (
	"bytes"
	"encoding/json"
)

type encodingJsonAdapter struct {
	config *config
}

func (a *encodingJsonAdapter) Marshal(v any) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(a.config.escapeHTML)
	err := encoder.Encode(v)
	return buf.Bytes(), err
}

func (a *encodingJsonAdapter) Unmarshal(data []byte, v any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if a.config.useNumber {
		decoder.UseNumber()
	}
	return decoder.Decode(v)
}

func (a *encodingJsonAdapter) MarshalString(v any) (string, error) {
	bytes, err := a.Marshal(v)
	return string(bytes), err
}

func (a *encodingJsonAdapter) UnmarshalString(data string, v any) error {
	return a.Unmarshal([]byte(data), v)
}

func (a *encodingJsonAdapter) ToString(v any) string {
	str, _ := a.MarshalString(v)
	return str
}

func (a *encodingJsonAdapter) ToByte(v any) []byte {
	bytes, _ := a.Marshal(v)
	return bytes
}

func newEncodingJsonAdapter(config *config) JsonAdapter {
	return &encodingJsonAdapter{config: config}
}
