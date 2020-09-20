package ts

import (
	"dpb/common"
	"fmt"
	"io/ioutil"
	"strconv"
)

func WriteCode(msgs []*common.MessageStruct) {

	str := "export module " + common.NameSpace + "{\n"
	for _, v := range msgs {
		str += "\texport interface " + v.Title + " {\n"
		for _, c := range v.Cacels { //tag type name isArray
			isArray := c[len(c)-1] == "1"
			if isArray {
				str += "\t\t" + c[2] + ":" + common.GetType(c[1]) + "[];\n"
			} else {
				str += "\t\t" + c[2] + ":" + common.GetType(c[1]) + ";\n"
			}
		}
		str += "\t}\n"
	}
	str += "}"

	var d = []byte(str)
	err := ioutil.WriteFile("out/Proto.ts", d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
}

func WriteCmd(msgs []*common.MessageStruct) {

	// str := "export module " + common.NameSpace + "{\n"
	str := "export enum ProtoCmd" + " {\n"
	for _, v := range msgs {
		if v.Cmd > 0 {
			str += "\t" + v.Title + " = " + strconv.Itoa(int(v.Cmd)) + ",\n"
		}
	}
	// str += "\t}\n"
	str += "}"

	var d = []byte(str)
	err := ioutil.WriteFile("out/ProtoCmd.ts", d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
}

func WriteConf(msgs []*common.MessageStruct) {

	str := "export module " + common.NameSpace + "{\n"
	str += "\texport class ProtoCfg {\n"
	cmd := "\t\t public static cmds:{ [key: number]: string }={\n"
	cfg := "\t\t public static cfgs:{ [key: string]: string[][] }={\n"
	f := true
	for j := 0; j < len(msgs); j++ {
		v := msgs[j]
		if v.Cmd > 0 {
			if f {
				f = false
				cmd += "\t\t\t" + strconv.Itoa(int(v.Cmd)) + ":" + common.GetString(v.Title)
			} else {
				cmd += ",\n\t\t\t" + strconv.Itoa(int(v.Cmd)) + ":" + common.GetString(v.Title)
			}
		}
		cfg += "\t\t\t" + common.GetString(v.Title) + ":[\n"
		for i := 0; i < len(v.Cacels); i++ {
			c := v.Cacels[i]
			cfg += "\t\t\t\t" + "["
			cfg += common.GetString(c[0]) + "," + common.GetString(c[1]) + "," + common.GetString(c[2])
			isArray := c[len(c)-1] == "1"
			if isArray {
				cfg += "," + common.GetString("1")
			}
			if i == len(v.Cacels)-1 {
				cfg += "]\n\t\t\t"
			} else {
				cfg += "],\n"
			}
		}
		if j == len(msgs)-1 {
			cfg += "]\n"
		} else {
			cfg += "],\n"
		}
	}
	cfg += "\t\t}\n"
	cmd += "\n\t\t}\n"
	str += cmd
	str += cfg
	str += "\t}\n"
	str += "}"

	var d = []byte(str)
	err := ioutil.WriteFile("out/Protocfg.ts", d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
}
