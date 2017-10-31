/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月31日 星期二
 * @ini 内部测试器
 */
package ini

import "fmt"

type Tester struct {
}

func (t *Tester) LinerTest()  {
	liner := Liner{}
	// 当行多列测试
	fmt.Println(liner.singleObject("{ key = { key2 ={ key3 = { key4 = { key5 = { key6 = value6 }}}}}}"))
} 