// server
package main

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"strings"
	"time"
)
import (
	"github.com/axgle/mahonia"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var itte *walk.TextEdit

//网络API模拟   网络服务页面
func (mw *MyMainWindow) ApiServer() {
	api := Addr{ip_info(), "2017"}
	fmw := new(MyMainWindow)
	//var itte *walk.TextEdit
	var leIp, lePort, leCmd *walk.LineEdit
	var exc_server, sbmt_http, submit, cancel, clearCmd *walk.PushButton
	//提前申明
	cmd_action := func() {
		//原始数据乱码需要进行代码转换
		itte.AppendText("\r\n" + now("f") + "\r\n")
		output := func(cmd *exec.Cmd, enc string) {
			itte.AppendText("\r\n")
			ret, err := cmd.Output()
			if enc == "UTF-8" || enc == "utf-8" {
				gbk2utf := mahonia.NewDecoder("GBK")
				if err != nil {
					itte.AppendText(gbk2utf.ConvertString(err.Error()))
					return
				}
				itte.AppendText(gbk2utf.ConvertString(string(ret)))
			} else {
				if err != nil {
					itte.AppendText(err.Error())
					return
				}
				itte.AppendText(string(ret))
			}

		}
		ipt := leCmd.Text()
		if ipt == "cls" { //清屏
			itte.SetText("")
			return
		}
		itte.AppendText("\r\n原始命令行：" + ipt)
		iptArr := strings.Split(ipt, " ")
		legth := len(iptArr)
		if legth > 2 {
			str := "\r\n原始命令行：" + iptArr[0] + " " + iptArr[1]
			if iptArr[2] == "" || iptArr[2] == "none" {
				cmd := exec.Command(iptArr[0], iptArr[1])
				itte.AppendText(str)
				output(cmd, "UTF-8")
			} else {
				cmd := exec.Command(iptArr[0], iptArr[1], iptArr[2])
				str += " " + iptArr[2]
				itte.AppendText(str)
				output(cmd, "UTF-8")
			}
		} else if legth == 2 {
			itte.AppendText("\r\n解析命令为：./server/server " + iptArr[0] + " " + iptArr[1])
			cmd := exec.Command("./server/server", iptArr[0], iptArr[1])
			output(cmd, "GBK")
		} else if legth == 1 {
			itte.AppendText("\r\n解析命令为：" + iptArr[0])
			cmd := exec.Command(iptArr[0])
			output(cmd, "UTF-8")
		} else {
			cmd := exec.Command("./server/server", "-r", "start")
			itte.AppendText("\r\n解析命令为：./server/server -r start")
			output(cmd, "GBK")
		}
	}
	(MainWindow{
		AssignTo: &fmw.MainWindow,
		Title:    "接口服务器",
		MinSize:  Size{600, 460},
		Layout:   VBox{}, //开启VBox时图层才可见
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 4}, //图层数据分列
				Children: []Widget{
					Label{
						Text: "服务器IP:",
					},
					LineEdit{
						AssignTo: &leIp,
						Text:     api.ip,
					},
					Label{
						Text: "服务器端口号:",
					},
					LineEdit{
						AssignTo: &lePort,
						Text:     api.port,
					},
					Label{
						Text: "cmd命令",
					},
					LineEdit{
						ColumnSpan: 3,
						AssignTo:   &leCmd,
						OnKeyDown: func(key walk.Key) {
							//确认键
							if key == walk.KeyReturn {
								cmd_action()
							}
						},
					},
				},
			},
			Composite{
				Layout: Grid{Columns: 5}, //图层数据分列
				Children: []Widget{
					PushButton{
						AssignTo: &exc_server,
						Text:     "运行server.exe", //此方法运行的exe可以隐藏起来
						OnClicked: func() {
							go cmd_action() //开启新的进程
						},
					},
					PushButton{
						AssignTo: &clearCmd,
						Text:     "重置命令框",
						OnClicked: func() {
							leCmd.SetText("")
						},
					},
					PushButton{
						AssignTo: &sbmt_http,
						Text:     "http",
						OnClicked: func() {
							port := lePort.Text()
							go webserv(port) //多进程-程序并发
						},
					},
					PushButton{
						AssignTo: &submit,
						Text:     "纯链接运行",
						OnClicked: func() {
							ip := leIp.Text()
							port := lePort.Text()
							go submit_clear(ip, port) //轻量级程序并发
						},
					},
					PushButton{
						AssignTo: &cancel,
						Text:     "关闭",
						OnClicked: func() {
							fmw.Close()
						},
					},
				},
			},
			TextEdit{
				AssignTo: &itte,
			},
		},
	}).Run()
}

//纯链接运行
func submit_clear(ip, port string) {
	ctt := "\r\n" + now("") + "  " + ip + ":" + port + "\r\n"
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		ctt += "错误:>>" + err.Error()
		itte.AppendText(ctt)
	} else {
		defer ln.Close()
		conn, err := ln.Accept()
		if err != nil {
			ctt += "读取失败:>>" + err.Error()
		} else {
			ctt += conn.RemoteAddr().String()
		}
		a := ln.Addr()
		ctt += "\r\n地址" + a.String() + " ，网络" + a.Network()
		itte.AppendText(ctt)
	}
}

//服务器处理
type web struct{}

//网页服务器(轻量级/) 2016年7月19日使用 go _goroutine_ 是由 Go 运行时环境管理的轻量级线程。实现单进程的影响
func webserv(port string) {
	mkdir("./www/") //网页脚本位置
	var w web
	ctt := "\r\n" + now("") + "  " + port + "\r\n"
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        w,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	itte.AppendText(ctt)
	ctt = ""
	ctt += s.ListenAndServe().Error()
	itte.AppendText(ctt)

}

//参数 server、client>>>
func (h web) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello My Web!\r\n")
	//itte.AppendText("ddhdh>>>")//程序出错
	fmt.Fprint(w, r)
}

func server_guaqi() {
	fmt.Println("Hello World!")
}
