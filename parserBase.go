package inigo

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// @Date：   2018/8/19 0019 14:25
// @Author:  Joshua Conero
// @Name:    基本 go 解析器
const (
	baseCommentReg = "^#|;"               // 注释符号
	baseSectionReg = "^\\[[^\\[\\]]+\\]$" // 节正则

	baseEqualToken = "="      // 等于符号
	baseSecRegPref = "__sec_" // 节前缀
)

type BaseParser struct {
	valid   bool
	section []string
	container
	filename string // 文件名
}

func (p *BaseParser) Raw(key string) string {
	raw := ""
	return raw
}

func (p *BaseParser) GetAllSection() []string {
	return p.section
}

func (p *BaseParser) Section(params ...interface{}) interface{} {
	var value interface{}
	var section, key string

	if nil == params {
		value = nil
	} else if len(params) == 1 {
		format := params[0].(string)
		if idx := strings.Index(format, "."); idx > -1 {
			section = format[0:idx]
			key = format[idx+1:]
		}
	} else if len(params) > -1 {
		section = params[0].(string)
		key = params[1].(string)
	}

	if section != "" && key != "" {
		if data, hasSection := p.Data[baseSecRegPref+section]; hasSection {
			dd := data.(map[interface{}]interface{})
			if v, hasKey := dd[key]; hasKey {
				value = v
			}
		}
	}
	return value
}

func (p *BaseParser) Set(key string, value interface{}) Parser {
	p.GetData()
	p.Data[key] = value
	return p
}

func (p *BaseParser) IsValid() bool {
	return p.valid
}

func (p *BaseParser) OpenFile(filename string) Parser {
	reader := &baseFileParse{}
	reader.read(filename)
	p.Data = reader.GetData()
	p.filename = filename
	return p
}

func (p *BaseParser) ReadStr(content string) Parser {
	return p
}

func (p *BaseParser) Save() bool {
	filename := p.filename
	return p.SaveAsFile(filename)
}

func (p *BaseParser) SaveAsFile(filename string) bool {
	successMk := true
	// 简单处理=字符串类型
	// @todo 需要做更多(20181105)
	iniTxt := "; power by (" + Name + "; V" + Version + "/" + Release + ")" +
		"\n;time: " + time.Now().String() +
		"\n; github.com/" + Name
	for k, v := range p.Data {
		switch k.(type) {
		case string:
			iniTxt += "\n" + k.(string) + "	= "
			if _, isStr := v.(string); isStr {
				iniTxt += v.(string)
			}
		}
	}
	// 0644 Append
	// 0755
	err := ioutil.WriteFile(filename, []byte(iniTxt), 0755)
	if err != nil {
		fmt.Println(err.Error())
		successMk = false
	}
	return successMk
}

// 获取驱动名称
func (p BaseParser) Driver() string  {
	return SupportNameIni
}
// =>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>(BaseStrParse)>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// baseStrParse
//
type BaseStrParse struct {
	data map[interface{}]interface{}
	line int
}

func (p *BaseStrParse) Line() int {
	return p.line
}

func (p *BaseStrParse) GetData() map[interface{}]interface{} {
	return p.data
}

func (p *BaseStrParse) LoadContent(content string) StrParser {
	p.data = map[interface{}]interface{}{}
	lineCtt := 0
	str2lines(content, func(line string) {
		lineCtt += 1
	})
	p.line = lineCtt
	return p
}
