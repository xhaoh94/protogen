package common

import (
	"fmt"
	"io/ioutil"
)

type (
	MessageStruct struct {
		Cmd    uint32
		Title  string
		Cacels [][]string
	}
)

var (
	NameSpace string
)

func GetString(str string) string {
	return "\"" + str + "\""
}
func GetType(str string) string {
	switch str {
	case "string":
		return "string"
	case "uint32":
		return "number"
	case "int32":
		return "number"
	case "float":
		return "number"
	default:
		return str
	}
}

func FilePathContent(path string, out *[]string) {
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
