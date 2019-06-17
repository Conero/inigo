package action

import (
	"../vars"
	"github.com/conero/inigo"
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
