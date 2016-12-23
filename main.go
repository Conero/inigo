// main
package main

import (
	"fmt"
)

var XHp *XHelper // 定义Brx对象
var cf *Conf     // 配置接口
// 系统常量
const (
	BR = "\r\n" // <br>
)

func main() {
	// 对象初始化
	XHp = XHelperInit()
	// 解析命令行
	fmt.Println(XHp.command())

	//go XHp.fserver("6120") // 内置文件服务器
	// 内置服务器测试
	//XHp.myServer("6100")
}
