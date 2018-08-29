package inigo

// @Date：   2018/8/19 0019 14:25
// @Author:  Joshua Conero
// @Name:    基本 go 解析器

type BaseParser struct {
	Data  map[string]interface{}
	valid bool
}

// data 数据检测
func (p *BaseParser) GetData() map[string]interface{} {
	if p.Data == nil {
		p.Data = map[string]interface{}{}
	}
	return p.Data
}

func (p *BaseParser) Get(key string) (bool, interface{}) {
	data := p.GetData()
	value, has := data[key]
	return has, value
}

func (p *BaseParser) HasKey(key string) bool {
	data := p.GetData()
	_, has := data[key]
	return has
}

func (p *BaseParser) Raw(key string) string {
	raw := ""
	return raw
}

func (p *BaseParser) Value(params ...interface{}) interface{} {
	// key, nil, def
	if len(params) > 2 {
		if has, value := p.Get(params[0].(string)); has {
			return value
		}
		return params[2]
	} else if len(params) > 1 {
		p.Set(params[0].(string), params[1])
	} else if len(params) == 1 {
		if has, value := p.Get(params[0].(string)); has {
			return value
		}
	}
	return nil
}

func (p *BaseParser) GetAllSection() []string {
	return nil
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

func (p *BaseStrParse) Line() int{
	return p.line
}

func (p *BaseStrParse) GetData() map[interface{}]interface{}{
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