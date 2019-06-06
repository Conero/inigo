package main

import (
	"../action"
	"../vars"
	"github.com/conero/uymas/bin"
)

/**
 * @DATE        2019/6/6
 * @NAME        Joshua Conero
 * @DESCRIPIT   inigo 语言包的命令行助手
 **/

// 主入口
func main() {
	cfg := vars.Cfg()

	// 空命令时，不破坏系统内置的功能
	if bin.IsEmptyCmd() {
		// 以 [$ ini] 命令打开
		open_use_ini := cfg.GetDef("open_use_ini", false)
		//fmt.Println(cfg.Raw("open_use_ini"), cfg.IsValid())
		if v, is := open_use_ini.(bool); is && v {
			bin.InjectArgs("ini")
		}
	}

	bin.RegisterApps(map[string]interface{}{
		"ini": &action.IniAction{},
	})
	bin.EmptyFunc(func() {
		vars.Welcome()
	})
	bin.Run()
}
