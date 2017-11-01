/* @ini-go V1.x
* @Joshua Conero
* @2017年10月28日 星期六
* @ini 文件解释器重写
 */

package rong

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ini 结构体
type Ini struct {
	FileName  string                 // ini 文件
	DataQueue map[string]interface{} // ini 解析后的数据
	IsSuccess bool                   // ini 文件是否解析成功
	FailMsg   string                 // 错误信息
	File      *File                  // 文件解析信息
}

// 入口文件
func Open(name string) *Ini {
	// 初始化对象
	inier := &Ini{
		FileName:  name,
		DataQueue: map[string]interface{}{},
		IsSuccess: false,
		FailMsg:   "",
		File: &File{
			line: 0,
		},
	}
	inier.reader()
	return inier
}

// 私有方法，文件读取
func (I *Ini) reader() {
	if len(I.FileName) > 0 {
		fs, err := os.Open(I.FileName)
		if err == nil {
			I.IsSuccess = true
			I.parseFile(fs)
		} else {
			I.FailMsg = err.Error()
		}
	}
}

//是否为注释行
func (I *Ini) isComment(line string) bool {
	matched, _ := regexp.MatchString(IniParseSettings["reg_comment"], line)
	return matched
}

// 章节处理
func (I *Ini) isSection(line string) (bool, string) {
	reg := regexp.MustCompile(IniParseSettings["reg_section"])
	matched := reg.MatchString(line)
	value := ""
	if matched {
		regSg := regexp.MustCompile(IniParseSettings["reg_section_sg"])
		value = regSg.ReplaceAllString(line, "")
	}
	return matched, value
}

// 解析文件
func (I *Ini) parseFile(fs *os.File) {
	// 节
	section := ""
	buf := bufio.NewReader(fs)
	for {
		line, err := buf.ReadString('\n')
		// 程序跳转前检测是否出错，出错直接中断循环，避免还没有检查错误时便继续进入循环(死循环)
		// -(2017年4月24日)新增的问题，最后一行还未完成时并提前结束
		isPanicError := false
		if err != nil {
			isPanicError = true
		}
		line = strings.TrimSpace(line)
		I.File.countLine()
		// 空行
		if !isPanicError && len(line) == 0 {
			continue
		}
		// 注释行
		if !isPanicError && I.isComment(line) {
			continue
		}
		// 章节判断
		isSct, sctKey := I.isSection(line)
		if isSct {
			section = sctKey
			I.DataQueue[section] = map[string]interface{}{}
			continue
		}
		// 值处理
		equlIdx := strings.Index(line, IniParseSettings["equal"])
		if equlIdx > -1 {
			key := strings.TrimSpace(line[0:equlIdx])
			value := strings.TrimSpace(line[equlIdx+1:])
			if section == "" {
				I.DataQueue[key] = value
			} else {
				sectionValue, has := I.DataQueue[section]
				sV := map[string]interface{}{}
				// 与历史值合并
				if has {
					sV = sectionValue.(map[string]interface{})
				}
				sV[key] = value
				I.DataQueue[section] = sV
			}
		}
		fmt.Println(line)
		if isPanicError {
			break
		}
	}
}

// 读取值
func (I *Ini) Get(key string) (bool, interface{}) {
	value, has := I.DataQueue[key]
	return has, value
}
func getStrByDQ(key string, dq map[string]interface{}) string {
	value := ""
	anyType, exist := dq[key]
	if exist {
		switch anyType.(type) {
		case string:
			value = anyType.(string)
		case int:
			value = strconv.Itoa(anyType.(int))
		}
	}
	return value
}

// 读取函数为字符串
// key 支持.二级操作
func (I *Ini) GetString(key string) string {
	value := ""
	pntIdx := strings.Index(key, ".")
	if pntIdx > -1 {
		fKey := key[0:pntIdx]
		mpDq, has := I.DataQueue[fKey]
		if has {
			switch mpDq.(type) {
			case map[string]interface{}:
				sKey := key[pntIdx+1:]
				value = getStrByDQ(sKey, mpDq.(map[string]interface{}))
			}
		}

	} else {
		value = getStrByDQ(key, I.DataQueue)
	}

	return value
}

// 是否存在值
func (I *Ini) HasKey(key string) bool {
	_, exist := I.DataQueue[key]
	return exist
}
