package main

import (
	"fmt"
	. "../../pkg/ini"
)

type IniTest struct {
}
func (it *IniTest) Console(){
	it.linerTest()
}

func (it *IniTest) linerTest(){
	line := &Tester{}
	//fmt.Println(line)
	line.LinerTest()
	fmt.Println(line)
}

func main() {
	//fmt.Println(VERSION)
	// 控制台
	(&IniTest{}).Console()
}
