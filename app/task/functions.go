package task

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/define"
	"NoticeServices/app/notifieer"
	"context"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func init() {
	var task1 Entity
	task1.FuncName = "TestMes"
	task1.Param = nil
	task1.Run = JobSendMessage
	Add(task1)

}

//JobSendMessage 通知发送任务
func JobSendMessage() {

	task := GetByName("JobSendMessage")
	if task == nil {
		return
	}
	for _, v := range task.Param {
		info, _ := dao.Info.Ctx(context.TODO()).Where("id", v).One()
		infoData := new(define.InfoData)
		if err := gconv.Struct(info, infoData); err != nil {
			glog.Error(err)
		}
		notifieer.Instance.GateWaySend(infoData)
	}

}
