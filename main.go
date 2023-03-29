package main

import (
	"fmt"
	"net"
	"os"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}

// 获取本地网卡的公网IPv6地址，不包含临时公网IPv6地址。
func GetEthIPv6Addr() string {
	eth_interface, err := net.InterfaceByName("以太网")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	eth_addr, err := eth_interface.Addrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	eth_addr_ipv6_cidr := eth_addr[0].String()
	return eth_addr_ipv6_cidr[:len(eth_addr_ipv6_cidr)-3]
}

func _main(args []*string) (_err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String("AccessKey ID"), tea.String("AccessKey Secret"))
	if _err != nil {
		return _err
	}

	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		//RecordId通过DescribeDomainRecords API查询获得
		RecordId: tea.String("182054496302945"), //已脱敏
		RR:       tea.String("nuc.home"),
		Type:     tea.String("AAAA"),
		Value:    tea.String(GetEthIPv6Addr()),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
