// brximl
package main

import (
	"log"
	"strconv"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

/*
import (
	"fmt"
)
*/
var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
	te  *walk.TextEdit
	tv  *walk.TableView
	tb  *walk.TabWidget
	dlg *walk.Dialog
}

/*
	互联网访问地址>> IP:PORT
*/
type Addr struct {
	ip   string
	port string
}

//服务器
var Host Addr

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)
	mw := new(MyMainWindow)
	//图标设置无效
	icon, err := walk.NewIconFromFile("./img/gui.ico")
	if err != nil {
		log.Fatal(err)
	}
	ic, err := walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}
	defer ic.Dispose()
	if err := ic.SetIcon(icon); err != nil {
		log.Fatal(err)
	}

	var openAction, showAboutBoxAction, dialTest, webServer, webApiServer, setting *walk.Action
	var recentMenu, locationMenu, computerMenu, dbMenu *walk.Menu
	//var tv *walk.TableView
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "GO图形化",
		MenuItems: []MenuItem{
			Menu{
				Text: "&菜单(M)",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "./img/open.png", //	../ 表示当前文件的上一级文件,./当前文件夹
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: mw.openAction_Triggered,
					},
					Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					Menu{
						AssignTo: &locationMenu,
						Text:     "本地",
					},
					Menu{
						AssignTo: &computerMenu,
						Text:     "计算机",
					},
					Menu{
						AssignTo: &dbMenu,
						Text:     "数据库",
					},
					Separator{},
					Action{
						Text:        "&退出",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&网络(I)",
				Items: []MenuItem{
					Action{
						AssignTo:    &dialTest,
						Text:        "&拨号测试",
						OnTriggered: mw.dialtest_Trig,
					},
				},
			},
			Menu{
				Text: "&服务器(S)",
				Items: []MenuItem{
					Action{
						AssignTo: &webServer,
						Text:     "&运行",
					},
					Action{
						AssignTo:    &webApiServer,
						Text:        "&网络接口调试",
						OnTriggered: mw.ApiServer,
					},
				},
			},
			Menu{
				Text: "&设置(T)",
				Items: []MenuItem{
					Action{
						AssignTo: &setting,
						Text:     "&设置",
					},
				},
			},
			Menu{
				Text: "&帮助(H)",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "关于我",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				ActionRef{&openAction},
				Menu{
					Text:  "New A",
					Image: "./img/document-new.png",
					Items: []MenuItem{
						Action{
							Text:        "A",
							OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text:        "B",
							OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text:        "C",
							OnTriggered: mw.newAction_Triggered,
						},
					},
					OnTriggered: mw.newAction_Triggered,
				},
				Separator{},
				Menu{
					Text:  "View",
					Image: "./img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text:        "X",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Y",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Z",
							OnTriggered: mw.changeViewAction_Triggered,
						},
					},
				},
				Separator{},
				Action{
					Text:        "Special",
					Image:       "./img/system-shutdown.png",
					OnTriggered: mw.specialAction_Triggered,
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},
		MinSize: Size{800, 500},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &mw.te,
			},
			TableView{
				AssignTo: &mw.tv,
			},
			//TabWidget{
			//	AssignTo: &mw.tb,
			//},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openAction_Triggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	//本地子菜单

	addlocationActions := func(text ...string) {
		for _, text := range text {
			a := walk.NewAction()
			a.SetText(text)
			switch text {
			case "IP":
				a.Triggered().Attach(mw.locip_Triggered)
			case "地理信息":
				a.Triggered().Attach(mw.locpos_Triggered)

			}
			//a.Triggered().Attach(mw.openAction_Triggered)
			locationMenu.Actions().Add(a)
		}
	}
	addlocationActions("IP", "地理信息")
	//计算机
	addcomputerActions := func(text ...string) {
		for _, text := range text {
			a := walk.NewAction()
			a.SetText(text)
			switch text {
			case "基本信息":
				a.Triggered().Attach(mw.cpinfo_Triggered)
			case "环境变量":
				a.Triggered().Attach(mw.cpEnvir_Triggered)
			}
			computerMenu.Actions().Add(a)
		}
	}
	addcomputerActions("基本信息", "环境变量")
	//数据库
	addDbMenuAction := func(text ...string) {
		for _, text := range text {
			a := walk.NewAction()
			a.SetText(text)
			switch text {
			case "Mysql":
				a.Triggered().Attach(mw.mysql_Triggered)
			case "Oracle":
				a.Triggered().Attach(mw.cpEnvir_Triggered)
			}
			dbMenu.Actions().Add(a)
		}
	}
	addDbMenuAction("Mysql", "Oracle")

	mw.Run()
}

//IP
func (mw *MyMainWindow) locip_Triggered() {
	ipstr := ip_info()
	mw.te.SetText(ipstr)
	//walk.MsgBox(mw, "IP", ipstr, walk.MsgBoxIconInformation)
}

//地理信息
func (mw *MyMainWindow) locpos_Triggered() {
	start := sec(0)
	ctt := pos_info() //内容
	ctt += "\r\n用时:" + strconv.FormatFloat(sec(start), 'f', -1, 64) + "s"
	mw.te.SetText(ctt)
	//walk.MsgBox(mw, "地理信息", ctt, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) openAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

//关于我
func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	// \n、\r\n、<br> 换行无效
	str := "作者: 杨华"
	str += "\n英文：Joshua Conero Doeeking"
	str += "\n项目说明：主要目的是实现通过对go以及图形化库的学习以及使用，加深桌面应用的编程方案"
	str += "\n其他：采用go语言，以及第三方Windows GUI库walk"
	str += "\nGo中文网址：https://go-zh.org/"
	str += "\n其他：https://github.com/lxn/walk"
	walk.MsgBox(mw, "关于", str, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
