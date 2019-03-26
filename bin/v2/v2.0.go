package main

import (
	"fmt"
	"github.com/conero/inigo"
	"time"
)

type Cs struct {
}

func (c Cs) NewBase() {
	ini := inigo.NewParser()
	ini.OpenFile("./base.ini")

	fmt.Println(ini.IsValid())
	fmt.Println(ini.GetData())

	fmt.Println(ini.Section("data", "name"))
}

func (c Cs) newRong() {
	rong := inigo.NewParser("rong")
	rong.OpenFile("./rong.ini")
	fmt.Println(rong.GetData())
}

func (c Cs) createBase()  {
	base := inigo.NewParser()
	base.Set("name", time.Now().String())
	fmt.Println(base.Driver())
	fmt.Println(base.GetData())
	base.SaveAsFile("./createBase.ini")
}

func main() {
	var cs Cs
	//cs.NewBase()
	//cs.newRong()
	cs.createBase()
}
