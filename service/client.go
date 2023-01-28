package service

import (
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/osins/aliyun-cli/config"
	log "github.com/sirupsen/logrus"
)

/**
 * 使用AK&SK初始化账号Client
 * @return Client
 * @throws Exception
 */
func NewClient() (_result *ecs20140526.Client, _err error) {
	aliyunConfig, _err := config.NewAliyunConfig()
	if _err != nil {
		log.Error(_err)
		os.Exit(1)
		return
	}

	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &aliyunConfig.Access.AccessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &aliyunConfig.Access.AccessKeySecret,
	}

	// 访问的域名
	config.Endpoint = &aliyunConfig.Endpoint
	return ecs20140526.NewClient(config)
}
