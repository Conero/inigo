package inigo

import "strings"

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
}

func (p *BaseParser) Raw(key string) string {
	raw := ""
	return raw
}

func (p *BaseParser) GetAllSection() []string {
	return p.section
}

func (p *BaseParser) Section(params ...interface{}) interface{}{
	var value interface{}
	var section, key string

	if nil == params{
		value = nil
	}else if len(params) == 1{
		format := params[0].(string)
		if idx := strings.Index(format, "."); idx > -1{
			section = format[0: idx]
			key = format[idx+1:]
		}
	}else if len(params) > -1{
		section = params[0].(string)
		key = params[1].(string)
	}

	if section != "" && key != ""{
		if data, hasSection := p.Data[baseSecRegPref+section]; hasSection{
			dd := data.(map[interface{}]interface{})
			if v, hasKey := dd[key]; hasKey{
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
	return p
}

func (p *BaseParser) ReadStr(content string) Parser {
	return p
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
