package models

import (
	"encoding/json"
	// "fmt"
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

// '客户端数据信息';
type Style struct {
	Id             int
	Client_token   string //client_token` varchar(64) DEFAULT NULL COMMENT '客户端唯一标示',
	App_id         string //`app_id` varchar(32) DEFAULT NULL COMMENT 'AppId',
	Source_name    string //`source_name` varchar(32) DEFAULT NULL COMMENT 'SourceName',
	Action_name    string //`action_name` varchar(32) DEFAULT NULL COMMENT 'ActionName',
	Data_format    string //`data_format` varchar(32) DEFAULT NULL COMMENT 'json',
	Create_time    int64  //`create_time` int(11) DEFAULT NULL,
	Client_version string //`client_version` varchar(32) DEFAULT NULL COMMENT 'Version 1.0',
	Md5_code       string //`md5_code` varchar(400) DEFAULT NULL COMMENT 'MD5HashoCode ',
	Value          string //`value` varchar(400) DEFAULT NULL,
	Status         int64  //`status` int(1) DEFAULT NULL COMMENT '处理状态  ',
}

// 数据表名称
func (this *Style) TableName() string {
	return "di_style_4"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(Style))
}

type StyleModel struct { //业务模型
}

var sources = []string{"1"} //数据来源，1表示来自管掌

//查找客户端信息, 有缓存则优先读缓存
func (this *StyleModel) FindStyles(app_id, client_token, client_version string) (data []Style, err error) {
	data = make([]Style, len(sources))
	for i := 0; i < len(sources); i++ {
		var style Style

		//从缓存中查找
		cacheKey := app_id + "_" + client_token + "_" + sources[i]
		cacheData := utils.Redis.Get(cacheKey) //从cache中查找客户端信息

		if cacheData != nil {	// 有缓存
			err := json.Unmarshal(cacheData.([]byte), &style)
			if err != nil {
				utils.Log.Error("解析对象错误 warnMsg : %v ", err)
				break
			}
		} else {
			utils.Log.Warn("cache中未找到的客户端信息 wranMsg : %s", cacheKey)

			//cache中未找到，从数据库中查找
			style = Style{App_id: app_id, Client_token: client_token, Source_name: sources[i]}
			cacheCondition := false  //控制不允许插入空的style
			style_list, err1 := this.FindOneStyle(style)

			 //查找成功，返回一个数组，取数组中的第一个
			if err1 == nil && len(style_list) > 0 {
				style = style_list[0]
				cacheCondition = true
			} else {
				//如果数据库中没有找到,则创建新的记录
				s, err2 := this.AddStyle(app_id, client_token, sources[i], client_version, "0")
				if err2 == nil {
					style = s
					cacheCondition = true
					utils.Log.Info("将Style对象插入数据库 Id ＝ %s ", style.Id)
				} else {//插入数据库出错，跳到下次循环
					utils.Log.Error("将Style对象插入数据库错误 errMsg ： %v ", err)
					break
				}
			}
			// 设置缓存
			if cacheCondition {
				err = this.setRedis(cacheKey, style)
				if err != nil {
					utils.Log.Warn("在cache中设置style 失败errMsg : %v ", err)
				}
			}
		}
		data[i] = style //向返回结果中添加数据
	}
	return
}

func (this *StyleModel) setRedis(cacheKey string, style Style) (err error) {
	jsonData, jErr := json.Marshal(style)

	if jErr == nil {
		pErr := utils.Redis.Put(cacheKey, jsonData, 3600*time.Second) // 缓存1小时
		if pErr != nil {
			err = pErr
			utils.Log.Warn("添加redis缓存失败, jsonData: %s", jsonData)
		}
	} else {
		err = jErr
		utils.Log.Warn("json编码失败, smos: %+v, err: %v", style, jErr)
	}
	return
}

func (this *StyleModel) FindOneStyle(s Style) (style_list []Style, err error) {
	o := orm.NewOrm()
	select_str := "select * from di_style where app_id = '" + s.App_id + "' and client_token = '" + s.Client_token + "'  and source_name = '" + s.Source_name + "' and status = 1"
	num, err := o.Raw(select_str).QueryRows(&style_list)
	if num == 0 && err != nil { //如果，数据库中未找到，则返回nil和err
		return nil, err
	}
	return //如查询成功返回style_list数组
}

func (this *StyleModel) AddStyle(app_id, client_token, source_name,
	client_version, action_name string) (s Style, err error) {

	s = Style{Client_token: client_token, App_id: app_id, Source_name: source_name,
		Action_name: action_name, Client_version: client_version}
	s.Create_time = helper.GetTimestamp()
	s.Status = 1
	s.Value = "5"

	o := orm.NewOrm()
	var Id int64
	Id, err = o.Insert(&s)
	s.Id = int(Id)

	return s, err
}
