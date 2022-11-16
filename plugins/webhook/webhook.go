package main

import (
	"NoticeServices/app/define"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
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
	g.Log().Debug(context.TODO(), "weebhook发送开始")
	pluginPath := g.Cfg().MustGet(context.TODO(), "system.PluginPath").String()

	cfgFile := pluginPath + "/webhook/config.toml"
	cfg, err := gcfg.NewAdapterFile(cfgFile)
	weConfigs := cfg.MustGet(context.TODO(), "webhook").Array()

	op := new(Options)
	op.Subject = msg.MsgTitle
	body := gjson.New(msg)
	bodyJson, err := body.ToJsonString()
	if err != nil {
		g.Log().Error(context.TODO(), "webhook转换数据出错！")
	}
	op.Body = bodyJson

	for _, opData := range weConfigs {
		if err := gconv.Struct(opData, op); err != nil {
			g.Log().Error(context.TODO(), err)
		}
		go PostData(op)
	}

}

//PostData 通过API修改数据
func PostData(o *Options) {

	c := g.Client()
	c.SetHeader("Secret", o.Secret)
	c.SetHeader("Accept", "application/json")
	c.SetHeader("Content-Type", "application/json")
	if r, e := c.Post(context.TODO(), o.PayloadURL, o.Body); e != nil {
		g.Log().Error(context.TODO(), e)
		return

	} else {
		defer r.Close()
		g.Log().Debug(context.TODO(), o.PayloadURL, r.StatusCode)
		body := []byte(r.ReadAllString())
		g.Log().Debug(context.TODO(), body)
	}
}
func main() {

	cfgFile := "config.toml"
	cfg, err := gcfg.NewAdapterFile(cfgFile)
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}
	weConfigs := cfg.MustGet(context.TODO(), "webhook").Array()

	for _, op := range weConfigs {
		o := new(Options)
		if err := gconv.Struct(op, o); err != nil {
			g.Log().Error(context.TODO(), err)
		}
		g.Log().Debug(context.TODO(), o.PayloadURL)
	}

}
