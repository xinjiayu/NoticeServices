package tencentcloud

import (
	"NoticeServices/app/model"
	"NoticeServices/plugins/sms/provider"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var logger *glog.Logger

func init() {
	logger = glog.New()
	logger.SetConfigWithMap(g.Map{
		"path":     "./log/sms",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
	})
}

type Instance struct {
}

func (i *Instance) SendSms(ctx *provider.Context, msg *model.InfoData) error {
	smsConfig := ctx.SmsConfig
	logger.Info("sms发送开始")
	secretKey := gconv.String(smsConfig["secretKey"])
	secretId := gconv.String(smsConfig["secretId"])
	signName := gconv.String(smsConfig["sign_name"]) //短信签名
	tplCode := gconv.String(ctx.SendParam["code"])
	var sendObjectList []model.SendObject
	err := gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		return err
	}

	//发送的信息内容采用|线进行内容分割
	TemplateParam := gstr.Explode("|", msg.MsgBody)

	var phoneNumbers []string
	for _, object := range sendObjectList {
		if object.Name == "sms" {
			phoneNumbers = append(phoneNumbers, object.Value)
		}
	}

	result, err := New(secretId, secretKey, signName).
		Request(tplCode, TemplateParam, phoneNumbers)

	if err != nil {
		logger.Error(err)
	}

	logger.Info(result)

	return nil
}
