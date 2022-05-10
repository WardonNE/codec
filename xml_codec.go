package codec

import (
	"encoding/xml"
	"io"
)

type XMLCodec struct{}

func (c *XMLCodec) Encode(data any) ([]byte, error) {
	return xml.Marshal(data)
}

func (c *XMLCodec) EncodeToString(data any) (string, error) {
	xmlBytes, err := c.Encode(data)
	if err != nil {
		return "", err
	}
	return string(xmlBytes), nil
}

func (c *XMLCodec) EncodeToWriter(data any, writer io.Writer) error {
	return xml.NewEncoder(writer).Encode(data)
}

func (c *XMLCodec) Decode(data []byte, container any) error {
	return xml.Unmarshal(data, container)
}

func (c *XMLCodec) DecodeFromString(data string, container any) error {
	return xml.Unmarshal([]byte(data), container)
}

func (c *XMLCodec) DecodeFromReader(reader io.Reader, container any) error {
	return xml.NewDecoder(reader).Decode(container)
}
