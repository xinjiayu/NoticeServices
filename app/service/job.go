package service

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/model"
	"NoticeServices/app/task"
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"strings"
)

func AutoAllTask() {
	//自动执行已开启的任务
	glog.Info("===========自动执行已开启的任务===========")
	jobs, err := GetJobs()
	if err != nil {
		glog.Error(err)
	}
	for _, job := range jobs {
		if err = JobStart(job); err != nil {
			//glog.Error(err)
		}
	}
}

//JobReqAdd 添加操作请求参数
type JobReqAdd struct {
	Name           string `p:"name" v:"required#任务名称不能为空"`
	Params         string `p:"params"` // 任务参数
	Group          string `p:"group" `
	InvokeTarget   string `p:"invoke_target" v:"required#执行方法不能为空"`
	CronExpression string `p:"cron_expression" v:"required#任务表达式不能为空"`
	MisfirePolicy  int    `p:"misfire_policy"`
	Concurrent     int    `p:"concurrent" `
	Status         int    `p:"status" v:"required#状态（0正常 1暂停）不能为空"`
	Remark         string `p:"remark" `
}

func GetJobs() (jobs []*model.Job, err error) {
	err = dao.Job.Ctx(context.TODO()).Where(dao.Job.Columns.Status, "0").Scan(&jobs)
	return
}

//添加计划任务
func JobAdd(jobData *model.Job) (id int64, err error) {
	glog.Info("======添加任务=======", jobData.Name)
	res, err := dao.Job.Ctx(context.TODO()).FieldsEx(dao.Job.Columns.Id).Insert(jobData)
	if err != nil {
		glog.Error(err)
		err = gerror.New("添加任务失败")
	}
	id, err = res.LastInsertId()
	if err != nil {
		glog.Error(err)
		err = gerror.New("添加任务失败")
	}
	return
}

//启动任务
func JobStart(job *model.Job) error {
	//可以task目录下是否绑定对应的方法
	glog.Info("======启动任务======", job.InvokeTarget)
	f := task.GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}
	//传参
	paramArr := strings.Split(job.Params, "|")
	task.EditParams(f.FuncName, paramArr)
	rs := gcron.Search(job.InvokeTarget)
	if rs == nil {
		if job.MisfirePolicy == 1 {
			taskJob, err := gcron.AddSingleton(job.CronExpression, f.Run, job.InvokeTarget)
			if err != nil || taskJob == nil {
				return err
			}
		} else {
			taskJob, err := gcron.AddOnce(job.CronExpression, f.Run, job.InvokeTarget)
			if err != nil || taskJob == nil {
				return err
			}
		}
	}
	gcron.Start(job.InvokeTarget)
	if job.MisfirePolicy == 1 {
		job.Status = 0
		if _, err := dao.Job.Ctx(context.TODO()).Data(job).Update(); err != nil {
			//glog.Error(err)
		}
	}
	return nil
}

//JobStop 停止任务
func JobStop(job *model.Job) error {
	//可以task目录下是否绑定对应的方法
	f := task.GetByName(job.InvokeTarget)
	if f == nil {
		return gerror.New("当前task目录下没有绑定这个方法")
	}
	rs := gcron.Search(job.InvokeTarget)
	if rs != nil {
		gcron.Remove(job.InvokeTarget)
	}
	job.Status = 1
	if _, err := dao.Job.Ctx(context.TODO()).Update(job); err != nil {
		//glog.Error(err)
	}
	return nil
}
