/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月31日 星期二
 * @ini 行处理
 */

package ini

import (
	"regexp"
	"strings"
)

// 行处理结构体
type Liner struct {
	MultiCommentIngMk bool // 多行注释标注
}

// 是否为注释行
func (line Liner) isComment(sLine string) bool {
	matched, _ := regexp.MatchString(IniParseSettings["reg_comment"], sLine)
	return matched
}
func (line *Liner) isMultiComment(sLine string) bool {
	shouldToNextLine := false
	// 多行注释介绍
	if line.MultiCommentIngMk && sLine == IniParseSettings["mcomment2"] {
		line.MultiCommentIngMk = false
		shouldToNextLine = true
	} else if line.MultiCommentIngMk { // 处于多行注释行中
		shouldToNextLine = true
	} else if !line.MultiCommentIngMk && sLine == IniParseSettings["mcomment1"] { // 多行注释开始
		line.MultiCommentIngMk = true
		shouldToNextLine = true
	}
	return shouldToNextLine
}

// 根据 "={" 检测数据， 返回值（是否为键值， 键值， 剩余line值）
// 获取基键
func (line Liner) getBaseKey(sLine string) (bool, string, string) {
	isBaseKey := false
	baseKey := ""
	idx := strings.Index(sLine, IniParseSettings["equal"])
	valueStr := ""
	// 有等于符号
	if idx > -1 {
		baseKey = strings.TrimSpace(sLine[0:idx])
		valueStr = strings.TrimSpace(sLine[idx+1:])
		//reg := regexp.MustCompile("\\s")
		reg := regexp.MustCompile(`\s`)
		eSg := strings.TrimSpace(reg.ReplaceAllString(valueStr, ""))
		if "{" == eSg {
			//if "={" == eSg {
			isBaseKey = true
		}
		//fmt.Println(eSg, "-<")
	} else {
		valueStr = sLine
	}
	return isBaseKey, baseKey, valueStr
}

// 字符串转移处理
func (line Liner) strTransform(sLine string) string {
	for k, v := range TranStrMap {
		sLine = strings.Replace(sLine, k, v, -1)
	}
	return sLine
}

// 转移字符恢复
func (line Liner) transRecover(sLine string) string {
	for k, v := range TranStrMap {
		sLine = strings.Replace(sLine, v, k, -1)
	}
	sLine = strings.Replace(sLine, "\\", "", -1)
	return sLine
}

// 行转变为键值对
func (line Liner) lineToKeyV(sLine string) (bool, string, interface{}) {
	hasEqual := false
	key := ""
	var value interface{}
	idx := strings.Index(sLine, IniParseSettings["equal"])
	if idx != -1 {
		hasEqual = true
		key = strings.TrimSpace(sLine[0:idx])
	}
	return hasEqual, key, value
}

// { key = { key2 ={ key3 = { key4 = { key5 = { key6 = value6 }}}}}}			=> map
// {v1, v2, v3}	=> []string
// map[key:map[key2:map[key3:map[key4:map[key5:map[key6:value6 是一个复杂的字符串{key= value}]]]]]]
// 单行多对象数组
func (line Liner) singleObject(cLine string) interface{} {
	cLine = line.strTransform(cLine)
	var value interface{}
	if strings.Index(cLine, IniParseSettings["equal"]) == -1 { // {}
		strLen := len(cLine)
		cLine = cLine[1 : strLen-2]
		baseStrArr := []string{}
		for _, v := range strings.Split(cLine, IniParseSettings["limiter"]) {
			baseStrArr = append(baseStrArr, line.transRecover(v))
		}
		value = baseStrArr
	} else {
		var baseValueMd map[string]interface{}
		var baseValue map[string]string
		reg := regexp.MustCompile(IniParseSettings["reg_scope"])
		// 安全计数器
		safeCounter := 20
		i := 0
		for {
			i = i + 1
			if len(cLine) == 0 {
				break
			}
			mStr := reg.FindString(cLine)
			cLine = strings.Replace(cLine, mStr, "", -1)
			mStrLen := len(mStr)
			if mStrLen > 0 && mStr[0] == '{' && mStr[mStrLen-1] == '}' {
				mStr = mStr[1 : mStrLen-1]
			}
			k := ""
			equalIdx := strings.Index(mStr, IniParseSettings["equal"])
			if equalIdx > -1 {
				k = strings.TrimSpace(mStr[0 : equalIdx-1])
				mStr = strings.TrimSpace(mStr[equalIdx+1:])
			}
			if "" != k && mStr != "" {
				baseValue = map[string]string{k: line.transRecover(mStr)}
			} else if "" != k && mStr == "" {
				if nil != baseValueMd {
					baseValueMd = map[string]interface{}{k: baseValueMd}
				} else {
					baseValueMd = map[string]interface{}{k: baseValue}
				}
			}
			if i >= safeCounter {
				break
			}
		}
		value = baseValueMd
	}
	return value
}

// key = value 中的value字符串转变为值
func (line Liner) strToData(sLine string) interface{} {
	transLine := strings.TrimSpace(line.strTransform(sLine))
	var value interface{}
	//print("\r\n", sLine, " ^^ ")
	//fmt.Println(regexp.MatchString(IniParseSettings["reg_scope_sg"], sLine))
	if matched, _ := regexp.MatchString(IniParseSettings["reg_scope_sg"], sLine); matched { // {}
		value = line.singleObject(sLine)
	} else if strings.Index(transLine, IniParseSettings["limiter"]) > -1 { // 逗号数组
		newArr := []string{}
		for _, v := range strings.Split(transLine, IniParseSettings["limiter"]) {
			newArr = append(newArr, line.transRecover(v))
		}
		value = newArr
	} else { // 字符串
		value = line.transRecover(sLine)
	}
	return value
}
