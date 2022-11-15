package boot

import (
	_ "NoticeServices/packed"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

const Instant string = "instant"         //即时发送
const Appointment string = "appointment" //预约发送
const Regular string = "regular"         //定期发送

func init() {
	err := gtime.SetTimeZone("Asia/Shanghai") //设置系统时区
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}
	logPath := g.Cfg().MustGet(context.TODO(), "logger.Path").String()

	err = glog.SetPath(logPath)
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}
}
