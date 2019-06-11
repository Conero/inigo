package action

import (
	"../vars"
	"bufio"
	"fmt"
	"github.com/conero/inigo"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/fs"
	"github.com/conero/uymas/util"
	"os"
	"path"
	"regexp"
	"strings"
)

/**
 * @DATE        2019/6/6
 * @NAME        Joshua Conero
 * @DESCRIPIT   描述 descript
 **/

type IniAction struct {
	bin.Command
}

// 交互时命令行入口
func (a *IniAction) Run() {
	vars.Welcome()

	input := bufio.NewScanner(os.Stdin)
	// 文本输入前缀
	fmt.Print("$ ")
	helper := &iniHelper{}
	// 初始化
	helper.oIRQueueManger.init()
	// 输入扫描
	for input.Scan() {
		text := input.Text()
		if text == vars.CmdExit {
			fmt.Println(" 欢迎下次使用 [inigo]， 再见！祝您愉快")
			break
		}
		// 助手入口
		helper.init(text)
		// 文本输入前缀
		fmt.Print("$ ")
	}
}

// 唯一键： filename, name|alias
// ini 文件打开资源
// 用于缓存数据
type openIniResource struct {
	ini      inigo.Parser
	useTime  string // 花费时间
	name     string // 加载的文件命名
	alias    string // 别名
	filename string // 文件名称

}

// oIRQueue 管理器
type oIRQueueManger struct {
	oIR      []openIniResource
	curIdx   int    // 当前的索引
	curFname string // 当前所属文件名
	curName  string // 当前显示文件名
}

// 初始化
func (oq *oIRQueueManger) init() {
	if oq.oIR == nil {
		oq.oIR = []openIniResource{}
	}
}

// filename, name(alias) 的唯一性判断
func (oq *oIRQueueManger) hasQueue(filename, name string) bool {
	exist := false
	for _, q := range oq.oIR {
		if filename == q.filename {
			if q.alias != "" && name == q.alias {
				exist = true
				break
			} else if q.alias == "" && name == q.name {
				exist = true
				break
			}
		}
	}
	return exist
}

// 获取的数据
func (op *oIRQueueManger) hasAndChangeByName(name string) bool {
	success := false
	for i, q := range op.oIR {
		if q.alias == name {
			success = true
		} else if q.name == name {
			success = true
		} else if fmt.Sprintf("%v%v", q.name, i) == name {
			success = true
		}
		if success {
			op.curName = q.name
			op.curFname = q.filename
			op.curIdx = i
			break
		}
	}
	return success
}

// @TODO needTodos 实现 $ open 不受相同文件名影响；支持别名/文件名 等唯一性不同区分

// 交互类
// ini 助手
type iniHelper struct {
	text string
	cmds []string
	xa   *bin.App
	//resource map[string]inigo.Parser // ini 资源集合
	//resource oIRQueue // ini 资源集合
	oIRQueueManger
	curName string // 当前的 ini 文件
}

// 初始化
func (h *iniHelper) init(text string) {
	// 字符串清洗
	reg := regexp.MustCompile("[\\s]{2,}")
	text = reg.ReplaceAllString(text, " ")

	h.text = text
	h.cmds = strings.Split(strings.TrimSpace(text), " ")

	// 命令分发
	h.router()
}

// 通用的格式化输出
func (h *iniHelper) output(params ...string) {
	br := "\r\n"
	outs := "" +
		h.curName + ">>" + br

	for _, s := range params {
		if s == "" {
			continue
		}
		outs += "  " + s + br
	}
	fmt.Println(outs)
	fmt.Println()
}

// 使用 bin 包启命令行程序
// 且进行注册
func (h *iniHelper) router() {
	bin.InjectArgs(h.cmds...)

	// 新的命令行程序
	h.xa = bin.GetApp()

	// 命令接口
	bin.RegisterFunc("open", h.cmdOpen)
	bin.RegisterFunc("use", h.cmdUse)
	bin.RegisterFunc("list", h.cmdList)
	bin.RegisterFunc("get", h.cmdGet)
	bin.RegisterFunc("help", h.cmdHelp)
	bin.EmptyFunc(h.cmdEmpty)
	bin.UnfindFunc(h.cmdUnfind)

	bin.Run()
}

// 打开 ini 文件
func (h *iniHelper) cmdOpen() {
	xa := h.xa
	filename := xa.Next(xa.Command)
	filename = fs.StdPathName(filename)
	if filename != "" {
		rtMk := util.SecCallStr()
		ini := inigo.NewParser()
		ini.OpenFile(filename)
		if !ini.IsValid() {
			h.output("打开错误："+ini.ErrorMsg(),
				"文件："+filename,
			)
		} else {
			_, name := path.Split(filename)
			nameQue := strings.Split(name, ".")
			vLen := len(nameQue)
			if vLen > 1 {
				name = strings.Join(nameQue[0:vLen-1], ".")
			}
			h.curName = name
			if h.oIRQueueManger.hasQueue(filename, name) {
				// 已经存在
			} else {
				// 不存在，则新建
				h.oIRQueueManger.oIR = append(h.oIRQueueManger.oIR, openIniResource{
					filename: filename,
					name:     name,
					ini:      ini,
				})
				h.oIRQueueManger.curName = name
				h.oIRQueueManger.curFname = filename
				h.oIRQueueManger.curIdx = len(h.oIRQueueManger.oIR) - 1
			}

			h.output(filename + "文件加载成功！")
			h.output("用时： " + rtMk())
		}
	} else {
		h.output("参数错误， 参考: $ open <filename>")
	}
}

// 使用 ini 资源
func (h *iniHelper) cmdUse() {
	xa := h.xa
	name := xa.Next(xa.Command)
	if name != "" {
		if name == h.oIRQueueManger.curName {
			h.output("您已经选择了 " + name)
		} else {
			if h.oIRQueueManger.hasAndChangeByName(name) {
				h.curName = name
			} else {
				h.output(name + " 不存在")
			}
		}
	} else {
		h.output("参数为空， $ use <name>")
	}
}

// 打印是资源列表
func (h *iniHelper) cmdList() {
	queue := []interface{}{}
	for k, _ := range h.resource {
		queue = append(queue, k)
	}
	h.output("资源列表如下：")
	fmt.Println(bin.FormatQue(queue))
}

// 显示帮助信息
func (h *iniHelper) cmdHelp() {
	h.output("ini文件加载测试器",
		"open <filename>   打开并加载 ini 文件",
		"use <name>        切换已经打开的资源",
		"list              列出全部的可用资源",
		"get <key>         获取键值",
		"exit              退出对话框",
	)
}

func (h *iniHelper) cmdGet() {
	xa := h.xa
	key := xa.Next(xa.Command)
	if key != "" {
		curName := h.curName
		if curName != "" {
			if rs, has := h.resource[curName]; has {
				if exist, v := rs.Get(key); exist {
					h.output(fmt.Sprintf("%v", v))
				} else {
					h.output("键值获取错误",
						"键值 "+key+" 不存在")
				}
			}
		} else {
			h.output("键值获取错误",
				"您还没有加载任何资源，请使用命令: open <filename> 加载资源")
		}
	} else {
		h.output("键值获取错误",
			"清楚参数有误: get <key> ")
	}
}

// 命令
func (h *iniHelper) cmdEmpty() {
	h.cmdHelp()
}

// 未知命令
func (h *iniHelper) cmdUnfind(cmd string) {
	h.output("命令错误",
		cmd+" 命令不存在！")
}
