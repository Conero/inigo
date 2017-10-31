// main
package main

import (
	"fmt"

	ict "../../ini-go" // 引用当前版本
	ini "../dist/v0.1/ini"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(ini.NAME)
	fmt.Println(ini.VERSION)
	// testIniFunc()
	//	buildInitFunc()
	InitCurTest()
}
func InitCurTest() {
	conf := ict.Open("./test.ini")
	conf.Parse()
	fmt.Println(conf.Cdt)
}

// test.ini 文件
func testIniFunc() {
	testIni := ini.Open("./test.ini")
	testIni.Parse()
	fmt.Println(testIni.Parse())
	testIni.Compile() // ini 编译
	fmt.Println(testIni)
	listen := testIni.GetString("listen")
	listenArry := testIni.GetArray("listen")
	fmt.Println(listen, listenArry)
	mul := testIni.GetString("mul")
	mulArray := testIni.GetArray("mul")
	fmt.Println(mul, mulArray)
}
func buildInitFunc() {
	//	inibuild := ini.EmptyBuilder()
	inibuild := ini.MapBuilder(map[string]string{})
	inibuild.SetPath("./build.ini")
	inibuild.SetValue("author", ini.AUTHOR)
	inibuild.SetValue("name", ini.NAME)
	fmt.Println(inibuild.Save())
}
