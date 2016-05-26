package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type Application struct {
	Id              int
	Catalog         int    //  `catalog` int(11) DEFAULT NULL COMMENT '分类  1 web 2 wx 3 ios 4 android 5 pc 6 os 7 api',
	Dev_id          int    // `dev_id` int(11) NOT NULL COMMENT '开发者',
	User_id         int    //  `user_id` int(11) NOT NULL COMMENT '开发者user',
	Code            string // `code` varchar(50) NOT NULL COMMENT '编码',
	Name            string // `name` varchar(50) DEFAULT NULL COMMENT '名称',
	Consumer_secret string // `consumer_secret` varchar(64) DEFAULT NULL COMMENT '开发密钥consumer_secret',
	Api_address     string // `api_address` varchar(256) DEFAULT NULL COMMENT 'API 地址',
	Api_id          string // `api_id` varchar(64) DEFAULT NULL COMMENT 'API ID',
	Api_key         string // `api_key` varchar(128) DEFAULT NULL COMMENT 'API Key',
	Api_version     string // `api_version` varchar(32) DEFAULT NULL COMMENT '版本',
	Group_code      string // `group_code` varchar(32) DEFAULT NULL COMMENT '编号',
	Is_active       int    // `is_active` tinyint(4) NOT NULL COMMENT '是否可用',
	Remark          string // `remark` varchar(100) DEFAULT NULL COMMENT '描述',
	Create_time     int    // `create_time` int(11) DEFAULT NULL COMMENT '创建日期',
	Update_time     int    // `update_time` int(11) DEFAULT NULL COMMENT '修改日期',
}

// 数据表名称
func (a Application) TableName() string {
	return "td_applications"
}

func init() {
	// 注册定义的model
	orm.RegisterModel(new(Application))
}

type ApplicationModel struct {
}

func (am ApplicationModel) FindOne(devId int, code string) (app Application, err error) {

	o := orm.NewOrm()
	o.Using("whale_openplat_db")

	queryStr := "SELECT * FROM td_applications WHERE dev_id=? AND code=? AND is_active=1"
	err = o.Raw(queryStr, devId, code).QueryRow(&app)

	return
}

// 获取趣运动api应用记录的配置信息
func (am ApplicationModel) GetQydApi() (app Application, err error) {
	var devId int = 1
	var code string = "qyd_api"

	app, err = am.FindOne(devId, code)

	if err == nil && app.Id <= 0 { // 没有错误，但是记录没有找到
		err = errors.New("没有找到qyd_api配置信息")
	}

	return
}
