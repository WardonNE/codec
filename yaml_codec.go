package codec

import (
	"io"

	"gopkg.in/yaml.v3"
)

type YAMLCodec struct {
}

func (c *YAMLCodec) Encode(data any) ([]byte, error) {
	return yaml.Marshal(data)
}

func (c *YAMLCodec) EncodeToString(data any) (string, error) {
	yamlBytes, err := c.Encode(data)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

func (c *YAMLCodec) EncodeToWriter(data any, writer io.Writer) error {
	return yaml.NewEncoder(writer).Encode(data)
}

func (c *YAMLCodec) Decode(data []byte, container any) error {
	return yaml.Unmarshal(data, container)
}

func (c *YAMLCodec) DecodeFromString(data string, container any) error {
	return yaml.Unmarshal([]byte(data), container)
}

func (c *YAMLCodec) DecodeFromReader(reader io.Reader, container any) error {
	return yaml.NewDecoder(reader).Decode(container)
}
