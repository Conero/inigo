/* @ini-go V1.x
 * @Joshua Conero
 * @2017年11月1日 星期三
 * @ini 的 DataQueue 转化为 json
 */

package ini

import (
	"strings"
)

var jsonSettingTrans map[string]string = map[string]string{
	`\"`: "__JC__JSON_SG",
}

// json 字符串转移处理 isReconvert = true <-
func jsonTransform(sJson string, isReconvert bool) string {
	for k, v := range jsonSettingTrans {
		if isReconvert {
			sJson = strings.Replace(sJson, v, k, -1)
		} else {
			sJson = strings.Replace(sJson, k, v, -1)
		}
	}
	return sJson
}

// queue 转化为 字符串
func ToJsonStr(queue map[string]interface{}) string {
	jsonStr := ""
	jsonStrArr := []string{}
	for k, v := range queue {
		switch v.(type) {
		case string:
			value := jsonTransform(v.(string), false)
			value = jsonTransform(strings.Replace(value, `"`, `\"`, -1), false)
			value = jsonTransform(value, true)
			jsonStrArr = append(jsonStrArr, `"`+k+`": "`+value+`"`)
		case map[string]interface{}:
			cJsonStr := ToJsonStr(v.(map[string]interface{}))
			jsonStrArr = append(jsonStrArr, `"`+k+`": `+cJsonStr)
		case []string:
			cJsonStr := strings.Join(v.([]string), `","`)
			jsonStrArr = append(jsonStrArr, `"`+k+`": ["`+cJsonStr+`"]`)
		}
	}
	if len(jsonStrArr) > 0 {
		jsonStr = `{` + strings.Join(jsonStrArr, ",") + `}`
	}
	return jsonStr
}
