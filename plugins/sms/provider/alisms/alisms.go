package alisms

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

type Sender struct {
	// 请求参数
	Params   map[string]string
	Client   *dysmsapi.Client
	SignName string
}

// 初始化一些参数
func New(regionId, keyId, secret, signName string) *Sender {

	client, err := dysmsapi.NewClientWithAccessKey(regionId, keyId, secret)
	if err != nil {
		panic(err)
	}
	sd := &Sender{}
	sd.SignName = signName

	sd.Client = client

	return sd
}

// 发送请求
func (sd *Sender) Request(TemplateCode, TemplateParam string, phoneNumbers []string) (string, error) {

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = strings.Replace(strings.Trim(fmt.Sprint(phoneNumbers), "[]"), " ", ",", -1)
	request.SignName = sd.SignName
	request.TemplateCode = TemplateCode
	request.TemplateParam = TemplateParam

	glog.Info(request)
	response, err := sd.Client.SendSms(request)
	if err != nil {
		return "", err
	}
	return gconv.String(response), nil
}
