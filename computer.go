// computer
package main

import (
	"fmt"
)
import (
	"os"
	"os/user"
	"runtime"
	"strings"
)
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

/*
	计算机处理
*/

//计算机基本信息
func (mw *MyMainWindow) cpinfo_Triggered() {
	var ctt string = ""
	name, _ := os.Hostname()
	ctt += "\r\n计算机名称：" + name
	ctt += "\r\n系统版本号：" + runtime.Version()
	ctt += "\r\n操作系统：" + runtime.GOOS
	ctt += "\r\n操作系统：" + runtime.GOARCH
	u, _ := user.Current()
	ctt += "\r\n用户ID" + u.Uid // user id
	ctt += "\r\n分组ID" + u.Gid //primary group id
	ctt += "\r\n用户名称" + u.Username
	ctt += "\r\n名称" + u.Name
	ctt += "\r\n主地址" + u.HomeDir
	mw.te.SetText(ctt)
}

//Foo 结构体换位大学开头是正常 ? value/key 抛出异常
type Foo struct {
	Index int
	Key   string
	Value string
}
type FooModel struct {
	//walk.TableModelBase
	//walk.SorterBase
	//sortColumn int
	//sortOrder  walk.SortOrder
	//evenBitmap *walk.Bitmap
	//oddIcon    *walk.Icon
	//walk.SortedReflectTableModelBase
	items []*Foo
}

//无此方法时抛出错误：dataSource must be assignable to []map[string]interface{}
func (m *FooModel) Items() interface{} {
	return m.items
}

//环境变量
func (mw *MyMainWindow) cpEnvir_Triggered() {
	/*
		var ctt string = ""
		cttarr := os.Environ()
		for i := 0; i < len(cttarr); i++ {
			ctt += "\r\n" + cttarr[i]
		}
		mw.te.SetText(ctt)
	*/

	//填充数据
	NewFooModel := func() *FooModel {
		//m := new(FooModel)
		//m.items = make([]*Foo, 1000)
		cttarr := os.Environ()
		//fmt.Println(len(cttarr)) //长度
		m := &FooModel{items: make([]*Foo, len(cttarr))}
		for i := range m.items {
			tmp := strings.Split(cttarr[i], "=")
			m.items[i] = &Foo{
				Index: i + 1,
				Key:   tmp[0],
				Value: tmp[1],
			}
		}
		/*
			for i := 0; i < len(cttarr); i++ {
				tmp := strings.Split(cttarr[i], "=")
				m.items[i] = &Foo{
					Index: i,
					key:   tmp[0],
					value: tmp[1],
				}
			}
		*/
		return m
	}
	//mw.tv.
	//var dlg *walk.Dialog
	n, err := Dialog{
		AssignTo: &mw.dlg,
		Title:    "系统环境变量",
		Layout:   VBox{}, //设置它时才能显示图层内容
		MinSize:  Size{800, 500},
		Children: []Widget{
			TableView{
				AssignTo:              &mw.tv,
				AlternatingRowBGColor: walk.RGB(255, 255, 224),
				ColumnsOrderable:      true,
				MultiSelection:        true,
				Columns: []TableViewColumn{
					{Name: "Index", Title: "序号", Width: 50},
					{Name: "Key", Title: "环境变量名", Width: 200},
					{Name: "Value", Title: "环境变量值", Width: 350},
				},
				Model: NewFooModel(),
			},
		},
	}.Run(mw)

	if err != nil {
		fmt.Println(err.Error())
		//mw.te.SetText(err.Error())
	} else {
		fmt.Println(n, "《返回值")
	}
}
