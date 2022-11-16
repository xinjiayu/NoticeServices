package main

import (
	"NoticeServices/app/define"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"gopkg.in/gomail.v2"
	"strings"
)

var logger *glog.Logger

type Options struct {
	MailHost string
	MailPort int
	MailUser string // 发件人
	MailPass string // 发件人密码
	MailTo   string // 收件人 多个用,分割
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

func init() {
	logger = glog.New()
	logger.SetConfigWithMap(g.Map{
		"path":     "./log/mail",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
	})
}

//Send 发送
func Send(sendParam map[string]interface{}, msg *define.InfoData) {

	pluginPath := g.Cfg().MustGet(context.TODO(), "system.PluginPath").String()

	cfgFile := pluginPath + "/webhook/config.toml"
	cfg, err := gcfg.NewAdapterFile(cfgFile)
	mailHostCfg := cfg.MustGet(context.TODO(), "MailHost").String()

	if mailHostCfg == "" {
		g.Log().Error(context.TODO(), "发送失败：邮件发送配置文件有误")
		return
	}
	op := new(Options)
	op.MailHost = cfg.MustGet(context.TODO(), "MailHost").String()
	op.MailPort = cfg.MustGet(context.TODO(), "MailPort").Int()
	op.MailUser = cfg.MustGet(context.TODO(), "MailUser").String()
	op.MailPass = cfg.MustGet(context.TODO(), "MailPass").String()

	var sendObjectList []define.SendObject
	err = gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}

	for _, object := range sendObjectList {
		if object.Name == "mail" {
			op.MailTo = object.Value
			op.Subject = msg.MsgTitle
			op.Body = msg.MsgBody
			sendMail(op)

		}
	}

}

func sendMail(o *Options) {

	m := gomail.NewMessage()
	//设置发件人
	m.SetHeader("From", o.MailUser)
	//设置发送给多个用户
	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)
	//设置邮件主题
	m.SetHeader("Subject", o.Subject)

	//设置邮件正文
	m.SetBody("text/html", o.Body)
	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)

	err := d.DialAndSend(m)
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}
	g.Log().Debug(context.TODO(), mailArrTo, "邮件发送完成")
}
