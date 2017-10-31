// unit
// @ 2017年4月6日 星期四
// @ 单元测试
package unit

import (
	"fmt"
)

//func PutWhenBool(isTest bool,...interface{}){}
// 条件输出测试
func PutWhenBool(isTest bool, args ...interface{}) {
	if isTest && len(args) > 0 {
		fmt.Println(args)
	}
}
