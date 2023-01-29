// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"os"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/osins/aliyun-cli/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	errMessage := "请输入正确的参数, 依次为: api方法名，参数列表"
	if len(os.Args[1:]) == 0 {
		log.Error(errMessage)
		os.Exit(1)
		return
	}

	log.Debugf("args: %s", strings.Join(os.Args[1:], ", "))

	commandString := os.Args[1:][0]
	params := tea.StringSlice(os.Args[1:])

	commands := service.NewCommands()
	if cli, ok := commands[commandString]; ok {
		err := cli.Run(params)
		if err != nil {
			log.Errorf(err.Error())
			os.Exit(1)
			return
		}
	} else {
		log.Errorf("您要执行的命令: %s, 不存在.", commandString)
		os.Exit(1)
		return
	}

	log.Debugf("run command[ %s ] complete", commandString)

	os.Exit(0)
}
