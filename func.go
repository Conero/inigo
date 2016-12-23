// func
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 生成随机数
func randInt() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100000)
	return strconv.Itoa(num)
}

// 通过解析dir解析出zip文件名
func getName(dir string) string {
	dir = strings.Replace(dir, "/", "\\", -1)
	tmp := strings.Split(dir, "\\")
	name := tmp[len(tmp)-1]
	if name == "" {
		name = tmp[len(tmp)-2]
	}
	return name
}

// 自动生成目录
func mkdir(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		err := os.Mkdir(dir, os.ModeDir)
		if err != nil {
			return false
		} else {
			return true
		}
	}
	return true
}

// 文件或者目录是否存在
func hasFile(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		return false
	} else {
		return true
	}
}

func isDir(p string) bool {
	f, err := os.Stat(p)
	if err != nil {
		return false
	}
	return f.IsDir()
}

//写文件/可新建，主要是实现文件覆盖~重写
func put_content(name, content string) error {
	err := ioutil.WriteFile(name, []byte(content), 0x644)
	return err
}

//读入文件
func get_content(name string) string {
	ctt, err := ioutil.ReadFile(name)
	if err != nil {
		log := name + "文件读取失败！=》" + err.Error()
		write("./runtime/error.log", log)
	}
	cttStr := string(ctt)
	return cttStr
}

//写文件/可新建，主要是实现文件追加
func write(name, content string) (int, error) {
	fl, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer fl.Close()
	n, err := fl.Write([]byte(content))
	if err == nil && n < len(content) {
		return 0, err
	}
	return n, err
}
func funcguaqi() {
	fmt.Println("Hello World!")
}
