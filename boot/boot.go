package boot

import (
	_ "NoticeServices/packed"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

const Instant string = "instant"         //即时发送
const Appointment string = "appointment" //预约发送
const Regular string = "regular"         //定期发送

func init() {
	err := gtime.SetTimeZone("Asia/Shanghai") //设置系统时区
	if err != nil {
		glog.Error(err)
	}
	logPath := g.Config().GetString("logger.Path")
	err = glog.SetPath(logPath)
	if err != nil {
		glog.Error(err)
	}
}
