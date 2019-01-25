package inigo

import (
	"fmt"
	"testing"
)

// @Date：   2018/8/19 0019 15:03
// @Author:  Joshua Conero
// @Name:    ini test
func TestNewParser(t *testing.T) {
	p := NewParser()

	// int
	p.Set("test", 5)
	//fmt.Println(p.GetData())
	has, value := p.Get("test")
	if !has || value.(int) != 5 {
		t.Fatal("[\"test\"=5] 设置值无效")
	}

	// bool
	p.Set("bool", true)
	//fmt.Println(p.GetData())
	has, value = p.Get("bool")
	if !has || value.(bool) != true {
		t.Fatal("[\"bool\"=true] 设置值无效")
	}

}

func TestNewParserRong(t *testing.T) {
	rong := NewParser(nil, SupportNameRong)
	fmt.Println(rong)
	fmt.Println(rong.Driver())

	if rong.Driver() != SupportNameRong{
		t.Fatal("Driver 默认生成无效！")
	}
}

func TestNewParserIni(t *testing.T) {
	ini := NewParser(nil, SupportNameIni)
	ini.Set("test", 8).
		Set("name", "Full")
	fmt.Println(ini)
	if ini.Driver() != SupportNameIni{
		t.Fatal("Driver 默认生成无效！")
	}
}
