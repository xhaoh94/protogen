package ts

import (
	"dpb/common"
	"fmt"
	"io/ioutil"
	"strconv"
)

var (
	UseModule bool
)

//Write 写入
func Write() {
	fmt.Printf("write ts start")
	str := "namespace " + common.NameSpace + "{\n"
	if UseModule {
		str = "export " + str
	}
	str += writeCmd() + "\n"
	str += writeConf() + "\n"
	for _, v := range common.Enums {
		str += "\texport const enum " + v.Title + " {\n"
		f := true
		for _, c := range v.Datas {
			if f {
				f = false
				str += "\t\t" + c[1] + "=" + c[0]
			} else {
				str += ",\n\t\t" + c[1] + "=" + c[0]
			}
		}
		str += "\n\t}\n"
	}
	for _, v := range common.Messages {
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
	err := ioutil.WriteFile(common.OutPath+"/ProtoCode.ts", d, 0666)
	if err != nil {
		fmt.Println("write ts fail")
	} else {
		fmt.Println("write ts success")
	}
}

func writeCmd() string {

	str := "\texport const enum Cmds" + " {\n"
	for _, v := range common.Messages {
		if v.Cmd > 0 {
			str += "\t\t" + v.Title + " = " + strconv.Itoa(int(v.Cmd)) + ",\n"
		}
	}
	str += "\t}"

	return str
}

func writeConf() string {

	cmd := "\texport var cmds:{ [key: number]: string }={\n"
	cfg := "\texport var cfgs:{ [key: string]: string[][] }={\n"
	f := true
	for j := 0; j < len(common.Messages); j++ {
		v := common.Messages[j]
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
		if j == len(common.Messages)-1 {
			cfg += "]\n"
		} else {
			cfg += "],\n"
		}
	}
	cfg += "\t}\n"
	cmd += "\n\t}\n"
	return cmd + cfg

}
