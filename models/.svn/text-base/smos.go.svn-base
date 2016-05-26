package models

import (
	"encoding/json"
	"errors"
	//"fmt"
	"gate.guanzhang.me/api"
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/utils"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Smos struct {
	Id           int    // `id` int(11) NOT NULL AUTO_INCREMENT,
	App_id       string // `app_id` varchar(32) DEFAULT NULL,
	App_key      string // `app_key` varchar(64) DEFAULT NULL,
	Status       int    // `status` tinyint(1) DEFAULT NULL,
	Start_date   int    // `start_date` int(11) DEFAULT NULL COMMENT '生成日期',
	Venue_id     int    // `venue_id` int(10) DEFAULT NULL COMMENT '场地ID',
	Supplier_id  int    // `supplier_id` int(10) DEFAULT NULL COMMENT '商户ID',
	Telephone    string // `telephone` varchar(15) DEFAULT NULL COMMENT '电话',
	Name         string // `name` varchar(128) DEFAULT NULL COMMENT '场馆名称',
	Create_by    int    // `create_by` int(11) DEFAULT NULL,
	Create_time  int    // `create_time` int(11) DEFAULT NULL,
	Qyd_venue_id string // `qyd_venue_id` varchar(50) DEFAULT NULL COMMENT '趣运动场馆Id',
}

// 数据表名称
func (s Smos) TableName() string {
	return "bs_smos"
}

func init() {
	// 注册定义的model
	orm.RegisterModel(new(Smos))
}

const OBJECT_APP_ID_KEY = "business_Smos_AppId"

type SmosModel struct {
}

// 通过app_id获取app_key
func (this *SmosModel) GetAppkey(app_id string) (string, error) {
	smos, err := this.GetSmosFromCache(app_id)

	if err != nil {
		return "", err
	}

	if smos.Id > 0 {
		return smos.App_key, nil
	} else {
		return "", errors.New("没有查找到app_key")
	}
}

// 获取smos信息，有缓存会优先读取缓存
func (this *SmosModel) GetSmosFromCache(app_id string) (smos Smos, retErr error) {
	cacheKey := OBJECT_APP_ID_KEY + ":" + app_id // 缓存键名
	cacheData := utils.Redis.Get(cacheKey)       // 优先从缓存中获取

	if cacheData != nil { // 有缓存
		err := json.Unmarshal(cacheData.([]byte), &smos)
		if err == nil {
			// 解析成功, 直接返回
			return
		} else {
			// 解析失败，不影响主业务，只记录log
			utils.Log.Warn("解析json数据出错, app_id: %s, cacheData:%v, err: %v", app_id, cacheData, err)
		}
	}

	// 从数据库读取
	smos, err := this.GetSmos(app_id)
	if err != nil { // 获取出错
		utils.Log.Error("获取出错smos数据出错, app_id: %s, err: %v", app_id, err)
		retErr = err
		return
	}

	if smos.Id <= 0 { // 记录没有找到
		utils.Log.Error("smos记录没有找到, app_id: %s,", app_id)
		retErr = errors.New("记录不存在")
		return
	}

	// 成功读取数据，进行缓存
	/*jsonData, jErr := json.Marshal(smos)

	if jErr == nil {
		pErr := utils.Redis.Put(cacheKey, jsonData, 3600*time.Second) // 缓存1小时
		if pErr != nil {
			utils.Log.Warn("添加redis缓存失败, jsonData: %s", jsonData)
		}
	} else {
		utils.Log.Warn("json编码失败, smos: %+v, err: %v", smos, jErr)
	}*/

	return
}

// 获取smos信息， 如果本库没有，则会调用趣运动接口，并初始化本库的smos信息
func (this *SmosModel) GetSmos(appId string) (smos Smos, retErr error) {

	smos, err := this.FindSmos(appId)
	if err != nil { // 查询出错， 直接返回
		retErr = err
		return
	}

	if smos.Id > 0 { // 本地数据存在
		if len(smos.Qyd_venue_id) == 0 { // 趣运动场馆id不存在
			this.UpdateQydVenueId(smos, appId)
		}
		return
	}

	// 本地数据库查询不到smos信息, 初始化馆掌信息
	this.InitializeSmos(appId)
	return
}

