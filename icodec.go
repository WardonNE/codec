package codec

import "io"

type ICodec interface {
	Encode(data any) ([]byte, error)
	EncodeToString(data any) (string, error)
	EncodeToWriter(data any, writer io.Writer) error
	Decode(data []byte, container any) error
	DecodeFromString(data string, container any) error
	DecodeFromReader(reader io.Reader, container any) error
}
