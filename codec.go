package codec

import (
	"fmt"
	"io"
	"sync"
)

type CodecType string

const (
	CODEC_TYPE_JSON CodecType = "json"
	CODEC_TYPE_XML  CodecType = "xml"
	CODEC_TYPE_YML  CodecType = "yml"
	CODEC_TYPE_YAML CodecType = "yaml"
	CODEC_TYPE_TOML CodecType = "toml"
)

var codecFactory *CodecFactory

type CodecFactory struct {
	codecs sync.Map
}

func init() {
	codecFactory = NewCodecFactory()
}

func NewCodecFactory() *CodecFactory {
	if codecFactory == nil {
		codecFactory = new(CodecFactory)
		codecFactory.RegisterCustomCodec(CODEC_TYPE_JSON, new(JSONCodec))
		codecFactory.RegisterCustomCodec(CODEC_TYPE_XML, new(XMLCodec))
		codecFactory.RegisterCustomCodec(CODEC_TYPE_YML, new(YAMLCodec))
		codecFactory.RegisterCustomCodec(CODEC_TYPE_YAML, new(YAMLCodec))
		codecFactory.RegisterCustomCodec(CODEC_TYPE_TOML, new(TOMLCodec))
	}
	return codecFactory
}

func (factory *CodecFactory) Get(typ CodecType) ICodec {
	codec, ok := factory.codecs.Load(typ)
	if ok {
		return codec.(ICodec)
	}
	panic(fmt.Errorf("invalid codec named `%s`", typ))
}

func (factory *CodecFactory) Encode(typ CodecType, data any) ([]byte, error) {
	return factory.Get(typ).Encode(data)
}

func (factory *CodecFactory) EncodeToString(typ CodecType, data any) (string, error) {
	return factory.Get(typ).EncodeToString(data)
}

func (factory *CodecFactory) EncodeToWriter(typ CodecType, data any, w io.Writer) error {
	return factory.Get(typ).EncodeToWriter(data, w)
}

func (factory *CodecFactory) Decode(typ CodecType, data []byte, container any) error {
	return factory.Get(typ).Decode(data, container)
}

func (factory *CodecFactory) DecodeFromString(typ CodecType, data string, container any) error {
	return factory.Get(typ).DecodeFromString(data, container)
}

func (factory *CodecFactory) DecodeFromReader(typ CodecType, r io.Reader, container any) error {
	return factory.Get(typ).DecodeFromReader(r, container)
}

func (factory *CodecFactory) JSONEncode(data any) ([]byte, error) {
	return factory.Encode(CODEC_TYPE_JSON, data)
}

func (factory *CodecFactory) JSONEncodeToString(data any) (string, error) {
	return factory.EncodeToString(CODEC_TYPE_JSON, data)
}

func (factory *CodecFactory) JSONEncodeToWriter(data any, w io.Writer) error {
	return factory.EncodeToWriter(CODEC_TYPE_JSON, data, w)
}

func (factory *CodecFactory) JSONDecode(data []byte, container any) error {
	return factory.Decode(CODEC_TYPE_JSON, data, container)
}

func (factory *CodecFactory) JSONDecodeFromString(data string, container any) error {
	return factory.DecodeFromString(CODEC_TYPE_JSON, data, container)
}

func (factory *CodecFactory) JSONDecodeFromReader(r io.Reader, container any) error {
	return factory.DecodeFromReader(CODEC_TYPE_JSON, r, container)
}

func (factory *CodecFactory) XMLEncode(data any) ([]byte, error) {
	return factory.Encode(CODEC_TYPE_XML, data)
}

func (factory *CodecFactory) XMLEncodeToString(data any) (string, error) {
	return factory.EncodeToString(CODEC_TYPE_XML, data)
}

func (factory *CodecFactory) XMLEncodeToWriter(data any, w io.Writer) error {
	return factory.EncodeToWriter(CODEC_TYPE_XML, data, w)
}

func (factory *CodecFactory) XMLDecode(data []byte, container any) error {
	return factory.Decode(CODEC_TYPE_XML, data, container)
}

func (factory *CodecFactory) XMLDecodeFromString(data string, container any) error {
	return factory.DecodeFromString(CODEC_TYPE_XML, data, container)
}

func (factory *CodecFactory) XMLDecodeFromReader(r io.Reader, container any) error {
	return factory.DecodeFromReader(CODEC_TYPE_XML, r, container)
}

func (factory *CodecFactory) YAMLEncode(data any) ([]byte, error) {
	return factory.Encode(CODEC_TYPE_YAML, data)
}

func (factory *CodecFactory) YAMLEncodeToString(data any) (string, error) {
	return factory.EncodeToString(CODEC_TYPE_YAML, data)
}

func (factory *CodecFactory) YAMLEncodeToWriter(data any, w io.Writer) error {
	return factory.EncodeToWriter(CODEC_TYPE_YAML, data, w)
}

func (factory *CodecFactory) YAMLDecode(data []byte, container any) error {
	return factory.Decode(CODEC_TYPE_YAML, data, container)
}

func (factory *CodecFactory) YAMLDecodeFromString(data string, container any) error {
	return factory.DecodeFromString(CODEC_TYPE_YAML, data, container)
}

func (factory *CodecFactory) YAMLDecodeFromReader(r io.Reader, container any) error {
	return factory.DecodeFromReader(CODEC_TYPE_YAML, r, container)
}

func (factory *CodecFactory) TOMLEncode(data any) ([]byte, error) {
	return factory.Encode(CODEC_TYPE_TOML, data)
}

func (factory *CodecFactory) TOMLEncodeToString(data any) (string, error) {
	return factory.EncodeToString(CODEC_TYPE_TOML, data)
}

func (factory *CodecFactory) TOMLEncodeToWriter(data any, w io.Writer) error {
	return factory.EncodeToWriter(CODEC_TYPE_TOML, data, w)
}

func (factory *CodecFactory) TOMLDecode(data []byte, container any) error {
	return factory.Decode(CODEC_TYPE_TOML, data, container)
}

func (factory *CodecFactory) TOMLDecodeFromString(data string, container any) error {
	return factory.DecodeFromString(CODEC_TYPE_TOML, data, container)
}

func (factory *CodecFactory) TOMLDecodeFromReader(r io.Reader, container any) error {
	return factory.DecodeFromReader(CODEC_TYPE_TOML, r, container)
}

func (factory *CodecFactory) RegisterCustomCodec(key CodecType, codec ICodec) {
	factory.codecs.Store(key, codec)
}
