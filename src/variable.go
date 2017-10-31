/**
 * 全局常量
 * 2017年3月16日 星期四
 */
package ini

// 基本信息
const (
	AUTHOR    = "Joshua Conero" // @author 作者
	VERSION   = "0.3.0"         // @version	版本号
	NAME      = "go ini 文件解析库"  // @name 名称
	START     = "2017-01-19"    // @start 开始时间
	COPYRIGHT = "@Conero"       // @copyright 版权
)

type setType map[string]string

// 配合值 - 常量
var Setting = map[string]setType{
	"conf": map[string]string{
		"equal":     "=",   // 等号符
		"comment":   "#|;", // 注释符号
		"mcomment1": "'''", // 多行注释 - 开始
		"mcomment2": "'''", // 多行注释 - 结束
		"limiter":   ",",   // 分隔符
		"scope1":    "{",   // 作用域 - 开始
		"scope2":    "}",   // 作用域 - 结束
	},
}
