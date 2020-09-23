package common

import (
	"fmt"
	"io/ioutil"
)

type (
	MessageStruct struct {
		Cmd   uint32
		Title string
		Datas [][]string
	}
	EnumStruct struct {
		Title string
		Datas [][]string
	}
)

var (
	NameSpace string
	OutPath   string

	Messages []*MessageStruct
	Enums    []*EnumStruct
)

// var types: { [key: string]: string } = {
// 	"0": "bool", "1": "string", "2": "bytes",
// 	"3": "float", "4": "double", "5": "enum",
// 	"6": "int32", "7": "int64",
// 	"8": "uint32", "9": "uint64",
// 	"10": "sint32", "11": "sint64",
// 	"12": "fixed32", "13": "fixed64",
// 	"14": "sfixed32", "15": "sfixed64"
// }
func GetString(str string) string {
	return "\"" + str + "\""
}
func GetType(str string) string {
	switch str {
	case "bool":
		return "boolean"
	case "string":
		return "string"
	case "bytes":
		return "Uint8Array"
	case "float":
		return "number"
	case "double":
		return "number"
	case "enum":
		return "enum"
	case "int32":
		return "number"
	case "int64":
		return "number|Long"
	case "uint32":
		return "number"
	case "uint64":
		return "number|Long"
	case "sint32":
		return "number"
	case "sint64":
		return "number|Long"
	case "fixed32":
		return "number"
	case "fixed64":
		return "number|Long"
	case "sfixed32":
		return "number"
	case "sfixed64":
		return "number|Long"
	default:
		return str
	}
}
func GetId(str string) string {
	return GetString(cov(str))
}
func cov(str string) string {
	switch str {
	case "bool":
		return "0"
	case "string":
		return "1"
	case "bytes":
		return "2"
	case "float":
		return "3"
	case "double":
		return "4"
	case "enum":
		return "5"
	case "int32":
		return "6"
	case "int64":
		return "7"
	case "uint32":
		return "8"
	case "uint64":
		return "9"
	case "sint32":
		return "10"
	case "sint64":
		return "11"
	case "fixed32":
		return "12"
	case "fixed64":
		return "13"
	case "sfixed32":
		return "14"
	case "sfixed64":
		return "15"
	default:
		return str
	}
}
func FilePathContent(path string, out *[]string) {
	path += "/"
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			FilePathContent(path+file.Name()+"/", out)
		} else {
			f, err := ioutil.ReadFile(path + file.Name())
			if err != nil {
				fmt.Println("read fail", err)
			}
			*out = append(*out, string(f))
		}
	}
}
