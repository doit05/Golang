package api

import (
	"encoding/json"
	"errors"
	"gate.guanzhang.me/helper"
	"gate.guanzhang.me/utils"
	"github.com/astaxie/beego/httplib"
)

type Venue struct {
	Base
}

type ResponseData struct {
	App_id  string
	App_key string
}

type Response struct {
	Status string
	Msg    string
	Data   *ResponseData
}

type ApiResponse struct {
	Status string
	Msg    string
	Data   json.RawMessage
}

// 获取api_key
func (v *Venue) GetAppKeyById(address string, params map[string]string) (string, error) {
	resKey, exist := params["res_key"]
	if exist == false {
		return "", errors.New("缺少res_key参数")
	}
	delete(params, "res_key")

	params["action"] = "GetAppKeyById"
	params = v.AddPublicParams(params)
	params["api_sign"] = helper.MakeSign(params, resKey) // 添加api_sign

	url := v.BuildUrl(address, params)
	req := httplib.Get(url)

	utils.Log.Info("GetAppKeyById:请求接口,url: %s", url) // 记录log

	str, err := req.String()
	if err != nil { // 获取失败
		utils.Log.Error("GetAppKeyById:请求接口失败, res: %s, err: %v", str, err) // 记录log
		return "", err
	}

	response := Response{}
	jErr := json.Unmarshal([]byte(str), &response)

	if jErr != nil { // 解析json字符串失败
		utils.Log.Error("GetAppKeyById:解析json字符串失败, res: %s, err: %v", str, jErr) // 记录log
		return "", errors.New("解析json字符串失败")
	}

	if response.Status == "0000" && len(response.Data.App_key) != 0 {
		utils.Log.Info("GetAppKeyById:正确返回,Data: %+v", response.Data) // 记录log
		return response.Data.App_key, nil
	} else {
		utils.Log.Error("GetAppKeyById:获取app_key失败, res: %s", str) // 记录log
		return "", errors.New(response.Msg)
	}
}

// 用于GetCourtInfoList接口返回的Categories字段数据结构
type CourtInfoCategory struct {
	Cat_id        string
	Cat_name      string
	Is_card_order string
}

// 用于GetCourtInfoList接口返回的数据结构
type CourtInfo struct {
	App_id       string
	App_key      string
	Venues_id    string
	Name         string
	Address      string
	City_id      string
	Telephone    string
	City_name    string
	Is_soms_test string
	Categories   []*CourtInfoCategory
}

// 获取场馆基本信息
func (v *Venue) GetCourtInfoList(address string, params map[string]string) (courtInfo CourtInfo, retErr error) {
	resKey, exist := params["res_key"]
	if exist == false {
		retErr = errors.New("缺少res_key参数")
		return
	}

	delete(params, "res_key")

	params["action"] = "GetCourtInfoList"
	params = v.AddPublicParams(params)
	params["api_sign"] = helper.MakeSign(params, resKey) // 添加api_sign

	url := v.BuildUrl(address, params)
	req := httplib.Get(url)

	utils.Log.Info("GetCourtInfoList:请求接口,url: %s", url) // 记录log

	str, err := req.String()
	if err != nil { // 获取失败
		utils.Log.Error("GetCourtInfoList:请求接口失败, res: %s, err: %v", str, err) // 记录log
		retErr = err
		return
	}

	response := ApiResponse{}
	jErr := json.Unmarshal([]byte(str), &response)

	if jErr != nil { // 解析json字符串失败
		utils.Log.Error("GetCourtInfoList:解析json字符串失败, res: %s, err: %v", str, jErr) // 记录log
		retErr = errors.New("解析json字符串失败")
		return
	}

	if response.Status == "0000" { // 请求接口成功
		jErr2 := json.Unmarshal(response.Data, &courtInfo)
		if jErr2 == nil { // data字段解析成功
			utils.Log.Info("GetCourtInfoList:正确返回,Data: %+v", courtInfo) // 记录log
			return
		} else {
			utils.Log.Error("GetCourtInfoList:解析data字段数据失败, res: %s, err: %v", str, jErr) // 记录log
			retErr = errors.New("解析data字段数据失败")
			return
		}
	} else {
		utils.Log.Error("GetCourtInfoList:接口返回失败, err: %s", response.Msg) // 记录log
		retErr = errors.New(response.Msg)
		return
	}
}
