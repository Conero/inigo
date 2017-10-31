// 数据算法 - 2017年3月28日 星期二
// 数组相关
package common

// 是否在存在于数组中
func InArray(value string, arr []string) (int, bool) {
	vauleExist := false
	idx := -1
	for k, v := range arr {
		if v == value {
			vauleExist = true
			idx = k
			break
		}
	}
	return idx, vauleExist
}

// 是否在存在于数组中 - 支持类型： string/int
func In_Array(value, arr interface{}) (int, bool) {
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
