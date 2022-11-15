package main

import (
	"NoticeServices/app/service"
	_ "NoticeServices/boot"
	"NoticeServices/library/version"
	_ "NoticeServices/router"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	service.AutoAllTask()
	g.Server().Run()
}
