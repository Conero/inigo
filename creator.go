/* @ini-go V1.x
 * @Joshua Conero
 * @2017年11月1日 星期三
 * @ini 的 DataQueue 转化为 json, 生成器
 */

package inigo

import (
	"strconv"
	"strings"
	"time"
)

// 结构体 生成器
type Creator struct {
	DataQueue map[string]interface{}
}

// 生成器配置参数
var creatorSetting = map[string]string{
	"delimiter": "	", // 分割长度, 1 tab
}

// 生成器常量
const (
	BR = "\r\n"
)

// 生成器初始化
func MkCreator(queue map[string]interface{}) *Creator {
	crt := &Creator{
		DataQueue: queue,
	}
	return crt
}

// ini 文件声明前缀
func (crt *Creator) iniFilePreText() string {
	return `
;	文档生成日期 = ` + time.Now().Format("2006-01-02 15:04:05") + `
;	pkg/ini 包版本 = v` + VERSION + `(` + BUILD + `)
`
}

// ini 格式化私有方法
// queue 数据列队
// class 级别，从0开始
func (crt *Creator) iniFormat(queue map[string]interface{}, class int) string {
	iniString := ""
	prefStr := ""
	if class > 0 {
		prefStr = strings.Repeat(creatorSetting["delimiter"], class)
	}
	for k, v := range queue {
		switch v.(type) {
		case string:
			iniString = iniString + prefStr + k + " = " + v.(string) + BR
		case int:
			iniString = iniString + prefStr + k + " = " + strconv.Itoa(v.(int)) + BR
		case []string:
			iniString = iniString +
				prefStr + k + " = {" + BR +
				strings.Repeat(creatorSetting["delimiter"], class+1) +
				strings.Join(v.([]string), BR+strings.Repeat(creatorSetting["delimiter"], class+1)) + BR +
				prefStr + "}" + BR
		case map[string]string:
			cPrefStr := strings.Repeat(creatorSetting["delimiter"], class+1)
			iniString = iniString + prefStr + k + " = {" + BR
			for kk, vv := range v.(map[string]string) {
				iniString = iniString + cPrefStr + kk + " = " + vv + BR
			}
			iniString = iniString + prefStr + "}" + BR
		case map[string]interface{}:
			iniString = iniString +
				prefStr + k + " = {" + BR +
				crt.iniFormat(v.(map[string]interface{}), class+1) +
				prefStr + "}" + BR
		}
	}
	if class == 0 {
		iniString = crt.iniFilePreText() + BR + iniString
	}
	return iniString
}

// 生成格式化ini字符串(标准化)
func (crt *Creator) ToIniString() string {
	return crt.iniFormat(crt.DataQueue, 0)
}

// ini 格式化私有方法
// queue 数据列队
// class 级别，从0开始
func (crt *Creator) jsonFormat(queue map[string]interface{}, class int) string {
	jsonString := ""
	delimiter := creatorSetting["delimiter"]
	jsonStrArr := []string{}
	for k, v := range queue {
		switch v.(type) {
		case string:
			jsonStrArr = append(jsonStrArr, strings.Repeat(delimiter, class+1)+`"`+k+`":"`+v.(string)+`"`)
		case int:
			jsonStrArr = append(jsonStrArr, strings.Repeat(delimiter, class+1)+`"`+k+`":`+strconv.Itoa(v.(int)))
		case []string:
			arrayStr := strings.Repeat(delimiter, class+1) + `"` + k + `":[` + BR +
				strings.Repeat(delimiter, class+2) + `"` + strings.Join(v.([]string), `",`+BR+strings.Repeat(delimiter, class+2)+`"`) + `"` + BR +
				strings.Repeat(delimiter, class+1) + "]"
			jsonStrArr = append(jsonStrArr, arrayStr)
		case map[string]string:
			cJsonStr := ""
			cArr := []string{}
			for kk, vv := range v.(map[string]string) {
				//cArr = append(cArr, strings.Repeat(delimiter, class+2)+`"`+kk+`":"`+vv+`"`)
				cArr = append(cArr, strings.Repeat(delimiter, class+1)+`"`+kk+`":"`+vv+`"`)
			}
			if len(cArr) > 0 {
				cJsonStr = strings.Repeat(delimiter, class+1) + `"` + k + `":{` + BR +
					strings.Repeat(delimiter, class+1) + strings.Join(cArr, ","+BR+strings.Repeat(delimiter, class+1)) + BR +
					strings.Repeat(delimiter, class+1) + "}"
			}else{
				cJsonStr += strings.Repeat(delimiter, class+1) + `"` + k + `":`
			}
			jsonStrArr = append(jsonStrArr, cJsonStr)
		case map[string]interface{}:
			cMapStr := strings.Repeat(delimiter, class+1) + `"` + k + `":{` + BR +
				crt.jsonFormat(v.(map[string]interface{}), class+1)+ BR +
				strings.Repeat(delimiter, class+1) + "}"
			jsonStrArr = append(jsonStrArr, cMapStr)
		}
	}
	if len(jsonStrArr) > 0{
		jsonString = strings.Join(jsonStrArr, ","+BR)
	}
	if class == 0{
		jsonString = "{" + BR + jsonString + BR + "}"
	}
	return jsonString
}

// 生成格式化json字符串(标准化)
func (crt *Creator) ToJsonString() string {
	return crt.jsonFormat(crt.DataQueue, 0)
}
