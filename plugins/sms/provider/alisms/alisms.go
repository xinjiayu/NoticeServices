package alisms

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gogf/gf/util/gconv"
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
func (sd *Sender) Request(TemplateCode string, TemplateParam, phoneNumbers []string) (string, error) {

	templates, _ := json.Marshal(TemplateParam)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.SignName = sd.SignName //签名不能为空
	request.TemplateCode = TemplateCode
	request.TemplateParam = string(templates)

	for _, phone := range phoneNumbers {
		request.PhoneNumbers = phone
		//这里需要解析下 判断返回状态码
		response, err := sd.Client.SendSms(request)
		if err != nil {
			return gconv.String(response), err
		}
	}

	return "", nil
}
