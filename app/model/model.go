// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

// App is the golang structure for table app.
type App struct {
	Id          string `orm:"id"           json:"id"`          //
	Name        string `orm:"name"         json:"name"`        //
	Explain     string `orm:"explain"      json:"explain"`     //
	AccessToken string `orm:"access_token" json:"accessToken"` //
	CreateTime  int    `orm:"create_time"  json:"createTime"`  //
}

// Config is the golang structure for table config.
type Config struct {
	Id          string `orm:"id"           json:"id"`          //
	AppId       string `orm:"app_id"       json:"appId"`       //
	Name        string `orm:"name"         json:"name"`        //
	SendGateway string `orm:"send_gateway" json:"sendGateway"` //
	CreateTime  int    `orm:"create_time"  json:"createTime"`  //
	Type        string `orm:"type"         json:"type"`        //
}

// Info is the golang structure for table info.
type Info struct {
	Id         int    `orm:"id"          json:"id"`         //
	AppId      string `orm:"app_id"      json:"appId"`      //
	ConfigId   string `orm:"config_id"   json:"configId"`   //
	ComeFrom   string `orm:"come_from"   json:"comeFrom"`   //
	Method     string `orm:"method"      json:"method"`     //
	MsgTitle   string `orm:"msg_title"   json:"msgTitle"`   //
	MsgBody    string `orm:"msg_body"    json:"msgBody"`    //
	MsgUrl     string `orm:"msg_url"     json:"msgUrl"`     //
	UserIds    string `orm:"user_ids"    json:"userIds"`    //
	PartyIds   string `orm:"party_ids"   json:"partyIds"`   //
	Totag      string `orm:"totag"       json:"totag"`      //
	State      string `orm:"state"       json:"state"`      //
	CreateTime int    `orm:"create_time" json:"createTime"` //
	MethodCron string `orm:"method_cron" json:"methodCron"` //
	MethodNum  int    `orm:"method_num"  json:"methodNum"`  //
}

// Job is the golang structure for table job.
type Job struct {
	Id             int    `orm:"id"              json:"id"`             //
	Name           string `orm:"name"            json:"name"`           //
	Params         string `orm:"params"          json:"params"`         //
	Group          string `orm:"group"           json:"group"`          //
	InvokeTarget   string `orm:"invoke_target"   json:"invokeTarget"`   //
	CronExpression string `orm:"cron_expression" json:"cronExpression"` //
	MisfirePolicy  int    `orm:"misfire_policy"  json:"misfirePolicy"`  //
	Concurrent     int    `orm:"concurrent"      json:"concurrent"`     //
	Status         int    `orm:"status"          json:"status"`         //
	CreateTime     int    `orm:"create_time"     json:"createTime"`     //
	Remark         string `orm:"remark"          json:"remark"`         //
}

// Template is the golang structure for table template.
type Template struct {
	Id          string `orm:"id"           json:"id"`          //
	ConfigId    string `orm:"config_id"    json:"configId"`    //
	SendGateway string `orm:"send_gateway" json:"sendGateway"` //
	Code        string `orm:"code"         json:"code"`        //
	Title       string `orm:"title"        json:"title"`       //
	Content     string `orm:"content"      json:"content"`     //
	CreateTime  int    `orm:"create_time"  json:"createTime"`  //
}

// UserInfo is the golang structure for table userInfo.
type UserInfo struct {
	Id         int    `orm:"id"          json:"id"`         //
	AppId      string `orm:"app_id"      json:"appId"`      //
	InfoId     int    `orm:"info_id"     json:"infoId"`     //
	Status     string `orm:"status"      json:"status"`     //
	CreateTime int    `orm:"create_time" json:"createTime"` //
	UserId     string `orm:"user_id"     json:"userId"`     //
}
