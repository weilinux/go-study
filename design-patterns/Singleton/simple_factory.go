package main

import "fmt"

// ConfigParser 配置解析接口
type ConfigParser interface {
	Parse(p []byte)
}

// JsonParser jsonParser JSON 文件解析器
type JsonParser struct {
}

func (j *JsonParser) Parse(p []byte) {
	fmt.Println(p)
}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}

// YamlParser Yaml 文件解析器
type YamlParser struct {
}

func (y *YamlParser) Parse(p []byte) {
	fmt.Println(p)
}

func NewYamlParser() *YamlParser {
	return &YamlParser{}
}

type ConfigType uint8

const (
	JsonType ConfigType = 1 << iota
	YamlType
	HtmlType
	Html4Type
)

func NewConfig(t ConfigType) ConfigParser {
	switch t {
	case JsonType:
		return NewJsonParser()
	case YamlType:
		return NewYamlParser()
	default:
		return nil
	}
}

func main() {
	// 调用方法码
	jsonParser := NewConfig(JsonType)
	yamlParser := NewConfig(YamlType)
	fmt.Printf("%p\n", jsonParser)
	fmt.Printf("%p\n", yamlParser)
	fmt.Println(JsonType, YamlType, HtmlType, Html4Type)
}
