package helper

import (
	"crypto/md5"
	"encoding/hex"
	"gate.guanzhang.me/helper/apicode"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

// md5加密
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 判断某个字符串是否在slice中, 类似php的in_array()函数
func InSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// 写入log到文件
func Log2File(file string, content string) {

	// log 目录
	logDir := GetRootPath() + "/logs/" + time.Now().Format("20060102")

	// 创建目录
	err := os.MkdirAll(logDir, 0666)
	if err != nil {
		log.Fatalf("error can't mkdir logs: %v", err)
	}

	// log文件路径
	file = logDir + "/" + file + ".log"

	// 打开文件
	f, ferr := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if ferr != nil {
		log.Fatalf("error opening file: %v", ferr)
	}
	defer f.Close()

	// 记录错误到文件
	log.SetOutput(f)

	log.Println(" " + content)
}

// 获取入口文件的绝对路径
func GetRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// 是否是生产环境
func IsProductionEnv() bool {
	if os.Getenv("PAY_HOST") == "" { // 生产环境
		return true
	}
	// 开发或测试环境
	return false
}

// 获取配置的前缀
// 例如开发环境返回 dev:: 生产环境返回 prod::
func GetConfigPrifix() string {
	str := "::"
	if IsProductionEnv() {
		str = "prod" + str
	} else {
		str = "dev" + str
	}

	return str
}

// 手机号码格式是否正确
func IsMobile(mobile string) bool {
	var validID = regexp.MustCompile(`^\d{11}$`)
	return validID.MatchString(mobile)
}

// 生成api_sign
func MakeSign(params map[string]string, screct_key string) string {

	if len(params) == 0 {
		return ""
	}

	// 获取所有键名，用于排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	// 按键正向排序
	sort.Strings(keys)

	// 将map拼接成字符串
	var sign_data string
	for _, v := range keys {
		sign_data += v + "=" + params[v] + "&"
	}

	// 末尾添加密钥
	// 例如：action=get_order_list&client_time=1451830538&user_id=3332&4b111cc14a33b88e37e2e2934f493458
	sign_data += screct_key

	return Md5(sign_data)
}

type ApiRes struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// 初始化api返回数据
func InitApiRes() *ApiRes {
	apiRes := new(ApiRes)
	apiRes.Status = apicode.Success
	apiRes.Msg = apicode.Msg(apicode.Success)
	apiRes.Data = make(map[string]interface{})

	return apiRes
}
