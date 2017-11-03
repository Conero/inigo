/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月28日 星期六
 * @ini 变量列表
 */

package ini

// 系统常量
const (
	AUTHOR    = "Joshua Conero" // @author 作者
	VERSION   = "1.0.4"         // @version	版本号
	NAME      = "go ini 文件解析库"  // @name 名称
	START     = "20171103"      // @start 开始时间
	COPYRIGHT = "@Conero"       // @copyright 版权
)

// ini-parse 设置
var IniParseSettings map[string]string = map[string]string{
	"equal":          "=",                    // 等号符
	"comment":        "#|;",                  // 注释符号
	"mcomment1":      "'''",                  // 多行注释 - 开始
	"mcomment2":      "'''",                  // 多行注释 - 结束
	"limiter":        ",",                    // 分隔符
	"scope1":         "{",                    // 作用域 - 开始
	"scope2":         "}",                    // 作用域 - 结束
	"reg_comment":    "^[#;]",                // 注释符号
	"reg_section":    "^\\[[^\\[^\\]}]*\\]$", // 是否为章节正则检测
	"reg_section_sg": "(\\[)|(\\])",          // 章节标点符号处理
	"reg_scope":      "\\{[^\\{^\\}]*\\}",    // 作用域开始于结束正则
	//"reg_scope_sg": "$\\{[^\\{^\\}]*\\}^", // 单行作用域解析
	"reg_scope_sg": "^\\{.*\\}$", // 单行作用域解析
}

// 转移字符解析
var TranStrMap map[string]string = map[string]string{
	`\,`: "_JC__COMMA", // 逗号转移符
	`\{`: "_L__BRACE",  // 左大括弧号
	`\}`: "_R__BRACE",  // 右大括弧号
	`\=`: "_JC__EQUAL", // 等于符号转移替代
}
