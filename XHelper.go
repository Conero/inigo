// XHelper
package main

import (
	"archive/zip"
	"bytes"
	"fmt" //测试包
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type XHelper struct {
	version string // 版本
	date    string // 日期
	author  string // 作者
	name    string // 名称
	update  string // 更新记录
}

// 构造函数
func XHelperInit() *XHelper {
	X := &XHelper{
		version: "V1.5.1",
		date:    "20160927",
		author:  "Joshua Conero",
		name:    "XHelper PHP插件帮助工具",
		update:  "1.0.0=20160927;1.1.0=20161002;1.1.5=20161007;1.5.1=20161011",
	}
	return X
}

// 解析命令行
func (X *XHelper) command() string {
	args := os.Args
	cmd := ""
	val := ""
	if len(args) >= 2 {
		cmd = args[1]
	}
	if len(args) >= 3 {
		val = args[2]
	}
	switch cmd {
	case "-zip": // zip 压缩
		_, err := os.Stat(val)
		if err == nil {
			return X.zip(val)
		} else {
			return err.Error()
		}
	case "-server": // 内置服务器
		if val == "" {
			val = "8080"
		}
		//go X.myServer(val)// 单独此句无效
		X.myServer(val)
	case "-fserver": // 内置文件服务器
		if val == "" {
			val = "7777"
		}
		X.fserver(val)
	case "-version":
		return X.version
	case "-name":
		return X.name
	case "-config":
		mkdir("./XHelper/")
		return cf.run(val)
	case "-i":
		return X.getInfo(val)
	default:
		// 自动检查是否有扩展可以
		if cmd != "" {
			exeName := strings.Replace(cmd, "-", "", -1)
			exeName = "./extends/" + exeName + ".exe"
			if hasFile(exeName) {
				etd := exec.Command(exeName, val)
				ret, err := etd.Output()
				if err != nil {
					return err.Error()
				} else {
					return string(ret)
				}
			}
		}
		// 否则输出有关系统
		desc := X.getDecrip()
		fmt.Println(desc)
	}
	return ""
}

// 压缩目录为zip
func (X *XHelper) zip(dir string) string {
	buf := new(bytes.Buffer)
	myzip := zip.NewWriter(buf)

	flist, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fs := range flist {
		if !fs.IsDir() {
			fname := fs.Name()
			f, _ := myzip.Create(fname)
			ct, _ := ioutil.ReadFile(dir + "/" + fname)
			f.Write(ct)
		}

	}

	myzip.Close()
	mkdir("./_runtime")
	zipname := "./_runtime/" + getName(dir) + "_" + randInt() + ".zip"
	f, err := os.OpenFile(zipname, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
	fmt.Println(dir)
	return zipname
}

//服务器处理
type web struct{}

// 开启内置服务器
func (X *XHelper) myServer(port string) {
	var w web
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        w,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println(port + "端口号的服务已经启动")
	log.Fatal(s.ListenAndServe())
}

// 内置文件服务器
func (X *XHelper) fserver(port string) {
	fmt.Println(port, "端口号文件服务器已经开启")
	// 文件服务器
	mkdir("./tmp/")
	http.Handle("/", http.FileServer(http.Dir("./tmp/")))
	/*
		hd := http.Handler{
			ServeHTTP: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "55555")
			},
		}
		http.ListenAndServe(":"+port, hd)
	*/
	http.ListenAndServe(":"+port, nil)
}

//参数 server、client>>>
func (h web) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if url == "/" {
		url = "/index"
	}
	p := new(Page)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	xhtml := ""
	switch url {
	case "/os":
		xhtml += p.html_os()
	case "/index":
		xhtml += p.index(w, r)
	default:
		fmt.Fprintln(w, r.URL.String())
	}
	fmt.Fprintln(w, xhtml)
}

// 返回XHelper软件详情基本概述
func (X *XHelper) getDecrip() string {
	content := `
		系统介绍>>
			名称：` + X.name + `
			作者：` + X.author + `
			版本：` + X.version + `
			时间：` + X.date + `
		命令参数介绍>>
			1.XHelper.exe -zip dir/目录			压缩目录下所有文件
			2.XHelper.exe -server port/端口号	开启以port为端口号的服务器
			3.XHelper.exe -fserver port/端口号	开启以port为端口号的文件服务器
			4.XHelper.exe -version				查看服务版本
			5.XHelper.exe -name					查看服务命名
			6.XHelper.exe -config ini/xml/json/delete	选择配置格式，生成名字conf.ini/等;delete删除配置文件;
			7.XHelper.exe -i 文件名[XHelper.md] 		 	生成系统的说明详情文	
			8.XHelper.exe -插件名称 参数 		 			系统集成群化处理；目录 -/extends/插件.exe			
		`
	return content
}

// 获取XHelper.exe 的说明文件
func (X *XHelper) getInfo(v string) string {
	if v == "" {
		v = "./XHelper.md"
	} else if isDir(v) {
		v += "/XHelper.md"
	}
	content := X.getDecrip()
	err := put_content(v, content)
	if err == nil {
		return BR + "”" + v + "“文件名已经生成"
	}
	return BR + v + BR + err.Error()
}
func guaqi() {
	fmt.Println("ppp")
}
