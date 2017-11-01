/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月28日 星期六
 * @ini 变量列表
 */

package rong

// 系统常量
const (
	AUTHOR    = "Joshua Conero" // @author 作者
	VERSION   = "1.0.0"         // @version	版本号
	NAME      = "go ini 文件解析库"  // @name 名称
	START     = "2017-10-28"    // @start 开始时间
	COPYRIGHT = "@Conero"       // @copyright 版权
)

// ini-parse 设置
var IniParseSettings map[string]string = map[string]string{
	"equal":          "=",                    // 等号符
	"reg_comment":    "^[#;]",                // 注释符号
	"reg_section":    "^\\[[^\\[^\\]}]*\\]$", // 是否为章节正则检测
	"reg_section_sg": "(\\[)|(\\])",          // 章节标点符号处理
}
