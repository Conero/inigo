// page
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Page struct{}

// page 头部
func (p *Page) head(title, app string) string {
	var xhtml string
	xhtml = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>` + title + `</title>
			` + app + `
		</head>
	`
	return xhtml
}

// page body
func (p *Page) body(bodyString string) string {
	var xhtml string
	xhtml = `
		<body>
			` + bodyString + `
		</body>
		</html>
	`
	return xhtml
}

// os 专项页面
func (p *Page) html_os() string {
	env := os.Environ()
	envXhtml := ""
	for _, v := range env {
		envXhtml += "<li>" + v + "</li>"
	}
	bodyStr := `
		<nav>
			<a href="#envs">envs</a>
		</nav>
		<div id="envs">
			<h4>系统常量</h4>
			<ol>
				` + envXhtml + `
			</ol>
		</div>
	`
	xhtml := p.head("OS"+randInt(), "")
	xhtml += p.body(bodyStr)
	return xhtml
}

// 首页
func (p *Page) index(w http.ResponseWriter, r *http.Request) string {
	xhtml := p.head("XHelper By Joshua", "")
	bhtml := `
		<div id="basc_info">
			<h4>基本信息</h4>
			<br>时间：` + time.Now().String() + `
			<br>请求地址：` + r.URL.String() + `
			<br>主页：` + r.Host + `
			<br>方式：` + r.Method + `
			<br>Proto：` + r.Proto + `
			<br>Referer：` + r.Referer() + `
			<br>RemoteAddr：` + r.RemoteAddr + `
			<br>RequestURI：` + r.RequestURI + `
			<br>代理：` + r.UserAgent() + `
		</div>
	`
	tmp := ""
	for k, v := range w.Header() {
		tmp += "<li>" + k + " : " + v[0] + "</li>"
	}
	if tmp != "" {
		bhtml += tmp
	}
	bhtml += `
		<div id="about">
			<h4>关于插件</h4>
			<br>作者：` + XHp.author + `
			<br>名称：` + XHp.name + `
			<br>日期：` + XHp.date + `
			<br>版本：` + XHp.version + `
		</div>
	`
	g := r.URL.Query()
	dir := g.Get("dir")
	if dir != "" && hasFile(dir) {
		bhtml += "<p>生成zip包：" + XHp.zip(dir)
	}
	xhtml += p.body(bhtml)
	return xhtml
}

func pguiqi() string {
	fmt.Println(randInt())
	return randInt()
}
