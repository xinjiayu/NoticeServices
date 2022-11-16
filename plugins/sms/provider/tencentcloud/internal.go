package tencentcloud

import (
	"NoticeServices/app/define"
	"NoticeServices/plugins/sms/provider"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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

func (i *Instance) SendSms(ctx *provider.Context, msg *define.InfoData) error {
	smsConfig := ctx.SmsConfig
	g.Log().Debug(context.TODO(), "sms发送开始")
	secretKey := gconv.String(smsConfig["secretKey"])
	secretId := gconv.String(smsConfig["secretId"])
	signName := gconv.String(smsConfig["sign_name"]) //短信签名
	tplCode := gconv.String(ctx.SendParam["code"])
	var sendObjectList []define.SendObject
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
		g.Log().Error(context.TODO(), err)
	}

	g.Log().Debug(context.TODO(), result)

	return nil
}
