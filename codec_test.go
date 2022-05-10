package codec

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testUser struct {
	ID      int
	Name    string
	Country struct {
		ID   int
		Name string
	}
}

var user = &testUser{
	ID:   1,
	Name: "testUser",
	Country: struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "testCountry",
	},
}

var userJson = `{"ID":1,"Name":"testUser","Country":{"ID":1,"Name":"testCountry"}}`
var userXml = `<testUser><ID>1</ID><Name>testUser</Name><Country><ID>1</ID><Name>testCountry</Name></Country></testUser>`
var userYaml = `id: 1
name: testUser
country:
    id: 1
    name: testCountry
`
var userToml = `ID = 1
Name = 'testUser'
[Country]
ID = 1
Name = 'testCountry'

`

func TestCodecFactoryGet(t *testing.T) {
	codecFactory := NewCodecFactory()
	fmt.Println(codecFactory)
	assert.Equal(t, new(JSONCodec), codecFactory.Get(CODEC_TYPE_JSON).(*JSONCodec))
	assert.Equal(t, new(XMLCodec), codecFactory.Get(CODEC_TYPE_XML).(*XMLCodec))
	assert.Equal(t, new(YAMLCodec), codecFactory.Get(CODEC_TYPE_YML).(*YAMLCodec))
	assert.Equal(t, new(YAMLCodec), codecFactory.Get(CODEC_TYPE_YAML).(*YAMLCodec))
	assert.Equal(t, new(TOMLCodec), codecFactory.Get(CODEC_TYPE_TOML).(*TOMLCodec))
}

func TestJSONEncode(t *testing.T) {
	jsonBytes, err := codecFactory.JSONEncode(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte(userJson), jsonBytes)
}

func TestJSONEncodeToString(t *testing.T) {
	jsonStr, err := codecFactory.JSONEncodeToString(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, userJson, jsonStr)
}

func TestJSONEncodeToWriter(t *testing.T) {
	err := codecFactory.JSONEncodeToWriter(user, os.Stdout)
	assert.Equal(t, nil, err)
}

func TestJSONDecode(t *testing.T) {
	c := new(testUser)
	err := codecFactory.JSONDecode([]byte(userJson), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}

func TestJSONDecodeFromString(t *testing.T) {
	c := new(testUser)
	err := codecFactory.JSONDecodeFromString(userJson, c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}

func TestJSONDecodeFromReader(t *testing.T) {
	c := new(testUser)
	err := codecFactory.JSONDecodeFromReader(strings.NewReader(userJson), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}

func TestXMLEncode(t *testing.T) {
	userBytes, err := codecFactory.XMLEncode(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte(userXml), userBytes)
}

func TestXMLEncodeToString(t *testing.T) {
	userStr, err := codecFactory.XMLEncodeToString(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, userXml, userStr)
}

func TestXMLEncodeToWriter(t *testing.T) {
	err := codecFactory.XMLEncodeToWriter(user, os.Stdout)
	assert.Equal(t, nil, err)
}

func TestXMLDecode(t *testing.T) {
	c := new(testUser)
	err := codecFactory.XMLDecode([]byte(userXml), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}
func TestXMLDecodeFromString(t *testing.T) {
	c := new(testUser)
	err := codecFactory.XMLDecodeFromString(userXml, c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}

func TestXMLDecodeFromReader(t *testing.T) {
	c := new(testUser)
	err := codecFactory.XMLDecodeFromReader(strings.NewReader(userXml), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}
func TestYAMLEncode(t *testing.T) {
	yamlBytes, err := codecFactory.YAMLEncode(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte(userYaml), yamlBytes)
}

func TestYAMLEncodeToString(t *testing.T) {
	yamlStr, err := codecFactory.YAMLEncodeToString(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, userYaml, yamlStr)
}

func TestYAMLEncodeToWriter(t *testing.T) {
	err := codecFactory.YAMLEncodeToWriter(user, os.Stdout)
	assert.Equal(t, nil, err)
}

func TestYAMLDecode(t *testing.T) {
	c := new(testUser)
	err := codecFactory.YAMLDecode([]byte(userYaml), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}
func TestYAMLDecodeFromString(t *testing.T) {
	c := new(testUser)
	err := codecFactory.YAMLDecodeFromString(userYaml, c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}

func TestYAMLDecodeFromReader(t *testing.T) {
	c := new(testUser)
	err := codecFactory.YAMLDecodeFromReader(strings.NewReader(userYaml), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}
func TestTOMLEncode(t *testing.T) {
	tomlBytes, err := codecFactory.TOMLEncode(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte(userToml), tomlBytes)
}

func TestTOMLEncodeToString(t *testing.T) {
	tomlStr, err := codecFactory.TOMLEncodeToString(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, userToml, tomlStr)
}

func TestTOMLEncodeToWriter(t *testing.T) {
	err := codecFactory.TOMLEncodeToWriter(user, os.Stdout)
	assert.Equal(t, nil, err)
}

func TestTOMLDecode(t *testing.T) {
	c := new(testUser)
	err := codecFactory.TOMLDecode([]byte(userToml), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}
func TestTOMLDecodeFromString(t *testing.T) {
	c := new(testUser)
	err := codecFactory.TOMLDecodeFromString(userToml, c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}

func TestTOMLDecodeFromReader(t *testing.T) {
	c := new(testUser)
	err := codecFactory.TOMLDecodeFromReader(strings.NewReader(userToml), c)
	assert.Equal(t, nil, err)
	assert.Equal(t, user, c)
}

type testCustomCodec struct{}

func (t *testCustomCodec) Encode(data any) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

func (t *testCustomCodec) EncodeToString(data any) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (t *testCustomCodec) EncodeToWriter(data any, writer io.Writer) error {
	panic("not implemented") // TODO: Implement
}

func (t *testCustomCodec) Decode(data []byte, container any) error {
	panic("not implemented") // TODO: Implement
}

func (t *testCustomCodec) DecodeFromString(data string, container any) error {
	panic("not implemented") // TODO: Implement
}

func (t *testCustomCodec) DecodeFromReader(reader io.Reader, container any) error {
	panic("not implemented") // TODO: Implement
}

func TestRegisterCustomCodec(t *testing.T) {
	codecFactory.RegisterCustomCodec(CodecType("custom"), new(testCustomCodec))
	assert.Equal(t, new(testCustomCodec), codecFactory.Get(CodecType("custom")))
}
