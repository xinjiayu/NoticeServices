package main

import (
	"NoticeServices/app/model"
	"github.com/gogf/gf/os/glog"
)

func Send(msg *model.InfoData) {
	glog.Info("通过邮件进行发送...")
	glog.Info(msg)
}
