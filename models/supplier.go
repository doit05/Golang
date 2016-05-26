package models

import (
	"gate.guanzhang.me/api"
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/utils"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Supplier struct {
	Id                 int    // `id` int(11) NOT NULL AUTO_INCREMENT,
	Code               string // `code` varchar(50) DEFAULT NULL COMMENT '编码',
	Source_type        int    // `source_type` int(2) DEFAULT NULL COMMENT '来源类型： 1 趣运动 2 申请 3 馆掌 4 后台系统 ',
	Source_id          int    // `source_id` int(11) DEFAULT NULL COMMENT '来源地id',
	Full_name          string // `full_name` varchar(50) DEFAULT NULL COMMENT '全称',
	Simple_name        string // `simple_name` varchar(50) DEFAULT NULL COMMENT '简称',
	City               string // `city` varchar(50) DEFAULT NULL COMMENT '城市',
	Contact_user       string // `contact_user` varchar(50) DEFAULT NULL COMMENT '联系人',
	Org_phone          string // `org_phone` varchar(50) DEFAULT NULL COMMENT '单位电话',
	Mobile             string // `mobile` varchar(50) DEFAULT NULL COMMENT '手机号码',
	Email              string // `email` varchar(50) DEFAULT NULL COMMENT '电子邮件',
	Biz_lic_code       string // `biz_lic_code` varchar(50) DEFAULT NULL COMMENT '营业执照编号',
	Tax_code           string // `tax_code` varchar(50) DEFAULT NULL COMMENT '税务登记证号',
	Pinyin_code        string // `pinyin_code` varchar(50) DEFAULT NULL COMMENT '拼音码',
	Org_address        string // `org_address` varchar(50) DEFAULT NULL COMMENT '单位地址',
	Is_active          int    // `is_active` tinyint(4) DEFAULT NULL COMMENT '是否激活',
	Status             int    // `status` int(2) DEFAULT NULL COMMENT '签约状态',
	Start_date         int    // `start_date` int(11) DEFAULT NULL COMMENT '上线时间',
	Remark             string // `remark` varchar(50) DEFAULT NULL,
	Create_by          int    // `create_by` int(11) NOT NULL DEFAULT '0',
	Create_time        int    // `create_time` int(11) unsigned NOT NULL DEFAULT '0',
	Last_modify_by     int    // `last_modify_by` int(11) NOT NULL DEFAULT '0',
	Last_modify_time   int    // `last_modify_time` int(11) unsigned NOT NULL DEFAULT '0',
	Is_test            int    // `is_test` int(11) DEFAULT '0',
	Is_enable_wxmember int    // `is_enable_wxmember` tinyint(1) DEFAULT '0' COMMENT '是否启用微信会员',
}

// 数据表名称
func (s Supplier) TableName() string {
	return "st_suppliers"
}

func init() {
	// 注册定义的model
	orm.RegisterModel(new(Supplier))
}

type SupplierModel struct {
}

// 保存商家信息
func (this *SupplierModel) SaveSupplier(courtInfo api.CourtInfo) (int, error) {
	o := orm.NewOrm()
	o.Using("whale_business_db")

	qydVenueId, err := strconv.Atoi(courtInfo.Venues_id)

	utils.Log.Info("SaveSupplier:类型转换:qydVenueId: %d, err: %v", qydVenueId, err)

	var supplier Supplier
	supplier.Org_address = courtInfo.Address         // 单位地址
	supplier.City = courtInfo.City_name              // 城市
	supplier.Mobile = courtInfo.Telephone            // 手机号码
	supplier.Simple_name = courtInfo.Name + "-肖颜春"   // 简称
	supplier.Start_date = int(helper.GetTimestamp()) // 上线时间
	supplier.Is_active = 1                           // 激活
	supplier.Source_type = 1                         // 类型
	supplier.Source_id = qydVenueId                  // 来源地id
	supplier.Status = 1                              // 签约
	supplier.Create_time = int(helper.GetTimestamp())

	utils.Log.Info("SaveSupplier:商家信息:supplier: %+v", supplier)

	// 插入数据
	insertId, err := o.Insert(&supplier)

	var id int

	if err == nil && insertId > 0 { // 将int64转换为int类型返回
		id = int(insertId)
		utils.Log.Info("SaveSupplier:insert_success,insertId: %d, id: %d", insertId, id)
	} else {
		utils.Log.Error("SaveSupplier:insert_failed,err: %v", err)
	}

	return id, err
}
