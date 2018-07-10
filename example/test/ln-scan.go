// 文件扫描
// 2018年7月10日 星期二
package main

import (
	"../../src/ini"
	"fmt"
)
func main() {
	scan1()
}


func scan1(){
	t1 := ini.NewLnRer("./resource/scan1.ini")
	t1.Scan(func(line string) {
		//fmt.Println(line)
		fmt.Print(line)
	})
}
