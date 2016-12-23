// internet  网络菜单
package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
import (
	"fmt"
)

//网络测试面板
func (mw *MyMainWindow) dialtest_Trig() {

	//var dlg *walk.Dialog
	fmw := new(MyMainWindow)
	var itte, postdata *walk.TextEdit
	var le *walk.LineEdit
	var GetBtn, PostBtn, cleItte, clearLe, actWeb *walk.PushButton
	var wv *walk.WebView
	//Tab控件
	WebTab := func() {
		tab := itte.Visible()
		if tab == true { //Text 视图
			actWeb.SetText("Web")
			wv.SetVisible(tab)
			itte.SetVisible(!tab)
		} else { //Web视图
			actWeb.SetText("Text")
			wv.SetVisible(tab)
			itte.SetVisible(!tab)
		}
		url := le.Text()
		if url != "" {
			wv.SetURL(url)
		}
	}
	get_data := func(req string) {
		//数据抓取
		if req == "GET" || req == "POST" {
			ctt := "\r\n------" + req + "--------\r\n"
			urlStr := le.Text()
			ctt += urlStr + "\r\n"

			if req == "GET" { //GET
				res, err := http.Get(urlStr)
				if err != nil {
					ctt += err.Error() + "\r\n"
				} else {
					defer res.Body.Close()
					body, _ := ioutil.ReadAll(res.Body)
					ctt += "结果>>\r\n" + string(body) + "\r\n"
				}
				itte.AppendText(ctt)
			} else { //POST
				data := url.Values{"key": {"Value"}, "id": {"123"}}
				fmt.Println(postdata.Text())
				for key, v := range json_decode(postdata.Text()) {
					data.Set(key, v.(string))
				}
				res, err := http.PostForm(urlStr, data)
				if err != nil {
					ctt += err.Error() + "\r\n"
				} else {
					defer res.Body.Close()
					body, _ := ioutil.ReadAll(res.Body)
					ctt += "结果>>\r\n" + string(body) + "\r\n"
				}
				itte.AppendText(ctt)
			}

		}
	}
	//(Dialog{
	(MainWindow{
		AssignTo: &fmw.MainWindow,
		Title:    "网络测试面板",
		MinSize:  Size{600, 460},
		Layout:   VBox{}, //开启VBox时图层才可见
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2}, //图层数据分列
				Children: []Widget{
					Label{
						Text: "拨号主机",
					},
					LineEdit{
						AssignTo: &le,
						OnKeyDown: func(key walk.Key) {
							if key == walk.KeyReturn {
								url := le.Text()
								var desc string
								desc = "\r\n********************\r\n" + now("") + ">>"
								if url == "cls" || url == "CLS" { //清屏
									itte.SetText("")
									return //跳出函数~无返回
								}
								//默认80端口
								if strings.Index(url, ":") == -1 {
									url += ":80"
								}
								protocol := "tcp"
								if strings.Index(url, "//") > -1 {
									praseArr := strings.Split(url, "//")
									protocol = praseArr[0]
									url = praseArr[1]
								} else {
									desc += "tcp//"
								}
								desc += url
								conn, err := net.Dial(protocol, url)
								if err != nil {
									desc += "\r\n调试出错：" + err.Error()
									itte.AppendText(desc)
								} else {
									defer conn.Close()
									ip := conn.LocalAddr().String()
									desc += "\r\n本地地址：" + ip
									desc += "\r\n目标地址：" + conn.RemoteAddr().String()
									/*
										//十分的卡
										//buf := make([]byte, 512)
										buf := make([]byte, 2)
										n, err1 := conn.Read(buf)
										fmt.Println(n, err1)
										fmt.Println(string(buf[:n]))
									*/

									//实现系统自动打开浏览器链接   cmd.exe start http://127.0.0.1/
									//go_links := "http://127.0.0.1:80/"
									//fmt.Println(exec.Command(go_links).Run())
									fmt.Println(exec.Command("start", "http://"+ip+"/").Run())
									//fmt.Println(exec.Command("ipconfig").Run())
									//open_url("http://127.0.0.1/")

									itte.AppendText(desc)
								}
							}
						},
					},
				},
			},
			Composite{
				Layout: Grid{Columns: 6}, //图层数据分列
				Children: []Widget{
					PushButton{
						AssignTo: &GetBtn,
						Text:     "GET",
						OnClicked: func() {
							get_data("GET")
						},
					},
					PushButton{
						AssignTo: &PostBtn,
						Text:     "POST",
						OnClicked: func() {
							get_data("POST")
						},
					},
					PushButton{
						AssignTo: &clearLe,
						Text:     "清理输入框",
						OnClicked: func() {
							le.SetText("")
						},
					},
					PushButton{
						AssignTo: &cleItte,
						Text:     "清理结果",
						OnClicked: func() {
							itte.SetText("")
						},
					},
					PushButton{
						AssignTo: &actWeb,
						Text:     "Web",
						OnClicked: func() {
							WebTab()
							/*
								tab := itte.Visible()
								if tab == true { //Text 视图
									actWeb.SetText("Web")
									itte.SetVisible(!tab)
									wv.SetVisible(tab)
								} else { //Web视图
									actWeb.SetText("Text")
									itte.SetVisible(tab)
									wv.SetVisible(!tab)
								}
								url := le.Text()
								if url != "" {
									wv.SetURL(url)
								}
							*/
						},
					},
					CheckBox{
						Name:    "postHide",
						Text:    "显示POST数据框",
						Checked: false,
					},
				},
			},
			Composite{
				Visible: Bind("postHide.Checked"),
				Layout:  Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "post数据",
					},
					TextEdit{
						AssignTo: &postdata,
					},
				},
			},
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						ColumnSpan: 2,
						Text:       "结果:",
					},
					TextEdit{
						AssignTo:   &itte,
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Remarks"),
					},
					WebView{
						AssignTo: &wv,
						Visible:  false,
					},
				},
			},
		},
	}).Run()
}
func itt() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}
