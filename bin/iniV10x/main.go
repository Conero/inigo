package main

import (
	"../../pkg/ini"
	"../../pkg/rong"
	"fmt"
)

// ini-test
func testV10x() {
	// 相对地址，在包里无法读取文件
	//test := ini.Open("./test-v1.0.x.ini")
	test := ini.Open("D:/Joshua/Active/go/ini-go/bin/iniV10x/test-v1.0.x.ini")
	if !test.IsSuccess {
		fmt.Println(test.FailMsg)
	}
	// 输出解析后的对象
	fmt.Println(test.DataQueue)
	fmt.Println(test.ToJsonString())
}

// Rong test
func RongTest() {
	su := rong.Open("D:/Joshua/Active/go/ini-go/bin/iniV10x/rong-v1.0.x.ini")
	if !su.IsSuccess {
		fmt.Println(su.FailMsg)
		return
	}
	fmt.Println(su.DataQueue)
	fmt.Println(su.GetString("author"))
	fmt.Println(su.GetString("rong.lastname"))

}

// 入口测试
func main() {
	fmt.Println(ini.VERSION)
	testV10x()
	//RongTest()
}
