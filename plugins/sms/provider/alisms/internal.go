package alisms

import (
	"NoticeServices/app/model"
	"NoticeServices/plugins/sms/provider"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
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
	ak := gconv.String(smsConfig["access_key_id"])
	sk := gconv.String(smsConfig["access_key_secret"])
	signName := gconv.String(smsConfig["sign_name"]) //短信签名
	tplCode := gconv.String(ctx.SendParam["code"])

	var sendObjectList []model.SendObject
	err := gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		return err
	}
	for _, object := range sendObjectList {
		if object.Name == "sms" {
			//op.MailTo = object.Value
			//op.Subject = msg.MsgTitle
			//op.Body = msg.MsgBody
			//sendMail(op)
			logger.Info("发短信到：", object.Value)
			result, err := New(ak, sk).
				Request(object.Value, signName, tplCode, msg.MsgBody)

			if err != nil {
				logger.Error(err)
			}

			logger.Info(result)
		}
	}

	return nil
}
