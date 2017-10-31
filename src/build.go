/* @Joshua Conero
 * @2017年3月18日 星期六
 * @ini 文件解生成器
 */

package ini

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Builder struct {
	Lenght   int               // 数据长度
	Pathname string            // 文件名称
	Bdt      map[string]string // 生成器数据字典
}

// 用字典数据初始化
func MapBuilder(data map[string]string) *Builder {
	_builder := &Builder{
		Bdt: map[string]string{},
	}
	_builder.Lenght = len(_builder.Bdt)
	for k, v := range data {
		_builder.Lenght = _builder.Lenght + 1
		_builder.Bdt[k] = v
	}
	return _builder
}

// 初始化空数据栈
func EmptyBuilder() *Builder {
	return &Builder{
		Bdt: map[string]string{},
	}
}

// 设置值
func (bder *Builder) SetValue(key, value string) int {
	_, has := bder.Bdt[key]
	if !has {
		bder.Lenght = bder.Lenght + 1
	}
	bder.Bdt[key] = value
	return bder.Lenght
}

// 数字数组
func (bder *Builder) SetArray(key string, arr []string) {
	_, has := bder.Bdt[key]
	if !has {
		bder.Lenght = bder.Lenght + 1
	}
	bder.Bdt[key] = strings.Join(arr, ",")
}

// 设置生成文件路劲
func (dber *Builder) SetPath(path string) {
	dber.Pathname = path
}

// 生成文字
func (dber *Builder) ToString() string {
	value := ""
	lenght := dber.Lenght
	if lenght > 0 {
		var tmpArray []string
		for k, v := range dber.Bdt {
			str := k + " = " + v
			tmpArray = append(tmpArray, str)
		}
		value = strings.Join(tmpArray, "\r\n")
	}
	return value
}

// 保存文件
func (dber *Builder) Save() bool {
	pathname := dber.Pathname
	if len(pathname) == 0 {
		dber.SetPath("./" + AUTHOR + "_" + strconv.Itoa(int(time.Now().Unix())) + ".ini")
		pathname = dber.Pathname
	}
	content := dber.ToString()
	if len(content) == 0 {
		return false
	}
	error := ioutil.WriteFile(pathname, []byte(content), 0x644)
	if error != nil {
		return false
	}
	return true
}

// 删除值
func (dber *Builder) Del(key string) bool {
	_, has := dber.Bdt[key]
	if has {
		delete(dber.Bdt, key)
		return true
	}
	return false
}
