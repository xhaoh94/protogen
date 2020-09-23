package main

import (
	"dpb/common"
	"dpb/ts"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	messageReg      *regexp.Regexp
	messageTitleReg *regexp.Regexp
	enumReg         *regexp.Regexp
	cmdReg          *regexp.Regexp
	contextReg      *regexp.Regexp

	msgs []*common.MessageStruct

	fileName string
)

func main() {
	messageReg = regexp.MustCompile(`message ([^}]+)}`)
	messageTitleReg = regexp.MustCompile(`message ([^{]+)`)
	enumReg = regexp.MustCompile(`enum ([^}]+)}`)
	cmdReg = regexp.MustCompile(`cmd:([\d]+)`)

	contextReg = regexp.MustCompile(`{([^}]+)}`)
	out := make([]string, 0)
	common.FilePathContent("file/", &out)
	for _, v := range out {
		parse(v)
	}
	common.NameSpace = "pb"
	ts.WriteCode(msgs)
	// ts.WriteCmd(msgs)
	// ts.WriteConf(msgs)
	writeConfJson()
}

func parse(str string) {
	strs := messageReg.FindAllString(str, -1)
	for _, context := range strs {
		s := &common.MessageStruct{}
		s.Datas = make([][]string, 0)
		cmdMatched := cmdReg.FindStringSubmatch(context)
		if len(cmdMatched) == 2 {
			num, err := strconv.Atoi(cmdMatched[1])
			if err != nil {
				fmt.Printf("cmd非int类型")
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
		msgs = append(msgs, s)
	}
}

func writeConfJson() {
	fileName = "out/ProtoCfg.json"
	str := "{\n"
	cmd := "\t" + common.GetString("cmds") + ":{\n"
	cfg := "\t" + common.GetString("cfgs") + ":{\n"
	f := true
	for j := 0; j < len(msgs); j++ {
		v := msgs[j]
		if v.Cmd > 0 {
			if f {
				f = false
				cmd += "\t\t" + common.GetString(strconv.Itoa(int(v.Cmd))) + ":" + common.GetString(v.Title)
			} else {
				cmd += ",\n\t\t" + common.GetString(strconv.Itoa(int(v.Cmd))) + ":" + common.GetString(v.Title)
			}
		}
		cfg += "\t\t" + common.GetString(v.Title) + ":[\n"
		for i := 0; i < len(v.Datas); i++ {
			c := v.Datas[i]
			cfg += "\t\t\t" + "["
			cfg += common.GetString(c[0]) + "," + common.GetString(c[2]) + "," + common.GetId(c[1])
			isArray := c[len(c)-1] == "1"
			if isArray {
				cfg += "," + common.GetString("1")
			}
			if i == len(v.Datas)-1 {
				cfg += "]"
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
	cfg += "\t}\n"
	cmd += "\n\t},\n"
	str += cmd
	str += cfg
	str += "}"

	var d = []byte(str)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}
