package action

import (
	"github.com/conero/inigo"
	"github.com/conero/inigo/cmd/ini/vars"
)

/**
 * @DATE        2019/6/17
 * @NAME        Joshua Conero
 * @DESCRIPIT   action 默认应用
 **/

var gCfg inigo.Parser

// 用用初始化
func init() {
	gCfg = vars.Cfg()
}
