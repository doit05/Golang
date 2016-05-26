package validate

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/validation"
	"strings"
)

// 接收“上传消息接口”参数结构体
type EtlUploadParams struct {
	App_id      string `form:"app_id"`
	Client_id   string `form:"client_id"`
	Client_time int    `form:"client_time"`
	Metadata    string `form:"metadata"`
	Api_sign    string `form:"api_sign"`
}

// 接收“上传客户端操作信息(operate)”参数结构体
type EtlOperateParams struct {
	App_id      string `form:"app_id"`
	Client_id   string `form:"client_id"`
	Client_time int    `form:"client_time"`
	Metadata    string `form:"metadata"`
	Api_sign    string `form:"api_sign"`
}

// 接收“获取时间消息接口”参数结构体
type EtlGetTimeParams struct {
	App_id       string `form:"app_id"`
	Client_token string `form:"client_id"`
	Client_time  int    `form:"client_time"`
	Api_sign     string `form:"api_sign"`
}

// 第一层Metadata数据
type Metadata struct {
	Metadata string
	Value    string
}

// 接收operate接口metadata数据中的Metadata部分
type MMetadata struct {
	AppId      string
	SourceName string
	ActionName string
	CreateTime string
	Format     string
	Version    string
	ClientID   string
}

// 接收operate接口metadata数据中的Value部分
type MValue struct {
	Id           string
	Code         string
	Name         string
	Operator     string
	OperatorTime string
	CreateBy     string
	CreateTime   string
}

/**
[
	{
		"Metadata":{
			"AppId":"DtiVxKZyBc3t3F35",
			"SourceName":"911001",
			"ActionName":"1",
			"CreateTime":"1463559749",
			"Format":"0",
			"Version":"soms2-alpha-v3.3.2.8769",
			"ClientID":"622A597A91D31A6C20AA1461469B198C"
		},
		"Value":{
			"Id":"d5c516e814f944aeb65d69af02a8bf19",
			"Code":"911001",
			"Name":"场地售卖",
			"Operator":"admin",
			"OperatorTime":"2016-05-18 16:22:29",
			"CreateBy":"admin",
			"CreateTime":"2016-05-18 16:22:29"
		}
	}
]
*/
// 接收operate接口metadata数据
type OperateMetadata struct {
	Metadata MMetadata
	Value    MValue
}

// 解析operate接口的Metadata数据
func ParseOperateMetadata(metadata string) (md OperateMetadata, err error) {
	var metadataArr []OperateMetadata
	err = json.Unmarshal([]byte(metadata), &metadataArr)
	if err != nil {
		return
	}

	if len(metadataArr) > 0 {
		md = metadataArr[0]
	} else {
		err = errors.New("metadata数据为空")
	}

	return md, err
}

// 解析第一层Metadata数据
func ParseMetadata(metadata string) (md Metadata, err error) {

	var metadataList []map[string]json.RawMessage

	err = json.Unmarshal([]byte(metadata), &metadataList)
	if err != nil {
		return
	}

	md.Metadata = string(metadataList[0]["Metadata"])
	md.Value = string(metadataList[0]["Value"])

	return md, err
}

// 验证参数结构体
type ValidationParams struct {
}

// 验证Metadata数据
func (v *ValidationParams) Metadata(metadata string) bool {

	if strings.Contains(metadata, "Metadata") == false { // metadata没有包含Metadata字段
		return false
	}

	if strings.Contains(metadata, "Value") == false { // metadata没有包含Value字段
		return false
	}

	_, err := ParseMetadata(metadata)

	if err != nil { // 解析第一层数据出错
		return false
	}

	return true
}

// 验证接收“上传消息接口”参数
func (v *ValidationParams) EtlUpload(params EtlUploadParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.App_id, "app_id").Message("app_id不能为空")

	valid.Required(params.Client_id, "client_id").Message("client_id不能为空")

	valid.Required(params.Client_time, "client_time").Message("client_time不能为空")
	valid.Min(params.Client_time, 1, "client_time").Message("client_time不能为小于1")

	valid.Required(params.Metadata, "metadata").Message("metadata不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else if v.Metadata(params.Metadata) == false {
		errMsg = "metadata格式错误"
	} else {
		ok = true
	}

	return ok, errMsg
}

// 验证“上传客户端操作信息(operate)”参数
func (v *ValidationParams) EtlOperate(params EtlOperateParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.App_id, "app_id").Message("app_id不能为空")

	valid.Required(params.Client_id, "client_id").Message("client_id不能为空")

	valid.Required(params.Client_time, "client_time").Message("client_time不能为空")
	valid.Min(params.Client_time, 1, "client_time").Message("client_time不能为小于1")

	valid.Required(params.Metadata, "metadata").Message("metadata不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else if v.Metadata(params.Metadata) == false {
		errMsg = "metadata格式错误"
	} else {
		ok = true
	}

	return ok, errMsg
}

// 验证接收“获取时间消息接口”参数
func (v *ValidationParams) EtlGetTime(params EtlGetTimeParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.App_id, "app_id").Message("app_id不能为空")

	valid.Required(params.Client_token, "client_token").Message("client_token不能为空")

	valid.Required(params.Client_time, "client_time").Message("client_time不能为空")
	valid.Min(params.Client_time, 1, "client_time").Message("client_time不能为小于1")

	valid.Required(params.Api_sign, "api_sign").Message("api_sign不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}
