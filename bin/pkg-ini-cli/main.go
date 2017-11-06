/**
pkg/ini 命令行测试时数据
*/
package main

import (
	"../../pkg/ini"
	"../../pkg/rong"
	"../../pkg/running"
	"fmt"
	"os"
	"regexp"
	"strings"
	"io/ioutil"
)

const (
	Name    = "PkgIniCli"
	Version = "0.0.2"
	Build   = "20171106"
	Start   = "20171105"
	Author  = "Joshua Conero"
)

//命令行层次处理
type Cli struct {
	Action      string   // 命名
	Param       []string // 参数
	ParamArgsMk bool     // 参数标识
}

// 路由器正在匹配处理 => reg => action
var CliRouterMapRegs map[string]string = map[string]string{
	`^[-]*(\?)|(help)|(HELP)`: "help-action", // help
	`^[-]*(ini)|(INI)`:        "ini-action",  // ini
	`^[-]*(rong)|(Rong)`:      "rong-action", // rong
}

// 根据 os.Args 解析参数
func (cli *Cli) Init() {
	if cli.ParamArgsMk {
		return
	}
	args := os.Args
	if len(args) > 1 {
		cli.Action = args[1]
		cli.ParamArgsMk = true
	}
	if len(args) > 2 {
		cli.Param = args[2:]
	}
}

func (cli *Cli) regToAction(reg string) {
	switch reg {
	case "help-action":
		cli.HelpAction()
	case "ini-action":
		cli.IniAction()
	case "rong-action":
		cli.RongAction()
	}
}

// 命令行路由器
func (cli *Cli) Router() {
	// 命令行初始化
	cli.Init()
	CliMapMk := false
	if len(cli.Action) > 0 {
		// 正在匹配
		for regs, value := range CliRouterMapRegs {
			if matched, _ := regexp.MatchString(regs, cli.Action); matched {
				CliMapMk = true
				cli.regToAction(value)
				break
			}
		}
	}
	if !CliMapMk {
		cli.IndexAction()
	}
}
// 获取参数
func (cli *Cli) getParamByIdx(args ...interface{})string{
	idx := 0
	def := ""
	if len(args) > 0{
		idx = args[0].(int)
	}
	if len(args) > 1{
		def = args[1].(string)
	}
	if idx > 0 && idx <= len(cli.Param){
		def = strings.TrimSpace(cli.Param[idx-1])
	}
	return def
}
// 获取参数 自动解析key
func (cli *Cli) getParamKV(args ...interface{})(string, string){
	sP := cli.getParamByIdx(args...)
	key, value := sP, ""
	if len(sP) > 0{
		if idx := strings.Index(sP, "--");idx > -1{
			key = sP[0:idx]
			value = sP[idx+2:]
		}else if idx := strings.Index(sP, "=");idx > -1{
			key = sP[0:idx]
			value = sP[idx+1:]
		}
	}
	return key,value
}
// Index
func (cli *Cli) IndexAction() {
	fmt.Println(`
	欢迎使用 ini-go 库:)-
		* ` + Name + `	pkg/ini 测试命令行程序
		* 版本号	v` + Version + `(` + Build + `)
		* 开始日期	` + Start + `
		* 作者		` + Author + `
	pkg/ini@` + ini.VERSION + `-` + ini.BUILD + `
	pkg/rong@` + rong.VERSION + `-` + rong.BUILD + `
	`)
}

// Help 帮助文档
func (cli *Cli) HelpAction() {
	fmt.Println(`
	--ini		ini 包文件解析测试
		. <ini-file>  <json[=jsonname]>	解析文件输出为json 字符串
		. <ini-file>  <get=key1.key2.key3> 读取解析成功以后的值
	--roung		rong 包文件文件解析测试
	--help		程序帮助
	`)
}

// Ini 文件测试
func (cli *Cli) IniAction() {
	name := cli.getParamByIdx(1, "")
	if "" != name{
		show(name+"文件正在等解析.")
		is := ini.Open(name)
		Rt := running.CreateTimer()
		if !is.IsSuccess{
			show(is.FailMsg, "F")
			return
		}
		show("文件解析完成： "+Rt.GetSecString(4)+"s, 共"+is.File.GetLineString()+"行")
		oKey, oName := cli.getParamKV(2)
		oKey = strings.ToLower(oKey)
		// 保存解析文件为 json 文件格式
		if oKey == "json"{
			if oName != ""{
				oName = "./ini-"+oName+".json"
			}else{
				fName, _ := fileSplit(name)
				oName = "./ini-"+fName+"-v"+ini.VERSION+".json"
			}
			writeToFile(oName, is.ToJsonString())
		}else if oKey == "get"{
			// 输出整解析后的内容
			if oName == ""{
				fmt.Println("   ", is.DataQueue)
			}else{
				has, dd := is.Get(oName)
				if has{
					fmt.Println("  -->", dd)
				}else {
					fmt.Println("  ", oName+"值不存在！")
				}
			}
		}
	}
}

// Rong 文件测试
func (cli *Cli) RongAction() {
}

// 测试命程序
// 参数：  args => [action, param]
func (cli *Cli) TestArgs(args []string) *Cli {
	cli.Init()
	if !cli.ParamArgsMk {
		// action
		if len(args) > 0 {
			cli.ParamArgsMk = true
			cli.Action = args[0]
		}
		if len(args) > 1 {
			cli.Param = args[1:]
		}
	}
	return cli
}

func main() {
	(&Cli{}).Router()
	/*
	(&Cli{}).TestArgs([]string{
		"ini",
		"out/source/mutiline.ini",
		"json--ini",
		//"json=ini",
		}).Router()
		*/
}
/********************公共函数******************/
// 格式化输出 => string msg,  string type
func show(args ...interface{})  {
	str := ""
	if len(args) > 0{
		str = args[0].(string)
	}
	stype := ":)"
	if len(args) > 1{
		st := args[1].(string)
		if st == "F"{	// fail 失败
			stype = "-(:"
		}
	}
	fmt.Println(" go-ini "+stype+"  "+str)
}
// 文件是否存在
func file_exist(filename string) bool{
	exist := false
	if _, err := os.Stat(filename); os.IsExist(err){
		exist = true
	}
	return  exist
}

// 文件写入测试
func writeToFile(path, content string) {
	ioutil.WriteFile(path, []byte(content), 0666)
}

func fileSplit(filename string) (string, string) {
	name := ""
	ext := ""
	filename = strings.Replace(filename, "\\", "/", -1)
	fT := strings.Split(filename , "/")
	lastArs := fT[len(fT)-1]
	pointIdx := strings.Index(lastArs, ".")
	name = lastArs[0:pointIdx]
	ext = lastArs[pointIdx+1:]
	return name, ext
}