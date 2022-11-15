package main

import (
	"NoticeServices/app/define"
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

	pluginPath := g.Config().GetString("system.PluginPath")
	cfgFile := pluginPath + "/mail/config.toml"
	cfg := gcfg.New(cfgFile)

	if cfg.GetString("MailHost") == "" {
		logger.Error("发送失败：邮件发送配置文件有误")
		return
	}
	op := new(Options)
	op.MailHost = cfg.GetString("MailHost")
	op.MailPort = cfg.GetInt("MailPort")
	op.MailUser = cfg.GetString("MailUser")
	op.MailPass = cfg.GetString("MailPass")

	var sendObjectList []define.SendObject
	err := gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
	}
	logger.Info(mailArrTo, "邮件发送完成")
}
