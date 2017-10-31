// main
package main

import (
	"fmt"

	ini "../../src"
)

type MainTest struct{}

func main() {
	p := &MainTest{}
	//p.BaseTest()

	p.ParseIni()
}

// 基本测试
func (m *MainTest) BaseTest() {
	fmt.Println(ini.NAME)
	fmt.Println("Hello World!")
}

func (m *MainTest) ParseIni() {
	Conf := ini.Open("./test0.3.0.ini")
	Conf.Parse()
	//Conf.FullParse()
	fmt.Println(Conf.Cdt)
	fmt.Println(Conf.ToJson())
}
