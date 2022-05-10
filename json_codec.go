package codec

import (
	"encoding/json"
	"io"
)

type JSONCodec struct{}

func (c *JSONCodec) Encode(data any) ([]byte, error) {
	return json.Marshal(data)
}

func (c *JSONCodec) EncodeToString(data any) (string, error) {
	jsonBytes, err := c.Encode(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (c *JSONCodec) EncodeToWriter(data any, writer io.Writer) error {
	return json.NewEncoder(writer).Encode(data)
}

func (c *JSONCodec) Decode(data []byte, container any) error {
	return json.Unmarshal(data, container)
}

func (c *JSONCodec) DecodeFromString(data string, container any) error {
	return json.Unmarshal([]byte(data), container)
}

func (c *JSONCodec) DecodeFromReader(reader io.Reader, container any) error {
	return json.NewDecoder(reader).Decode(container)
}
