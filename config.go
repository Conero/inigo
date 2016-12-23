// config
package main

import (
	"fmt"
	"os"
)

// 配置文件处理类
type Conf struct{}

// 配置类主方法
func (C *Conf) run(v string) string {
	cfs := [...]string{"/XHelper/config.json", "/XHelper/config.xml", "/XHelper/config.ini"}
	var fname, ret string
	switch v {
	case "json": // 生成json 格式配置文件
		put_content("./XHelper/config.json", "")
	case "xml": // 生成xml配置文件
		put_content("./XHelper/config.xml", "")
	case "delete": // 删除所有已经生成的配置文件
		for _, v := range cfs {
			fname = "." + v
			if hasFile(fname) {
				err := os.Remove(fname)
				if err == nil {
					ret += "\r\n删除文件：" + v
				}
			}
		}
		if ret == "" {
			ret = "未发现任何配置文件，删除操作失败"
		}
		return ret
	case "?":
		for _, v := range cfs {
			fname = "." + v
			if hasFile(fname) {
				ret += BR + v
			}
		}
		if ret == "" {
			ret = "您没有任何配置文件"
		} else {
			ret = "\r\n您的文件列表：" + ret
		}
		return ret
	default: // 生成ini配置文件-默认
		v = "ini"
		put_content("./XHelper/config.ini", "")
	}
	return v

}
func config_guaqi() {
	fmt.Println("Hello World!")
}
