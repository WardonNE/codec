Supported Format
- json
- xml
- yaml
- toml

Usage
```
go get -u github.com/wardonne/codec
```

Init
```golang
codecFactory := codec.NewCodecFactory()
```

JSON Encode/Decode
```golang
// encode
codecFactory.JSONEncode(data)
// decode
codecFactory.JSONDecode(data, container)
```

XML Encode/Decode
```golang
// encode
codecFactory.XMLEncode(data)
// decode
codecFactory.XMLDecode(data, container)
```

YAML Encode/Decode
```golang
// encode
codecFactory.YAMLEncode(data)
// decode
codecFactory.YAMLDecode(data, container)
```

TOML Encode/Decode
```golang
// encode
codecFactory.TOMLEncode(data)
// decode
codecFactory.TOMLDecode(data, container)
```

Custom Codec
```golang
// register
codecFactory.RegisterCustomCodec(codec.CodecType("custom"), CustomCodec)
// load
codecFactory.Get(codec.CodecType("custom"))
```

Custom Codec Encode/Decode
```golang
// encode
codecFactory.Encode(codec.CodecType("custom"), data)
// decode
codecFactory.Decode(codec.CodecType("custom"), data, container)
```
