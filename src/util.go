/**
 * 2017年4月13日 星期四
 * 公共函数
 */
package ini

func ArrayFlip(arr []string) []string {
	vLen := len(arr)
	newArr := []string{}
	if vLen > 1 {
		for i := vLen; i >= 0; i-- {
			newArr = append(newArr, arr[i])
		}
	} else {
		newArr = arr
	}
	return newArr
}
