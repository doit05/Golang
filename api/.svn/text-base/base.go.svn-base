package api

import (
	"fmt"
	"gate.guanzhang.me/helper"
	"net/url"
)

type Base struct {
}

// 添加公共参数
func (b *Base) AddPublicParams(params map[string]string) map[string]string {
	params["client_time"] = fmt.Sprintf("%d", helper.GetTimestamp())
	params["utm_medium"] = "gate.guanzhang.me"

	return params
}

// 拼接url
// http://court.qydw.net?action=GetCourtInfoList&api_key=94f2309f7fd56a61&api_sign=c8fb756b1c403b37d0ed64785a9ff2fa&app_id=DtiVxKZyBc3t3F35&client_time=1463479031&utm_medium=gate.guanzhang.me&ver=1.0
func (b *Base) BuildUrl(path string, params map[string]string) string {
	v := url.Values{}

	for key, val := range params {
		v.Add(key, val)
	}

	url := path + "?" + v.Encode()

	return url
}
