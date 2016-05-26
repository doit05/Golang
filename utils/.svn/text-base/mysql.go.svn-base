package utils

import (
	"gate.guanzhang.me/helper"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	configPrifix := helper.GetConfigPrifix() // 获取配置前缀

	// 默认数据库(whale_di_db)
	err := orm.RegisterDataBase("default", "mysql", beego.AppConfig.String(configPrifix+"whale_di_db"))
	if err != nil {
		panic("设置whale_di_db数据库配置失败, err: " + err.Error())
	}

	// business数据库(whale_business_db)
	err2 := orm.RegisterDataBase("whale_business_db", "mysql", beego.AppConfig.String(configPrifix+"whale_business_db"))
	if err2 != nil {
		panic("设置whale_business_db数据库配置失败, err: " + err2.Error())
	}

	// openplat数据库(whale_openplat_db)
	err3 := orm.RegisterDataBase("whale_openplat_db", "mysql", beego.AppConfig.String(configPrifix+"whale_openplat_db"))
	if err3 != nil {
		panic("设置whale_openplat_db数据库配置失败, err: " + err3.Error())
	}

}
