/**
	LnReader 文件行阅读器
	2018年7月10日 星期二
 */
package ini

import (
	"os"
	"bufio"
)

// 行阅读器
type LnReader struct {
	Filename string		// 文件名
	error
}

// 启动阅读器
func NewLnRer(filename string) *LnReader{
	return &LnReader{
		Filename: filename,
	}
}

// 扫描
func (ln *LnReader) Scan(callback func(line string)) bool {
	fs, err := os.Open(ln.Filename)
	if err == nil {
		buf := bufio.NewReader(fs)
		for {
			line, err2 := buf.ReadString('\n')
			callback(line)
			// 错误
			if err2 != nil{
				break
			}
		}
	} else {
		ln.error = err
		return false
	}
	return true
}