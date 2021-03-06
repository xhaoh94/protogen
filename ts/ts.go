package ts

import (
	"fmt"
	"io/ioutil"
	"os"
	"protogen/common"
	"strconv"
)

var (
	UseModule bool
)

//Write 写入
func Write() {
	_, err := os.Stat(common.OutPath)
	if err != nil {
		os.Mkdir(common.OutPath, 0777)
	}

	str := "namespace " + common.NameSpace + "{\n"
	if UseModule {
		str = common.Title + "export " + str
	} else {
		str = common.Title + str
	}
	str += writeCmd() + "\n"
	if !common.CreateJson {
		str += writeConf() + "\n"
	}
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
	err = ioutil.WriteFile(common.OutPath+"/ProtoCode.ts", d, 0666)
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

	rpc := "\texport const rpcs:{ [key: string]: string }={\n"
	f := true
	for k := 0; k < len(common.Rpcs); k++ {
		v := common.Rpcs[k]
		if f {
			f = false
			rpc += "\t\t" + common.GetString(v.Req) + ":" + common.GetString(v.Rsp)
		} else {
			rpc += ",\n\t\t" + common.GetString(v.Req) + ":" + common.GetString(v.Rsp)
		}
	}
	rpc += "\n\t}\n"
	cmd := "\texport const cmds:{ [key: number]: string }={\n"
	cfg := "\texport const cfgs:{ [key: string]: string[][] }={\n"
	f = true
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

	r := cmd
	if len(common.Rpcs) > 0 {
		r += rpc
	}
	r += cfg
	return r

}
