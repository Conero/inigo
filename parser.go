package inigo

// @Date：   2018/8/19 0019 10:54
// @Author:  Joshua Conero
// @Name:    解析器

type Parser interface {
	// 读取参数
	Get(key string) (bool, interface{})
	HasKey(key string) bool

	// 支持多级数据访问，获取元素数据
	Raw(key string) string

	// 获取参数: key, value(nil), default
	Value(params ...interface{}) interface{}

	// 节处理
	GetAllSection() []string
	// 参数格式
	//		1.     fun Section(section, key string) 	二级访问
	//		2.     fun Section(format string) 			点操作
	Section(params ...interface{}) interface{}


	// 获取数据返回 nil
	GetData() map[interface{}]interface{}


	Set(key string, value interface{}) Parser

	// 文件检测有效性
	IsValid() bool
	OpenFile(filename string) Parser
	ReadStr(content string) Parser
}
