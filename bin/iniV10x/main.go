package main

import (
	"../../pkg/ini"
	"../../pkg/rong"
	"../../pkg/running"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// ini-test
func testV10x() {
	// 相对地址，在包里无法读取文件
	//test := ini.Open("./test-v1.0.x.ini")
	rt := running.CreateTimer()
	test := ini.Open("D:/Joshua/Active/go/ini-go/bin/iniV10x/test-v1.0.x.ini")
	if !test.IsSuccess {
		fmt.Println(test.FailMsg)
	}
	// 运行秒数
	fmt.Println("运行秒数(s): ", rt.GetSec())
	// 输出解析后的对象
	//fmt.Println(test.DataQueue)
	//fmt.Println(test.ToJsonString())
	//writeToFile(ini.VERSION+"-test-to.json", test.ToJsonString())
	fmt.Println(test.Get("mapv10.map9.map8.name"))
	writeToFile(ini.VERSION+"-test-ini-creator.ini", ini.MkCreator(test.DataQueue).ToIniString())
	writeToFile(ini.VERSION+"-test-ini-creator.json", ini.MkCreator(test.DataQueue).ToJsonString())

	// 生成器测试
	crt := ini.MkContainer()
	crt.Set("name", "pkg/ini 文件生成器").
		Set("date", time.Now().String()).
		Set("test", "test").
		Set("map", map[string]string{
			"name": "map 格式",
			"author": "Joshua Conero",
		}).
		Set("map_class2", map[string]interface{}{
		"map": map[string]interface{}{
				"name": "map 格式",
				"author": "Joshua Conero",
			},
		}).
		Set("array", []string{"I", "LOVE", "U", "Sus"})

	writeToFile(ini.VERSION+"-self-ini-creator.ini", crt.ToIniString())
	writeToFile(ini.VERSION+"-self-ini-creator.json", crt.ToJsonString())

	// 运行秒数
	fmt.Println("#2运行秒数(s): ", rt.GetSec())
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
	//fmt.Println(ini.VERSION)
	testV10x()
	//RongTest()
}

// 文件写入测试
func writeToFile(name, content string) {
	dir := "./tmps/"
	/*
		fi, _ := os.Lstat(dir)
		if !fi.IsDir(){
			os.Mkdir(dir, os.ModeDir)
		}
	*/
	//os.Mkdir(dir, os.ModeDir)
	//os.Mkdir(dir, 0666)
	//fmt.Println(os.Mkdir(dir, 0666))
	os.Mkdir(dir, os.ModePerm)

	path := dir + name
	ioutil.WriteFile(path, []byte(content), 0666)
}
