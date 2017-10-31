// 共享库或参数(全局变量或函数)
package common

// 接口参数
var Config = map[string]string{
	"config_dir": "./config/", // 配置文件目录
	"cache_dir":  "./cache/",  // 缓存库
}
var Packages = map[string]string{
	"common": "全局参数共享包",
	"ini":    "go-ini 库",
	"zmapp":  "zmapp 引擎",
}

// -> 实际项目中 新建一个 project.go 用于存放如下常量
const (
	AUTHOR    = "Joshua Conero" // @author 作者
	VERSION   = "0.0.5"         // @version	版本号
	NAME      = "zmapp框架非内嵌式插件" // @name 名称
	START     = "2017-03-16"    // @start 开始时间
	COPYRIGHT = "@Conero"       // @copyright 版权
)

// 根据不同版本号聊区分
// 基本信息 - 使用常量说明
var About = map[string]string{
	"author":    "Joshua Conero",  // @author 作者
	"version":   "0.0.1 - 170330", // @version	版本号
	"name":      "go - 程序公共处理函数库", // @name 名称
	"start":     "20170330",       // @start 开始时间
	"copyright": "@Conero",        // @copyright 版权
}
