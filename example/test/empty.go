/**
	2018年7月10日 星期二
	空 ini 文件处理
 */
package main

import (
	"fmt"
	"../../src/ini"
)

func main() {
	vini := ini.Open("./resource/uymas.ini")
	fmt.Println(vini.DataQueue)
	fmt.Println(2)
}
