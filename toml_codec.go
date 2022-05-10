package codec

import (
	"io"

	"github.com/pelletier/go-toml/v2"
)

type TOMLCodec struct{}

func (c *TOMLCodec) Encode(data any) ([]byte, error) {
	return toml.Marshal(data)
}

func (c *TOMLCodec) EncodeToString(data any) (string, error) {
	tomlBytes, err := c.Encode(data)
	if err != nil {
		return "", err
	}
	return string(tomlBytes), nil
}

func (c *TOMLCodec) EncodeToWriter(data any, writer io.Writer) error {
	return toml.NewEncoder(writer).Encode(data)
}

func (c *TOMLCodec) Decode(data []byte, container any) error {
	return toml.Unmarshal(data, container)
}

func (c *TOMLCodec) DecodeFromString(data string, container any) error {
	return toml.Unmarshal([]byte(data), container)
}

func (c *TOMLCodec) DecodeFromReader(reader io.Reader, container any) error {
	return toml.NewDecoder(reader).Decode(container)
}
