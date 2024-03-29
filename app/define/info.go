package define

import "NoticeServices/app/model"

type SendParam struct {
	aa map[string]string
}

type InfoData struct {
	AppId      string `orm:"app_id"      json:"app_id"`      //
	ConfigId   string `orm:"config_id"   json:"config_id"`   //
	ComeFrom   string `orm:"come_from"   json:"come_from"`   //
	Method     string `orm:"method"      json:"method"`      //
	MethodCron string `orm:"method_cron" json:"method_cron"` //
	MethodNum  int    `orm:"method_num"  json:"method_num"`  //
	MsgTitle   string `orm:"msg_title"   json:"msg_title"`   //
	MsgBody    string `orm:"msg_body"    json:"msg_body"`    //
	MsgUrl     string `orm:"msg_url"     json:"msg_url"`     //
	UserIds    string `orm:"user_ids"    json:"user_ids"`    //
	PartyIds   string `orm:"party_ids"   json:"party_ids"`   //
	Totag      string `orm:"totag"       json:"totag"`       //
}

// 组合模型，通知信息
type EntityInfo struct {
	UserInfo *model.UserInfo `json:"user_info"`
	Info     *MsgInfo        `json:"info"`
}

type MsgInfo struct {
	MsgTitle string `orm:"msg_title"   json:"msg_title"` //
	MsgBody  string `orm:"msg_body"    json:"msg_body"`  //
	MsgUrl   string `orm:"msg_url"     json:"msg_url"`   //
}
