// ini
package ini

import (
	"fmt"
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
