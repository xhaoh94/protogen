package golang

import (
	"fmt"
	"io/ioutil"
	"os"
	"protogen/common"
	"strconv"
)

func Write() {
	_, err := os.Stat(common.OutPath)
	if err != nil {
		os.Mkdir(common.OutPath, 0777)
	}

	str := "package " + common.NameSpace + "\n\nconst (\n"
	for _, v := range common.Messages {
		if v.Cmd > 0 {
			str += "\t" + v.Title + " uint32 = " + strconv.Itoa(int(v.Cmd)) + "\n"
		}
	}
	str += ")"

	var d = []byte(str)
	err = ioutil.WriteFile(common.OutPath+"/cmd.pb.go", d, 0666)
	if err != nil {
		fmt.Println("write golang fail")
	} else {
		fmt.Println("write golang success")
	}
}
