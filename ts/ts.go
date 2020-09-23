package ts

import (
	"dpb/common"
	"fmt"
	"io/ioutil"
	"strconv"
)

func WriteCode(msgs []*common.MessageStruct) {

	str := "export module " + common.NameSpace + "{\n"
	str += writeCmd(msgs) + "\n"
	str += writeConf(msgs) + "\n"
	for _, v := range msgs {
		str += "\texport interface " + v.Title + " {\n"
		for _, c := range v.Datas { //tag type name isArray
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

func writeCmd(msgs []*common.MessageStruct) string {

	// str := "export module " + common.NameSpace + "{\n"
	str := "\texport const enum Cmds" + " {\n"
	for _, v := range msgs {
		if v.Cmd > 0 {
			str += "\t\t" + v.Title + " = " + strconv.Itoa(int(v.Cmd)) + ",\n"
		}
	}
	// str += "\t}\n"
	str += "\t}"

	return str
}

func writeConf(msgs []*common.MessageStruct) string {

	// str := "export module " + common.NameSpace + "{\n"
	// str := "\texport class ProtoCfg {\n"
	cmd := "\texport var cmds:{ [key: number]: string }={\n"
	cfg := "\texport var cfgs:{ [key: string]: string[][] }={\n"
	f := true
	for j := 0; j < len(msgs); j++ {
		v := msgs[j]
		if v.Cmd > 0 {
			if f {
				f = false
				cmd += "\t\t" + strconv.Itoa(int(v.Cmd)) + ":" + common.GetString(v.Title)
			} else {
				cmd += ",\n\t\t" + strconv.Itoa(int(v.Cmd)) + ":" + common.GetString(v.Title)
			}
		}
		cfg += "\t\t" + common.GetString(v.Title) + ":["
		for i := 0; i < len(v.Datas); i++ {
			c := v.Datas[i]
			cfg += "[" + common.GetString(c[0]) + "," + common.GetString(c[2]) + "," + common.GetId(c[1])
			isArray := c[len(c)-1] == "1"
			if isArray {
				cfg += "," + common.GetString("1")
			}
			if i == len(v.Datas)-1 {
				cfg += "]"
			} else {
				cfg += "],"
			}
		}
		if j == len(msgs)-1 {
			cfg += "]\n"
		} else {
			cfg += "],\n"
		}
	}
	cfg += "\t}\n"
	cmd += "\n\t}\n"
	// str += cmd
	// str += cfg
	// str += "\t}\n"

	// str += "}"
	return cmd + cfg
	// var d = []byte(str)
	// err := ioutil.WriteFile("out/Protocfg.ts", d, 0666)
	// if err != nil {
	// 	fmt.Println("write fail")
	// }
}
