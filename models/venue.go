package models

import (
	"gate.guanzhang.me/api"
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/utils"
	"github.com/astaxie/beego/orm"
)

type Venue struct {
	Id             int     // `id` int(11) NOT NULL AUTO_INCREMENT,
	Name           string  // `name` varchar(128) DEFAULT NULL COMMENT '场馆名称',
	Qyd_court_code string  // `qyd_court_code` varchar(50) DEFAULT NULL COMMENT '趣运动场地编码',
	Supplier_id    int     // `supplier_id` int(11) DEFAULT NULL COMMENT '商户',
	City           string  // `city` varchar(20) DEFAULT NULL COMMENT '城市',
	Region         string  // `region` varchar(20) DEFAULT NULL COMMENT '区域',
	Sub_region     string  // `sub_region` varchar(20) DEFAULT NULL COMMENT '商圈',
	Tel            string  // `tel` varchar(15) DEFAULT NULL COMMENT '电话',
	Latitude       float64 // `latitude` float(10,5) DEFAULT NULL COMMENT '维度',
	Longitude      float64 // `longitude` float(10,5) DEFAULT NULL COMMENT '经度',
	Address        string  // `address` varchar(200) DEFAULT NULL COMMENT '地址',
	Hot            int     // `hot` smallint(4) DEFAULT NULL COMMENT '热度排序',
	Transit_info   string  // `transit_info` varchar(256) DEFAULT NULL COMMENT '公交信息',
	Metro_info     string  // `metro_info` varchar(256) DEFAULT NULL COMMENT '地铁信息',
	Is_delete      int     // `is_delete` tinyint(1) DEFAULT NULL COMMENT '是否删除',
	Display_sort   int     // `display_sort` int(11) DEFAULT NULL COMMENT '显示顺序',
	Status         int     // `status` int(2) DEFAULT NULL COMMENT '1 启用 2 禁用',
	Register_date  int     // `register_date` int(11) DEFAULT NULL COMMENT '注册时间',
	Active_date    int     // `active_date` int(11) DEFAULT NULL COMMENT '上线时间',
}

// 数据表名称
func (v Venue) TableName() string {
	return "st_venues"
}

func init() {
	// 注册定义的model
	orm.RegisterModel(new(Venue))
}

type VenueModel struct {
}

// 请求趣运动api获取获取场馆基本信息
func (this *VenueModel) GetCourtInfoList(appId string) (courtInfo api.CourtInfo, retErr error) {
	// 获取接口配置信息
	applicationModel := ApplicationModel{}
	application, err := applicationModel.GetQydApi()
	if err != nil { // 获取api配置信息出错
		utils.Log.Error("获取接口配置信息失败, err: %v", err) // 记录log
		retErr = err
		return
	}

	// 组织请求获取appkey的数据
	params := make(map[string]string)
	params["api_key"] = application.Api_id
	params["res_key"] = application.Api_key
	params["app_id"] = appId
	params["ver"] = application.Api_version

	// 调用趣运动接口获取获取场馆基本信息
	venueApi := api.Venue{}
	return venueApi.GetCourtInfoList(application.Api_address, params)
}

// 获取场馆信息
func (this *VenueModel) getVenueBySource(qydVenueId int) (venue Venue, err error) {
	o := orm.NewOrm()
	o.Using("whale_business_db")

	query_str := "select * from st_venues where qyd_court_code = ? and status=1 limit 1"
	err = o.Raw(query_str, qydVenueId).QueryRow(&venue)

	return
}

// 保存场馆信息
func (this *VenueModel) SaveVenue(courtInfo api.CourtInfo, supplierId int) (int, error) {
	o := orm.NewOrm()
	o.Using("whale_business_db")

	var venue Venue
	venue.Address = courtInfo.Address                // 地址
	venue.City = courtInfo.City_name                 // 城市
	venue.Tel = courtInfo.Telephone                  // 电话
	venue.Name = courtInfo.Name + "-肖颜春"             // 场馆名称
	venue.Supplier_id = supplierId                   // 商户id
	venue.Qyd_court_code = courtInfo.Venues_id       // 趣运动场场馆id
	venue.Active_date = int(helper.GetTimestamp())   // 上线时间
	venue.Register_date = int(helper.GetTimestamp()) // 注册时间
	venue.Status = 1                                 // 启用

	utils.Log.Info("SaveVenue:场馆信息,venue: %+v", venue)

	// 插入数据
	insertId, err := o.Insert(&venue)

	var id int

	if err == nil && insertId > 0 { // 将int64转换为int类型返回
		id = int(insertId)
		utils.Log.Info("SaveVenue:insert_success,insertId: %d, id: %d", insertId, id)
	} else {
		utils.Log.Error("SaveVenue:insert_failed,err: %v", err)
	}

	return id, err
}
