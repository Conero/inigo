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
	p.Set("test", 5)
	fmt.Println(p.GetData())
}

func TestNewParserRong(t *testing.T) {
	rong := NewParser(nil, "rong")
	fmt.Println(rong)
}

func TestNewParserIni(t *testing.T) {
	ini := NewParser(nil, "ini")
	ini.Set("test", 8).
		Set("name", "Full")
	fmt.Println(ini)
}
