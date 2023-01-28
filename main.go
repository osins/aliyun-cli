// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"os"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/osins/aliyun-cli/config"
	"github.com/osins/aliyun-cli/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := "请输入正确的参数, 依次为: api方法名，参数列表"
	if len(os.Args[1:]) == 0 {
		log.Error(err)
		os.Exit(1)
		return
	}

	aliyunConfig, _err := config.NewAliyunConfig()
	if _err != nil {
		log.Error(_err)
		os.Exit(1)
		return
	}

	switch api := os.Args[1:][0]; api {
	case "ModifySecurityGroupRule":
		sg := service.NewSecurityGroups(aliyunConfig)
		err := sg.Run(tea.StringSlice(os.Args[1:]))
		if err != nil {
			log.Errorf(err.Error())
			os.Exit(1)
			return
		}
	default:
		log.Error(err)
		os.Exit(1)
		return
	}

	os.Exit(0)
}
