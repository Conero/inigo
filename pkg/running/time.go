/**
@date 2017年11月3日 星期五
@author Joshua Conero
@name 运行时工具
*/
package running

import (
	"time"
	"fmt"
	"strconv"
)

// 时间运行器
type RTime struct {
	StartMtime time.Time // 起始时间戳
}

// 创建运行计时器
func CreateTimer() *RTime {
	rt := &RTime{
		StartMtime: time.Now(),
	}
	return rt
}

// 获取到运行纳秒
func (Rt *RTime) GetNanoSec() int64 {
	return time.Now().UnixNano() - Rt.StartMtime.UnixNano()
}

// 获取运行统计秒
func (Rt *RTime) GetSec() float64 {
	num, _ := strconv.ParseFloat(Rt.GetSecString(4), 64)
	//fmt.Println(strconv.ParseFloat(strNum, 64))
	//println(strNum)
	//println(DetSec, Rt.GetNanoSec(), float64(Rt.GetNanoSec())/float64(1E6))
	return num
}

// 获取运行秒为字符串
func (Rt *RTime) GetSecString(bit int)  string {
	DetSec := float64(Rt.GetNanoSec()) / float64(1E9)
	return fmt.Sprintf("%0."+strconv.Itoa(bit)+"f", DetSec)
}