// 初始化馆掌信息
func (this *SmosModel) InitializeSmos(appId string) {

	utils.Log.Info("InitializeSmos:start, appId: %s", appId)

	// 获取app_key
	appKey, err := this.GetAppKeyById(appId)
	if err != nil {
		utils.Log.Error("InitializeSmos:获取app_key失败, err: %v", err)
		return
	}

	utils.Log.Info("InitializeSmos:appKey: %s", appKey)

	venueModel := VenueModel{}
	// 请求趣运动api获取获取场馆基本信息
	courtInfo, err := venueModel.GetCourtInfoList(appId)
	if err != nil {
		utils.Log.Error("InitializeSmos:请求趣运动api获取获取场馆基本信息失败, err: %v", err)
		return
	}

	utils.Log.Info("InitializeSmos:courtInfo: %+v", courtInfo)

	// 获取趣运动venue_id
	var qydVenueId int
	qydVenueId, err = strconv.Atoi(courtInfo.Venues_id)

	utils.Log.Info("InitializeSmos:qydVenueId: %d, err: %v", qydVenueId, err)

	// 获取本地库的场馆信息
	venue, err := venueModel.getVenueBySource(qydVenueId)
	if err != nil {
		utils.Log.Error("InitializeSmos:获取本地库的场馆信息失败, err: %v", err)
		return
	}

	utils.Log.Info("InitializeSmos:venueInfo: %+v", venue)

	var supplierId int = 0
	var venueId int = 0

	if venue.Id > 0 { // 本地库场馆信息存在
		supplierId = venue.Supplier_id
		venueId = venue.Id
	} else {
		//保存商户信息
		supplierModel := SupplierModel{}
		supplierId, err = supplierModel.SaveSupplier(courtInfo)

		if err == nil {
			utils.Log.Info("InitializeSmos:保存商家信息成功,supplierId: %d", supplierId)
		} else {
			utils.Log.Error("InitializeSmos:保存商家信息失败,err: %v", err)
		}

		if supplierId > 0 {
			// 保存场馆信息
			venueId, err = venueModel.SaveVenue(courtInfo, supplierId)
			if err == nil {
				utils.Log.Info("InitializeSmos:保存场馆信息成功,venueId: %d", venueId)
			} else {
				utils.Log.Error("InitializeSmos:保存场馆信息失败,err:%v", err)
			}
		}
	}

	//保存smos
	lastInsertId, err := this.SaveSmos(appId, appKey, supplierId, venueId, qydVenueId, courtInfo.Telephone)
	if err == nil {
		utils.Log.Info("InitializeSmos:保存smos信息成功,last_insert_id: %d", lastInsertId)
	} else {
		utils.Log.Error("InitializeSmos:保存smos信息失败,err:%v", err)
	}

	//保存场地项目
	categoryModel := CategoryModel{}
	categoryModel.SaveCategory(appId, supplierId, venueId, courtInfo)
}

//保存smos信息
func (this *SmosModel) SaveSmos(appId string, appKey string, supplierId int, venueId int, qydVenueId int, telephone string) (int, error) {
	o := orm.NewOrm()
	o.Using("whale_business_db")

	var smos Smos
	smos.App_id = appId
	smos.App_key = appKey
	smos.Supplier_id = supplierId
	smos.Venue_id = venueId
	smos.Status = 1 //
	smos.Create_time = int(helper.GetTimestamp())
	smos.Qyd_venue_id = strconv.Itoa(qydVenueId) // 趣运动场馆id
	smos.Telephone = telephone

	utils.Log.Info("SaveSmos:start,smos: %+v", smos)

	// 插入数据
	insertId, err := o.Insert(&smos)

	var id int

	if err == nil && insertId > 0 { // 将int64转换为int类型返回
		id = int(insertId)
		utils.Log.Info("SaveSmos:inser_success,insertId: %d, id: %d", insertId, id)
	} else {
		utils.Log.Error("SaveSmos:inser_failed,err: %v", err)
	}

	return id, err
}

// 更新趣运动场馆id
func (this *SmosModel) UpdateQydVenueId(smos Smos, appId string) (affectedNum int64) {
	utils.Log.Info("UpdateQydVenueId:start, smos: %+v, appId: %s", smos, appId)

	venueModel := VenueModel{}
	// 请求趣运动api获取获取场馆基本信息
	courtInfo, err := venueModel.GetCourtInfoList(appId)
	if err != nil {
		utils.Log.Error("UpdateQydVenueId:请求趣运动api获取获取场馆基本信息失败, err: %v", err)
		return
	}
	utils.Log.Info("UpdateQydVenueId:courtInfo: %+v", courtInfo)

	// 设置数据
	smos.Qyd_venue_id = courtInfo.Venues_id

	o := orm.NewOrm()
	o.Using("whale_business_db")
	if num, err := o.Update(&smos); err == nil { // 更新成功
		utils.Log.Info("UpdateQydVenueId:update_success,受影响的行数:%d ", num)
		affectedNum = num
	} else { // 更新失败
		utils.Log.Error("UpdateQydVenueId:update_failed,err: %v", err)
	}

	return
}

// 获取appkey
func (this *SmosModel) GetAppKeyById(appId string) (string, error) {
	// 获取接口配置信息
	applicationModel := ApplicationModel{}
	application, err := applicationModel.GetQydApi()
	if err != nil { // 获取api配置信息出错
		utils.Log.Error("获取接口配置信息失败, err: %v", err) // 记录log
		return "", err
	}

	// 组织请求获取appkey的数据
	params := make(map[string]string)
	params["api_key"] = application.Api_id
	params["res_key"] = application.Api_key
	params["app_id"] = appId
	params["ver"] = application.Api_version

	// 调用趣运动接口获取appKey
	venueApi := api.Venue{}
	appKey, err := venueApi.GetAppKeyById(application.Api_address, params)
	if err != nil {
		utils.Log.Error("获取趣运动app_key失败, err: %v", err) // 记录log
		return "", errors.New("获取趣运动app_key失败")
	}

	return appKey, nil
}

// 查找smos
func (this *SmosModel) FindSmos(appId string) (smos Smos, err error) {
	o := orm.NewOrm()
	o.Using("whale_business_db")

	queryStr := "SELECT * FROM bs_smos WHERE app_id = ? AND status=1 limit 1" //return Smos
	err = o.Raw(queryStr, appId).QueryRow(&smos)

	return
}
