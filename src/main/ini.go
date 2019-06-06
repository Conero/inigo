package main

import (
	"../action"
	"../vars"
	"fmt"
	"github.com/conero/uymas/bin"
)

/**
 * @DATE        2019/6/6
 * @NAME        Joshua Conero
 * @DESCRIPIT   inigo 语言包的命令行助手
 **/

// 欢迎语
func welcome() {
	Br := "\r\n"
	welcomeStr :=
		"" +
			" 欢迎使用 [" + vars.Name + "]" + Br +
			" version " + vars.Version + "/" + vars.Release + Br +
			" Since " + vars.Since

	fmt.Println(welcomeStr)
}

// 主入口
func main() {
	bin.RegisterApps(map[string]interface{}{
		"ini": &action.IniAction{},
	})
	bin.EmptyFunc(func() {
		welcome()
	})
	bin.Run()
}
