package models

import (
	"encoding/json"
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/utils"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Message struct {
	Id           int
	Message      string // `message` longtext COMMENT '消息 ',
	Status       int    // `status` int(1) DEFAULT NULL COMMENT '处理状态  0 未处理 1 处理中 2 已经处理 '
	App_id       string // `app_id` varchar(32) DEFAULT NULL COMMENT '场馆app_id ',
	Client_time  int    // `client_time` int(11) DEFAULT NULL,
	Rs_key       string // `rs_key` varchar(32) DEFAULT NULL COMMENT ' 缓存标示 ',
	Client_token string // `client_token` varchar(64) DEFAULT NULL COMMENT '客户端唯一标示',
}

// 数据表名称
func (m Message) TableName() string {
	return "di_message4"
}

func init() {
	// 注册定义的model
	orm.RegisterModel(new(Message))
}

const (
		Message_RS_Key = "Message.rs_key"
		Message_Sign = "gate.guanzhang.me.Msg"
		processing = 1
		processed = 2
	)

// 消息模型
type MessageModel struct {
}

//消息在处理
func (mm *MessageModel) CheckMsgProcessing(api_sign string) bool {
	if utils.Redis.IsExist(getMsgSign(api_sign)) {
		cacheData := utils.Redis.Get(getMsgSign(api_sign))
		var ret int
		json.Unmarshal(cacheData.([]byte), &ret)
		return ret == processing
	}
	return false
}

//消息已经处理完成
func (mm *MessageModel) CheckMsgProcessed(api_sign string) bool {
	if utils.Redis.IsExist(getMsgSign(api_sign)) {
		cacheData := utils.Redis.Get(getMsgSign(api_sign))
		var ret int
		json.Unmarshal(cacheData.([]byte), &ret)
		return ret == processed
	}
	return false
}

//设置Msg正在处理
func (mm *MessageModel)SetMsgProcessing(api_sign string)  {
	utils.Redis.Put(getMsgSign(api_sign), processing, 120*time.Second) //预估处理时间为2分钟
}

//设置Msg处理完成
func (mm *MessageModel)SetMsgProcessed(api_sign string)  {
	utils.Redis.Put(getMsgSign(api_sign), processed, 3600*time.Second) //缓存1小时
}

func getMsgSign(api_sign string) string {
	return  Message_Sign + ":" + api_sign
}

//删除消息处理标志
func (mm *MessageModel) DelMsgProcess(api_sign string) error {
	return utils.Redis.Delete(getMsgSign(api_sign))
}

//消息已经处理完成
func (mm *MessageModel) CheckUploaded(api_sign string) bool {
	if utils.Redis.IsExist(api_sign) {
		cacheData := utils.Redis.Get(api_sign)
		var ret int
		err := json.Unmarshal(cacheData.([]byte), &ret)
		if err == nil {
			return ret > 0
		}

	}
	return false
}

//设置消息已经处理完成
func (mm *MessageModel) setUploaded(msg Message) error {
	//
	return utils.SetCacheOneHour(Message_RS_Key+":"+msg.Rs_key, msg.Id)
}

// 是否已经收到消息
func (mm *MessageModel) HasReceivedMessage(Rs_key string) (id int) {
	cacheData := utils.Redis.Get(Message_RS_Key + ":" + Rs_key)
	if cacheData != nil {
		json.Unmarshal(cacheData.([]byte), &id)
		return
	}
	return -1
}

// 插入消息
func (mm *MessageModel) InsertMessage(msg Message) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&msg)
	if err == nil {
		msg.Id = (int)(id)
		utils.Log.Info("Message id : %s ", msg.Id)
		//设置消息已经上传
		err = mm.setUploaded(msg)
		if err != nil {
			utils.Log.Warn("cache设置错误 ： %v", err)
		}
	}
	return
}

// 生成rs_key
func CreateRsKey(msg Message) (rsKey string) {
	start := strings.Index(msg.Message, "Value")
	end := len(msg.Message) - 1

	rsKey = helper.Md5(msg.App_id + msg.Message[start:end])

	return rsKey
}
