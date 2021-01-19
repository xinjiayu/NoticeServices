package task

import (
	"github.com/gogf/gf/os/glog"
)

func init() {
	var task1 Entity
	task1.FuncName = "TestMes"
	task1.Param = nil
	task1.Run = TestMes
	Add(task1)

}

func TestMes() {
	glog.Info("====w定时任务测试方法=============")
}

func JobSendMessage(infoId string) {
	//
	//info, err := service.GetOneInfo(infoId)
	//if err != nil {
	//	glog.Error(err)
	//}
	//infoData := new(model.InfoData)
	//gconv.Struct(info, infoData)
	//service.Msg.Send(infoData)
}
