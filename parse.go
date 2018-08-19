/* @ini-go V1.x
* @Joshua Conero
* @2017年10月28日 星期六
* @ini 文件解释器重写
 */

package inigo

import (
	"strconv"
	"strings"
	"regexp"
)

// ini 结构体
type Ini struct {
	Liner                            // 组合继承- 行处理
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
		ln := NewLnRer(I.FileName)
		// 行扫描
		bba := BBAnalyze(I)
		I.IsSuccess = ln.Scan(func(line string) {
			line = strings.TrimSpace(line)
			I.File.countLine()
			// 非错误
			//if !isPanicError {
			// 空行
			if len(line) == 0 {
				return
			}
			// 多行注释
			if I.isMultiComment(line) {
				return
			}
			// 单行注释
			if I.isComment(line) {
				return
			}
			// 键值结束
			if line == IniParseSettings["scope2"] {
				bba.CommitQueue()
				return
			}
			line = I.transComment(line, false)
			regCmt := regexp.MustCompile(IniParseSettings["reg_has_comment"])
			CmtIdx := regCmt.FindStringIndex(line)
			if CmtIdx != nil{
				line = strings.TrimSpace(line[0:CmtIdx[0]])
			}
			line = I.transComment(line, true)
			// 多行字符串/字符串数组
			isMl, isEnd, mKey, mValue := I.mLineString(line)
			if isMl {
				return
			}
			// 多行数组结束
			if isEnd {
				if len(mKey) > 0{
					bba.PushQueue(mKey, mValue)
				}else{
					bba.MultiLineToArray(mValue)
				}
				return
			}

			// 获取基键
			isBK, BK, nLine := I.getBaseKey(line)
			//fmt.Println(isBK, BK, nLine, line)
			if isBK { // 是基键
				bba.UpdateBaseKey(BK)
				return
			} else if BK != "" {
				//bba.PushQueue(BK, nLine)
				bba.PushQueue(BK, I.strToData(nLine))
			} else {
				bba.MultiLineToArray(nLine)
			}
		})
		if !I.IsSuccess{
			I.FailMsg = ln.Error()
		}
		I.DataQueue = bba.DataQueue
	}
}

// 读取值
// 支持点级多级数据查询
func (I *Ini) Get(key string) (bool, interface{}) {
	//value, has := I.DataQueue[key]
	has := false
	var value interface{}
	// 默认数据类型
	value = I.DataQueue
	for _,v := range strings.Split(key, "."){
		switch value.(type) {
		case map[string]interface{}:
			tValue, tHas := value.(map[string]interface{})[v]
			has = tHas
			if has{
				value = tValue
			}
		}
	}
	return has, value
}

// 读取函数为字符串
// 支持点操作，多级数据获取
func (I *Ini) GetString(key string) string {
	value := ""
	exist, anyType := I.Get(key)
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

// 是否存在值
func (I *Ini) HasKey(key string) bool {
	_, exist := I.DataQueue[key]
	return exist
}

// 转化为json字符串
func (I *Ini) ToJsonString() string {
	return MkCreator(I.DataQueue).ToJsonString()
}
