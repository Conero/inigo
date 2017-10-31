/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月31日 星期二
 * @ini 行处理
 */

package ini

import (
	"strings"
	"regexp"
)

// 行处理结构体
type Liner struct {
}

// 是否为注释行
func (line Liner) isComment(sLine string) bool {
	matched , _ := regexp.MatchString(IniParseSettings["reg_comment"], sLine)
	return matched
}

// 根据 "={" 检测数据， 返回值（是否为键值， 键值， 剩余line值）
// 获取基键
func (line Liner)getBaseKey(sLine string) (bool, string, string) {
	isBaseKey := false
	baseKey := ""
	idx := strings.Index(sLine, IniParseSettings["equal"])
	valueStr := ""
	if idx > -1{
		baseKey = strings.TrimSpace(sLine[0:idx])
		valueStr = strings.TrimSpace(sLine[idx+1:])
		reg := regexp.MustCompile("\\s")
		if "={" == reg.ReplaceAllString(valueStr, ""){
			isBaseKey = true
		}
	}
	return isBaseKey, baseKey, valueStr
}
// 字符串转移处理
func (line Liner) strTransform(sLine string) string {
	for k, v := range TranStrMap{
		sLine = strings.Replace(sLine, k, v, -1)
	}
	return sLine
}

// 转移字符恢复
func (line Liner) transRecover(sLine string) string{
	for k, v := range TranStrMap{
		sLine = strings.Replace(sLine, v, k, -1)
	}
	return sLine
}

// 行转变为键值对
func (line Liner) lineToKeyV(sLine string) (bool, string, interface{}){
	hasEqual := false
	key := ""
	var value interface{}
	idx := strings.Index(sLine, IniParseSettings["equal"])
	if idx != -1{
		hasEqual = true
		key = strings.TrimSpace(sLine[0:idx])
	}
	return hasEqual, key, value
}
// 单行多对象数组
func (line Liner)singleObject(cLine string) interface{} {
	var value interface{}
	return value
}

// key = value 中的value字符串转变为值
func (line Liner) strToData(sLine string) interface{} {
	transLine := strings.TrimSpace(line.strTransform(sLine))
	var value interface{}
	if matched, _ := regexp.MatchString(IniParseSettings["reg_scope_sg"], sLine); matched { // {}
	}else if strings.Index(transLine, IniParseSettings["limiter"]) > -1{	// 逗号数组
		newArr := []string{}
		for _,v := range strings.Split(transLine, IniParseSettings["limiter"]){
			newArr = append(newArr, line.transRecover(v))
		}
		value = newArr
	}else {		// 字符串
		value = line.transRecover(sLine)
	}
	return value
}