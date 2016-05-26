package models

import (
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/validate"
	"github.com/astaxie/beego/orm"
)

type Operate struct {
	ID             int    // `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	App_id         string // `app_id` varchar(32) DEFAULT NULL COMMENT 'AppId',
	Client_id      string // `client_id` varchar(100) DEFAULT NULL,
	Venue_id       int    // `venue_id` int(11) DEFAULT NULL COMMENT '场馆',
	Code           string // `code` varchar(32) DEFAULT NULL COMMENT '编码',
	Name           string // `name` varchar(32) DEFAULT NULL COMMENT '操作名称',
	Create_user    string // `create_user` varchar(32) DEFAULT NULL COMMENT '操作用户',
	Create_time    int    // `create_time` int(11) DEFAULT NULL,
	Client_version string // `client_version` varchar(32) DEFAULT NULL COMMENT 'Version 1.0',
	Status         int    // `status` int(1) DEFAULT NULL COMMENT '处理状态  ',
	Upload_time    int    // `upload_time` int(11) DEFAULT NULL,
}

const (
		Operater_Sign = "gate.guanzhang.me.Operate"
	)

// 数据表名称
func (o Operate) TableName() string {
	return "di_operate4"
}

func init() {
	orm.RegisterModel(new(Operate))
}

type OperateModel struct {
}

// 添加操作记录
func (om *OperateModel) InsertOperate(params validate.EtlOperateParams, metadata validate.OperateMetadata) (insertId int64, retErr error) {

	sm := SmosModel{}
	smos, err := sm.FindSmos(params.App_id)

	if err != nil {
		retErr = err
		return
	}

	var operate Operate
	operate.App_id = params.App_id                   // app_id
	operate.Venue_id = smos.Venue_id                 // 场馆id
	operate.Upload_time = int(helper.GetTimestamp()) // 上传时间
	operate.Status = 1                               // 状态

	operate.Client_id = metadata.Metadata.ClientID
	operate.Client_version = metadata.Metadata.Version

	operate.Code = metadata.Value.Code
	operate.Name = metadata.Value.Name
	operate.Create_user = metadata.Value.CreateBy

	operate.Create_time = int(helper.StrToTimestamp(metadata.Value.CreateTime))

	o := orm.NewOrm()
	insertId, retErr = o.Insert(&operate)

	return
}
