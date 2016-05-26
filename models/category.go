package models

import (
	"gate.guanzhang.me/api"
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/utils"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Category struct {
	Id               int     //	`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	Supplier_id      int     // `supplier_id` int(11) DEFAULT '0',
	Venue_id         int     // `venue_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '场馆ID',
	Cat_id           int     // `cat_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '分类ID  1. 篮球',
	Characteristic   string  // `characteristic` varchar(100) NOT NULL DEFAULT '' COMMENT '特色',
	Is_delete        int     // `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
	Is_order         int     // `is_order` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否可在线订购',
	Is_hot           int     // `is_hot` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否热门 1预订页面推荐',
	Hot_image        string  // `hot_image` varchar(128) DEFAULT '' COMMENT '预订页面推荐图片',
	Is_index         int     // `is_index` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否首页推荐 1推荐',
	Index_image      string  // `index_image` varchar(128) DEFAULT '' COMMENT '首页推荐图片',
	Price            float64 // `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '平均价格',
	Price_desc       string  // `price_desc` varchar(180) DEFAULT '' COMMENT '场馆价格说明',
	Start_hour       int     // `start_hour` smallint(3) NOT NULL DEFAULT '9' COMMENT '开门时间',
	End_hour         int     /// `end_hour` smallint(3) NOT NULL DEFAULT '23' COMMENT '关门时间',
	Sort             int     // `sort` smallint(5) unsigned NOT NULL DEFAULT '99' COMMENT '排序',
	Index_sort       int     // `index_sort` smallint(5) unsigned NOT NULL DEFAULT '99' COMMENT '热门推荐排序',
	Order_type       int     // `order_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '订单类型',
	Book_cycle       int     // `book_cycle` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '预订周期 0一周1三天',
	Book_interval    int     // `book_interval` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '预订时间 2二小时后,1一小时后',
	Course_count     int     // `course_count` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '场地数量',
	Is_refund        int     // `is_refund` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否可退款（0不可退款、1可退款）',
	Close_date       string  // `close_date` varchar(128) DEFAULT '' COMMENT '暂停营业日期',
	Image_url        string  // `image_url` varchar(200) DEFAULT '' COMMENT '场馆项目图片',
	Is_course        int     // `is_course` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否月卡',
	Course_notice    string  // `course_notice` varchar(128) DEFAULT '' COMMENT '通知',
	Create_by        int     // `create_by` int(10) NOT NULL DEFAULT '0' COMMENT '操作人员',
	Create_time      int     // `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
	Last_modify_by   int     // `last_modify_by` int(10) DEFAULT '0' COMMENT '操作人员',
	Last_modify_time int     // `last_modify_time` int(11) unsigned DEFAULT '0' COMMENT '最后更新时间',
}

// 数据表名称
func (c Category) TableName() string {
	return "st_categorys"
}

func init() {
	// 注册定义的model
	orm.RegisterModel(new(Category))
}

type CategoryModel struct {
}

// 保存分类
func (this *CategoryModel) SaveCategory(appId string, supplierId int, venueId int, courtInfo api.CourtInfo) (successNum int) {
	o := orm.NewOrm()
	o.Using("whale_business_db")

	utils.Log.Info("SaveCategory:start, appId: %s, supplierId: %d, venueId: %d, 共 %d 条分类",
		appId, supplierId, venueId, len(courtInfo.Categories))

	if len(courtInfo.Categories) == 0 { // 没有分类
		return
	}

	var category Category

	// 循环插入分类
	for i, cat := range courtInfo.Categories {
		catId, err := strconv.Atoi(cat.Cat_id)
		if err != nil { // 转换失败
			utils.Log.Error("SaveCategory:insert_%v, catId: %d, err: %v", i, catId, err)
		}

		isCardOrder, err := strconv.Atoi(cat.Is_card_order)
		if err != nil { // 转换失败
			utils.Log.Error("SaveCategory:insert_%v, isCardOrder: %d, err: %v", i, isCardOrder, err)
		}

		category.Supplier_id = supplierId                 // 商家id
		category.Venue_id = venueId                       // 场馆id
		category.Cat_id = catId                           // 分类id
		category.Create_time = int(helper.GetTimestamp()) // 添加时间
		category.Is_delete = 0                            // 是否删除
		category.Is_order = isCardOrder                   // 是否可在线订购
		category.Characteristic = cat.Cat_name + "-肖颜春"   // 分类名称

		utils.Log.Info("SaveCategory:insert_%v, category: %+v", i, category)

		insertId, err := o.Insert(&category)
		if err == nil && insertId > 0 { // 插入成功
			successNum++
			utils.Log.Info("SaveCategory:insert_%v_success, insertId: %v", i, insertId)
		} else { // 插入失败
			utils.Log.Error("SaveCategory:insert_%v_failed, err: %v", i, err)
		}
	}

	utils.Log.Info("SaveCategory:end, 成功添加 %d 条分类", successNum)

	return
}
