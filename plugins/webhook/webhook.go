package main

import (
	"NoticeServices/app/model"
	"github.com/gogf/gf/os/glog"
)

func Send(msg *model.InfoData) {
	glog.Info("通过webhook进行发送...")
	glog.Info(msg)
}
