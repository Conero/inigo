// main.go
package main

import (
	"fmt"
	"reflect"

	ini "../../dist/v0.2.Ini"
	. "../common"
	u "../unit"
)

func main() {
	//BaseTest()
	//testv02()
	testv02Full()
	//interfaceTest()
}

// 基本测试
func BaseTest() {
	fmt.Println(ini.GetVal("conf.equal"))
	fmt.Println(ini.GetVmap("conf"))
	fmt.Println(ini.GetVmap("conf6"))
}

// v0.2test.ini 测试
func testv02() {
	//	scopeMap := make(map[string]interface{})
	//	testMap := make(map[string]interface{})
	//	testMap["yang"] = 8
	//	testMap["ttt"] = []string{"55", "666"}
	//	fmt.Println(len(scopeMap), len(testMap))

	Conf := ini.Open("./v0.2test.ini")
	Conf.Parse()
	fmt.Println(Conf.Cdt)
	fmt.Println(Conf.ToJson())
	//	fmt.Println(ini.NAME)
	PutContent("./v0.2test.json", Conf.ToJson())
}

func testv02Full() {
	Conf := ini.Open("./v0.2test.ini")
	Conf.FullParse()
	fmt.Println(Conf.Cdt)
	//	return
	fmt.Println(Conf.ToJson())
	//	fmt.Println(ini.NAME)
	PutContent("./v0.2test.full.json", Conf.ToJson())
	// 数据对比
	fmt.Println(Conf.Get("tavelZh"))
	fmt.Println(Conf.Get("likes"))
	v, has := Conf.Get("test")
	u.PutWhenBool(has, v)
	fmt.Println(Conf.Get("place"))
}

// interface {} 类型学习与研究
func interfaceTest() {
	type Any interface{}
	var test Any

	// 整形
	test = 5
	test = test.(int) + 48
	AnyToType(test)
	fmt.Println(test)

	// 字符串
	test = "What's F**k"
	fmt.Println(test)

	// map[string]string
	test = map[string]string{
		"emma": "it's my wife",
	}
	Jk := test.(map[string]string)
	fmt.Println(test)
	Jk["lee"] = "lee in usa"
	AnyToType(Jk)
	fmt.Println(test)

	// map[string]interface{}
	wmt := make(map[string]interface{})
	wmt["int"] = 5
	wmt["str"] = "string"
	test = wmt

	AnyToType(test)
	fmt.Println(test)
}

func AnyToType(value interface{}) {
	fmt.Println(value, "-->", reflect.TypeOf(value))
}
