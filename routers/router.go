package routers

import (
	"gate.guanzhang.me/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 上传数据
	beego.Router("/etl/upload", &controllers.EtlController{}, "post:Upload")

	// 上传客户端操作信息
	beego.Router("/etl/operate", &controllers.EtlController{}, "post:Operate")

	// 获取上传时间
	beego.Router("/etl/get_time", &controllers.EtlController{}, "post:GetTime")
}
