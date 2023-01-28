package service

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/osins/aliyun-cli/config"
	log "github.com/sirupsen/logrus"
)

func NewSecurityGroups(aliyunConfig *config.AliyunConfig) Cli {
	sg := &SecurityGroups{
		conf: *aliyunConfig,
	}

	return sg
}

type SecurityGroups struct {
	conf config.AliyunConfig
}

/**
 * 使用AK&SK初始化账号Client
 * @return Client
 * @throws Exception
 */
func (s *SecurityGroups) createClient(accessKeyId *string, accessKeySecret *string) (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}

	// 访问的域名
	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	return ecs20140526.NewClient(config)
}

func (s *SecurityGroups) Run(args []*string) (_err error) {
	if len(args) < 5 {
		log.Error("参数不正确, 依次为: RegionId, SecurityGroupId, SecurityGroupRuleId, SourceCidrIp, PortRange, IpProtocol")
		return &tea.SDKError{
			Message: tea.String("参数不正确, 依次为: RegionId, SecurityGroupId, SecurityGroupRuleId, SourceCidrIp, PortRange, IpProtocol"),
		}
	}

	client, _err := s.createClient(&s.conf.Access.AccessKeyId, &s.conf.Access.AccessKeySecret)
	if _err != nil {
		log.Errorf("创建客户端失败，%s", _err.Error())

		return _err
	}

	modifySecurityGroupRuleRequest := &ecs20140526.ModifySecurityGroupRuleRequest{
		RegionId:            args[1],
		SecurityGroupId:     args[2],
		SecurityGroupRuleId: args[3],
		SourceCidrIp:        args[4],
		IpProtocol:          args[5],
		PortRange:           args[6],
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.ModifySecurityGroupRuleWithOptions(modifySecurityGroupRuleRequest, runtime)
		if _err != nil {
			log.Errorf("api运行失败, %s", _err.Error())

			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			log.Errorf("api运行失败%s", tryErr.Error())

			error.Message = tea.String(tryErr.Error())
		}

		result, _err := util.AssertAsString(error.Message)
		if _err != nil {
			log.Errorf("api运行失败, %s", _err.Error())

			return _err
		}

		log.Errorf("api运行失败, %s", *result)
	}
	return _err
}
