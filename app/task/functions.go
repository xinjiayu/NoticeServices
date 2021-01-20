package task

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/model"
	"NoticeServices/app/notifieer"
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
		info, _ := dao.Info.Where("id", v).One()
		infoData := new(model.InfoData)
		gconv.Struct(info, infoData)
		notifieer.Instance.GateWaySend(infoData)
	}

}
