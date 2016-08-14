// server
package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//以后可设置为常量
var cmdTip string = `可设置的r(server -r x/简化)参数如下
	-r version/v 版本号
	-r start/s 开启服务器
	-r time/t 显示系统时间
	-r port/p 显示服务器端口号
	-r author/a	作者
	-r infomation/i 系统信息
	-port n 设置服务器端口号
	`

const (
	V = "1.2.1"
	A = "Joshua Doeeking Conero"
	I = `
		该项目主要是学习golang，便用其来开发一个小型的服务器。可独立使用，原来是在gogui(go walk学习)项目下发展而来的。
	`
)

//端口号
var PORT string

func main() {
	//web_server()
	historyPort := get_port("")
	var port = flag.String("port", historyPort, "服务器默认端口2020，但会加载上一次设置的端口号")
	var cmd = flag.String("r", cmdTip, "命令")
	flag.Parse()

	var p string = *port
	var r string = *cmd
	if r == cmdTip {
		fmt.Println(r)
	} else {
		//执行命令
		exec_cmd(r)
	}
	if p != historyPort {
		fmt.Println(get_port(p))
	}

}

func get_port(set string) string {
	var fname string = "./src/port.log"
	//设置端口号
	if set != "" {
		fmt.Println(set)
		x, _ := strconv.Atoi(set)
		fmt.Println(x + 50)
		//端口范围1-65535
		if x < 1 || x > 65535 {
			return "端口无效，请设置数字[1-65535]"
		}
		PORT = set
		err := put_content(fname, set)
		fmt.Println(err)
		if err != nil {
			return err.Error()
		}
		return "端口成功设置为了" + set
	}
	//返回端口号
	var ret = get_content(fname)
	//默认为2020
	if ret == "" {
		ret = "2020"
	}
	PORT = ret
	return ret
}
func exec_cmd(cmd string) {
	if cmd == "version" || cmd == "v" {
		fmt.Println("V" + V)
	} else if cmd == "start" || cmd == "s" {
		fmt.Println("正在启动端口为：" + PORT + "的服务器...")
		web_server()
	} else if cmd == "time" || cmd == "t" {
		fmt.Println(now("f"))
	} else if cmd == "port" || cmd == "p" {
		fmt.Println(PORT)
	} else if cmd == "author" || cmd == "a" {
		fmt.Println(A)
	} else if cmd == "infomation" || cmd == "i" {
		fmt.Println(I)
	}
}

//服务器处理
type web struct{}

//参数 server、client>>>
func (h web) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("服务器运行成功...")
	fmt.Fprint(w, "Hello My Web!")
	fmt.Println(r)
}
func web_server() {
	var w web
	ctt := "\r\n" + time.Now().String() + "  " + PORT + "\r\n"
	s := &http.Server{
		Addr:           ":" + PORT,
		Handler:        w,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		ctt += "\r\n系统错误"
		fmt.Println(ctt)
	}
}

func server_guaqi() {
	fmt.Println("Hello World!")
}
