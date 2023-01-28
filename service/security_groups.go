package service

import (
	"errors"
	"fmt"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	log "github.com/sirupsen/logrus"
)

type SecurityGroups struct {
}

func (s *SecurityGroups) Run(args []*string) (_err error) {
	if len(args) < 5 {
		log.Error("参数不正确, 依次为: RegionId, SecurityGroupId, SecurityGroupRuleId, SourceCidrIp, PortRange, IpProtocol")
		return errors.New("参数不正确, 依次为: RegionId, SecurityGroupId, SecurityGroupRuleId, SourceCidrIp, PortRange, IpProtocol")
	}

	client, _err := NewClient()
	if _err != nil {
		log.Errorf("创建客户端失败，%s", _err.Error())

		return errors.New("参数不正确, 依次为: RegionId, SecurityGroupId, SecurityGroupRuleId, SourceCidrIp, PortRange, IpProtocol")
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

			return errors.New("参数不正确, 依次为: RegionId, SecurityGroupId, SecurityGroupRuleId, SourceCidrIp, PortRange, IpProtocol")
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

			return fmt.Errorf("api运行失败, %s", _err.Error())
		}

		log.Errorf("api运行失败, %s", *result)

		return fmt.Errorf("api运行失败, %s", *result)
	}

	return nil
}
