package utils

import (
	"gate.guanzhang.me/helper"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
	"encoding/json"

)

var Redis cache.Cache

func init() {
	configPrifix := helper.GetConfigPrifix() // 获取配置前缀

	bm, err := cache.NewCache("redis", beego.AppConfig.String(configPrifix+"redisconn"))

	if err == nil {
		Redis = bm
	} else {
		panic("redis连接失败: " + err.Error())
	}
}

//将对象存入Cache
func SetCache(cacheKey string, obj interface{}, timeout time.Duration) (err error) {

	jsonData, jErr := json.Marshal(obj)
	if jErr == nil {
		pErr := Redis.Put(cacheKey, jsonData, timeout)
		if pErr != nil {
			err = pErr
			Log.Warn("添加redis缓存失败, jsonData: %s", jsonData)
		}
	} else {
		err = jErr
		Log.Warn("json编码失败, smos: %+v, err: %v", obj, jErr)
	}
	return
}

func SetCacheOneHour(cacheKey string, obj interface{}) (err error) {
	return SetCache(cacheKey, obj, time.Second * 3600)
}
