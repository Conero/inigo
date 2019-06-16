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
	oIR    map[string]openIniResource
	curKey string // 当前的唯一键值
}

// 初始化
func (oq *oIRQueueManger) init() {
	if oq.oIR == nil {
		oq.oIR = map[string]openIniResource{}
	}
}

// 获取用户列表
func (op *oIRQueueManger) getNameList() []interface{} {
	queue := []interface{}{}
	for k, _ := range op.oIR {
		queue = append(queue, k)
	}

	return queue
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
		h.curKey + ">>" + br

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
	bin.RegisterFunc("open", h.cmdOpen)   // 打开资源
	bin.RegisterFunc("new", h.cmdNew)     // 新增资源
	bin.RegisterFunc("use", h.cmdUse)     // 资源切换
	bin.RegisterFunc("list", h.cmdList)   // 获取资源礼列表
	bin.RegisterFunc("get", h.cmdGet)     // 获取资源中的值
	bin.RegisterFunc("help", h.cmdHelp)   // 获取的帮助信息
	bin.RegisterFunc("about", h.cmdAbout) // 查询资源信息
	bin.RegisterFunc("set", h.cmdSet)     // 资源值新增或设置
	bin.RegisterFunc("save", h.cmdSave)   // 保存新的资源/或者新文件

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
			// 获取键值
			_, vkey := path.Split(filename)
			if vAlias, exist := h.xa.Data["alias"]; exist {
				vkey = vAlias.(string)
			} else {
				nameQue := strings.Split(vkey, ".")
				vLen := len(nameQue)
				if vLen > 1 {
					vkey = strings.Join(nameQue[0:vLen-1], ".")
				}
			}

			if _, keyExist := h.oIR[vkey]; keyExist {
				// 已经存在
				// @TODO 存在时可以采用询问的方式继续或者类似 `--force` 强制性执行的参数
			} else {
				// 不存在，则新建
				oIr := openIniResource{
					filename: filename,
					ini:      ini,
					useTime:  rtMk(),
				}
				h.oIRQueueManger.oIR[vkey] = oIr
				h.oIRQueueManger.curKey = vkey
			}

			h.output(filename+"文件加载成功！",
				"用时： "+rtMk())
		}
	} else {
		h.output("参数错误， 参考: $ open <filename>")
	}
}

// 新增空的资源
func (h *iniHelper) cmdNew() {
	rtMk := util.SecCallStr()
	xa := h.xa
	name := xa.Next(xa.Command)
	if name != "" {
		if _, exist := h.oIR[name]; exist {
			h.output("[" + name + "] 资源已经存在，无法新增对应的资源")
		} else {
			newRs := openIniResource{
				ini:     inigo.NewParser(),
				useTime: rtMk(),
			}
			h.oIR[name] = newRs
			h.curKey = name
		}
	} else {
		h.output("新资源错误，格式有误：$ new <name>")
	}
}

// 使用 ini 资源
func (h *iniHelper) cmdUse() {
	xa := h.xa
	name := xa.Next(xa.Command)
	if name != "" {
		if name == h.oIRQueueManger.curKey {
			h.output("您已经选择了 " + name)
		} else {
			if _, exist := h.oIRQueueManger.oIR[name]; exist {
				h.oIRQueueManger.curKey = name
				h.output(name + " 切换成功！")
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
	queue := h.oIRQueueManger.getNameList()
	h.output("资源列表如下：")
	fmt.Println(bin.FormatQue(queue))
}

// 显示帮助信息
func (h *iniHelper) cmdHelp() {
	h.output("$ ini 交互式文件加载测试命令集合：",
		"open <filename>      打开并加载 ini 文件",
		"new <name>           创建空的资源",
		"use <name>           切换已经打开的资源",
		"list                 列出全部的可用资源",
		"get <key>            获取键值",
		"set <key> [<value>]  设置/更新当前的资源，value不设置是为空值",
		"save [<filename>]    保存当前的资源，空资源必须设置文件名；反正可能覆盖资源文件",
		"about [<name>]       打印当前的资源写信息",
		"exit                 退出对话框",
	)
}

// 键值获取
func (h *iniHelper) cmdGet() {
	xa := h.xa
	key := xa.Next(xa.Command)
	curKey := h.curKey
	if curKey == "" {
		h.output("当前还没有加载任何资源，请示先使用 open 打开资源")
		return
	}

	if key != "" {
		if rs, exist := h.oIRQueueManger.oIR[curKey]; exist {
			ini := rs.ini
			if exist, v := ini.Get(key); exist {
				h.output(fmt.Sprintf("%v", v))
			} else {
				h.output("键值获取错误",
					"键值 "+key+" 不存在")
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

// 资源值更新
func (h *iniHelper) cmdSet() {
	curKey := h.curKey
	if curKey != "" {
		xa := h.xa
		key := xa.Next(xa.Command)
		value := xa.Next(key)

		if key == "" {
			h.output("参数设置失败，格式有误: $ set <key> [<value>]")
		} else {
			rs := h.oIR[curKey]
			rs.ini.Set(key, value)
			h.output("设置值成功！")
		}
	} else {
		h.output("参数设置失败，当前还没有任何资源！")
	}
}

// 资源保存为新文件
func (h *iniHelper) cmdSave() {
	curKey := h.curKey
	if curKey != "" {
		xa := h.xa
		filename := xa.Next(xa.Command)
		curRs := h.oIR[curKey]

		if curRs.filename != "" {
			// 保存当前文件
			if filename == "" {
				curRs.ini.Save()
			} else {
				curRs.ini.SaveAsFile(filename)
			}
		} else if filename != "" {
			curRs.ini.SaveAsFile(filename)
		} else {
			h.output("保存失败，参数无效！")
		}
	} else {
		h.output("当前没有任何资源，资源保存失败")
	}
}

// 答应当前的资源
func (h *iniHelper) cmdAbout() {
	// 获取键值
	xa := h.xa
	key := xa.Next(xa.Command)
	if key == "" {
		key = h.oIRQueueManger.curKey
	}

	// 打印
	if key == "" {
		h.output("参数错误或者当前未加载任何资源，无法获取任何资源")
	} else {
		rs, exist := h.oIRQueueManger.oIR[key]
		if !exist {
			h.output("[" + key + "] 资源不存在")
		} else {
			h.output("["+key+"] 信息如下：",
				"加载用时    "+rs.useTime,
				"文件命令    "+rs.filename)
		}
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
