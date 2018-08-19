package inigo

import (
	"testing"
	"fmt"
)

// @Date：   2018/8/19 0019 15:03
// @Author:  Joshua Conero
// @Name:    ini test

func TestNewParser(t *testing.T) {
	p := NewParser()
	p.Set("test", 5)
	fmt.Println(p.GetData())

	rong := NewParser(nil, "rong")
	fmt.Println(rong)
}

