// ini
package ini

import (
	"fmt"
	"strings"
	"time"
)

// 配置文件对象
type Conf struct {
	filename       string                 // 文件名称
	content        string                 // 内容
	mtime          time.Time              // 用于计算时间差-效率评估量
	endtime        time.Time              // 完成时间
	seekCount      int                    // 遍历次数
	lineCount      int                    // 行数
	Cdt            map[string]interface{} // 编辑成功后的值
	FullComplier   bool                   // 是否全编译，默认为 true
	CloseRender    bool                   // 关闭编译功能
	RenderVariable map[string]string      // 手动设置的系统可渲染值
}

// 不定函数用于避免非必要输出
func out(args ...interface{}) {
	fmt.Println("")
	for _, arg := range args {
		fmt.Print(arg, ",")
	}
}

// 获取设置值
// key -> conf.equal => =
func GetVal(key string) string {
	idx := strings.Index(key, ".")
	k1 := key[0:idx]
	vmap, has := Setting[k1]
	if !has {
		return ""
	}
	k2 := key[idx+1:]
	str, has2 := vmap[k2]
	if !has2 {
		return ""
	}
	return str
}

// 获取设置值为map
func GetVmap(key string) (setType, bool) {
	vmap, has := Setting[key]
	if !has {
		return setType{}, false
	}
	return vmap, has
}

// 是否在存在于数组中 - 支持类型： string/int
func InArray(value, arr interface{}) (int, bool) {
	vauleExist := false
	idx := -1
	switch value.(type) {
	case string:
		for k, v := range arr.([]string) {
			if v == value.(string) {
				vauleExist = true
				idx = k
				break
			}
		}
	case int:
		for k, v := range arr.([]int) {
			if v == value.(int) {
				vauleExist = true
				idx = k
				break
			}
		}
	}

	return idx, vauleExist
}

// 遍历 json 字符串 - 递归法
func ToJson(data map[string]interface{}) string {
	tmpArray := []string{}
	for k, v := range data {
		value := `""`
		switch v.(type) {
		case string:
			value = `"` + v.(string) + `"`
		case []string:
			value = `["` + strings.Join(v.([]string), `","`) + `"]`
		case map[string]interface{}:
			value = ToJson(v.(map[string]interface{}))
		case vClass:
			value = ToJson(v.(vClass))
		case sArr:
			value = `["` + strings.Join(v.(sArr), `","`) + `"]`
		}
		value = `"` + k + `":` + value
		tmpArray = append(tmpArray, value)
	}
	return "{\n" + strings.Join(tmpArray, ",\n") + "\n}"
}
