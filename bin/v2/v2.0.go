package main

import (
	"fmt"
	"inigo"
)

type Cs struct {
}

func (c Cs) NewBase() {
	ini := inigo.NewParser()
	ini.OpenFile("./base.ini")

	//fmt.Println(ini.GetData())

	fmt.Println(ini.Section("data", "name"))
}

func (c Cs) newRong() {
	rong := inigo.NewParser("rong")
	rong.OpenFile("./rong.ini")
	fmt.Println(rong.GetData())
}
func main() {
	var cs Cs
	cs.NewBase()
	//cs.newRong()
}
