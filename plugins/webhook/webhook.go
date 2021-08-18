package main

import (
	"NoticeServices/app/define"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type Options struct {
	PayloadURL string
	Secret     string
	Subject    string
	Body       string
}

type WebhookConfig struct {
	PayloadURL string
	Secret     string
}

var logger *glog.Logger

func init() {
	logger = glog.New()
	logger.SetConfigWithMap(g.Map{
		"path":     "./log/webhook",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
	})
}

//goland:noinspection GoUnusedExportedFunction
func Send(sendParam map[string]interface{}, msg *define.InfoData) {
	logger.Info("weebhook发送开始")
	pluginPath := g.Config().GetString("system.PluginPath")
	cfgFile := pluginPath + "/webhook/config.toml"
	cfg := gcfg.New(cfgFile)

	weConfigs := cfg.GetArray("webhook")

	op := new(Options)
	op.Subject = msg.MsgTitle
	body := gjson.New(msg)
	bodyJson, err := body.ToJsonString()
	if err != nil {
		logger.Error("webhook转换数据出错！")
	}
	op.Body = bodyJson

	for _, opData := range weConfigs {
		if err := gconv.Struct(opData, op); err != nil {
			glog.Error(err)
		}
		go PostData(op)
	}

}

//PostData 通过API修改数据
func PostData(o *Options) {

	c := ghttp.NewClient()
	c.SetHeader("Secret", o.Secret)
	c.SetHeader("Accept", "application/json")
	c.SetHeader("Content-Type", "application/json")
	if r, e := c.Post(o.PayloadURL, o.Body); e != nil {
		logger.Error(e)
		return

	} else {
		defer r.Close()
		logger.Info(o.PayloadURL, r.StatusCode)
		body := []byte(r.ReadAllString())
		logger.Info(body)
	}
}
func main() {
	cfgFile := "config.toml"
	cfg := gcfg.New(cfgFile)
	weConfigs := cfg.GetArray("webhook")
	for _, op := range weConfigs {
		o := new(Options)
		if err := gconv.Struct(op, o); err != nil {
			glog.Error(err)
		}
		glog.Info(o.PayloadURL)
	}

}
