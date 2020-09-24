package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"protogen/common"
	"protogen/golang"
	"protogen/ts"
	"regexp"
	"strconv"
	"strings"
)

var (
	messageReg      *regexp.Regexp
	messageTitleReg *regexp.Regexp
	enumReg         *regexp.Regexp
	enumTitleReg    *regexp.Regexp
	cmdReg          *regexp.Regexp
	rpcReg          *regexp.Regexp
	contextReg      *regexp.Regexp

	fileName string
)

func init() {
	messageReg = regexp.MustCompile(`message ([^}]+)}`)
	messageTitleReg = regexp.MustCompile(`message ([^{]+)`)

	enumReg = regexp.MustCompile(`enum ([^}]+)}`)
	enumTitleReg = regexp.MustCompile(`enum ([^{]+)`)

	cmdReg = regexp.MustCompile(`cmd=([\d]+)`)
	rpcReg = regexp.MustCompile(`rpc<([^>]+)>`)
	contextReg = regexp.MustCompile(`{([^}]+)}`)

}

func main() {
	parse()
}
func parse() {
	codeType := flag.String("code_type", "", "生产代码类型")
	inPath := flag.String("in_path", "file/", "proto文件目录")
	outPath := flag.String("out_path", "out/", "导出文件目录")
	ns := flag.String("namespace", "pb", "生成代码命名空间")
	createJSON := flag.Bool("create_json", true, "是否生成json配置")
	useModule := flag.Bool("use_module", false, "(typescript)是否使用模块模式")
	flag.Parse()
	if *codeType == "" {
		fmt.Println("生成失败！没有指定生成代码类型")
		return
	}
	common.CreateJson = *createJSON
	common.OutPath = *outPath
	common.NameSpace = *ns
	ts.UseModule = *useModule
	out := make([]string, 0)
	common.FilePathContent(*inPath, &out)
	for _, v := range out {
		b := parseMessage(v)
		if !b {
			return
		}
		parseEnum(v)
		parseRPC(v)
	}
	switch *codeType {
	case "ts":
		ts.Write()
		break
	case "csharp":
		break
	case "golang":
		golang.Write()
		break
	default:
		fmt.Println("代码类型未实现")
		return
	}
	if *createJSON {
		writeJSON()
	}
}

func parseRPC(str string) {
	strs := rpcReg.FindAllString(str, -1)
	for _, context := range strs {
		if strings.Index(context, ":") < 0 {
			continue
		}
		s := &common.RpcStruct{}
		rpcMatched := rpcReg.FindStringSubmatch(context)
		if len(rpcMatched) == 2 {
			str := strings.Replace(rpcMatched[1], " ", "", -1)
			str = strings.Replace(str, "<", "", -1)
			str = strings.Replace(str, ">", "", -1)
			ss := strings.Split(str, ":")
			s.Req = ss[0]
			s.Rsp = ss[1]
		}
		common.Rpcs = append(common.Rpcs, s)
	}
}
func parseEnum(str string) {
	strs := enumReg.FindAllString(str, -1)
	for _, context := range strs {
		s := &common.EnumStruct{}
		s.Datas = make([][]string, 0)
		titleMatched := enumTitleReg.FindStringSubmatch(context)
		if len(titleMatched) == 2 {
			s.Title = titleMatched[1]
		}

		contMatched := contextReg.FindStringSubmatch(context)
		if len(contMatched) == 2 {
			lines := strings.Split(contMatched[1], "\n")
			for i := 0; i < len(lines); i++ {
				line := lines[i]
				if strings.Index(line, "=") < 0 {
					continue
				}
				c := strings.Split(line, "=")
				if strings.Index(c[0], "//") >= 0 {
					continue
				}
				datas := make([]string, 0)

				c[1] = strings.Replace(c[1], " ", "", -1)
				endIndex := strings.Index(c[1], ";")
				tag := c[1][0:endIndex]
				datas = append(datas, tag)
				ts := strings.Split(c[0], " ")
				for _, v := range ts {
					if v != "" {
						datas = append(datas, v)
					}
				}
				s.Datas = append(s.Datas, datas)
			}
		}

		common.Enums = append(common.Enums, s)
	}
}

func parseMessage(str string) bool {
	strs := messageReg.FindAllString(str, -1)
	for _, context := range strs {
		s := &common.MessageStruct{}
		s.Datas = make([][]string, 0)
		cmdMatched := cmdReg.FindStringSubmatch(context)
		if len(cmdMatched) == 2 {
			num, err := strconv.Atoi(cmdMatched[1])
			if err != nil {
				fmt.Printf("cmd非int类型%s", cmdMatched[1])
				return false
			}
			s.Cmd = uint32(num)
		}
		titleMatched := messageTitleReg.FindStringSubmatch(context)
		if len(titleMatched) == 2 {
			s.Title = titleMatched[1]
		}
		contMatched := contextReg.FindStringSubmatch(context)
		if len(contMatched) == 2 {
			var startindex int
			if s.Cmd > 0 {
				startindex = 1
			}
			lines := strings.Split(contMatched[1], "\n")
			for i := startindex; i < len(lines); i++ {
				line := lines[i]
				if strings.Index(line, "=") < 0 {
					continue
				}
				c := strings.Split(line, "=")
				if strings.Index(c[0], "//") >= 0 {
					continue
				}

				datas := make([]string, 0)
				c[1] = strings.Replace(c[1], " ", "", -1)
				endIndex := strings.Index(c[1], ";")
				tag := c[1][0:endIndex]
				datas = append(datas, tag)

				ts := strings.Split(c[0], " ")
				for _, v := range ts {
					if v != "" && v != "repeated" {
						datas = append(datas, v)
					}
				}
				if strings.Index(c[0], "repeated") >= 0 {
					datas = append(datas, "1")
				} else {
					datas = append(datas, "0")
				}
				s.Datas = append(s.Datas, datas)
			}
		}
		common.Messages = append(common.Messages, s)
	}
	return true
}

func writeJSON() {
	fmt.Println("write json start")
	fileName = common.OutPath + "/ProtoCfg.json"
	str := "{\n"

	rpc := "\t" + common.GetString("rpcs") + ":{\n"
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
	rpc += "\n\t},\n"

	cmd := "\t" + common.GetString("cmds") + ":{\n"
	cfg := "\t" + common.GetString("cfgs") + ":{\n"
	f = true
	for j := 0; j < len(common.Messages); j++ {
		v := common.Messages[j]
		if v.Cmd > 0 {
			if f {
				f = false
				cmd += "\t\t" + common.GetString(strconv.Itoa(int(v.Cmd))) + ":" + common.GetString(v.Title)
			} else {
				cmd += ",\n\t\t" + common.GetString(strconv.Itoa(int(v.Cmd))) + ":" + common.GetString(v.Title)
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
	cmd += "\n\t},\n"
	str += cmd
	str += rpc
	str += cfg
	str += "}"

	var d = []byte(str)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write json fail")
	}
	fmt.Println("write json success")
}
