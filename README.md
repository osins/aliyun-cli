## How to Write Go Code

```
https://go.dev/doc/code
```

### 修改安全组ip策略
第一个参数为api的方法名，后续列表为参数列表，参数列表对应api的参数（注意根据提示的参数顺序填写对应的值字符串）
```
bin/aliyun-cli ModifySecurityGroupRule RegionId, SecurityGroupId, SecurityGroupRuleId, SourceCidrIp, PortRange, IpProtocol
```
