/**
文本域二进制测试
*/
package main

import (
	"os"
	"io/ioutil"
	"bytes"
	"encoding/binary"
	"fmt"
)

// binary for ini test version 1
type BfIniTestV1 struct {
	filename string				// 文案名称
}
// test
func BTv1(filename string)  *BfIniTestV1{
	return &BfIniTestV1{
		filename:filename,
	}
}
// 读取文本
func (bf *BfIniTestV1) Read() string {
	text := ""
	if _, err := os.Stat(bf.filename); os.IsExist(err){
		ct, err2 := ioutil.ReadFile(bf.filename)
		if err2 == nil{
			err3 := binary.Read(bytes.NewBuffer(ct),binary.LittleEndian, &text)
			if err3 != nil{
				fmt.Println(err3.Error())
			}
		}
	}
	return text
}
// 写入文本
func (bf *BfIniTestV1) Write(content string)  bool{
	isSuccess := false
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, []byte(content))
	if err == nil{
		isSuccess = true
		//text := buf.String()
		//err := ioutil.WriteFile("./"+bf.filename, buf.Bytes(), 0666)
		err := ioutil.WriteFile("./"+bf.filename, []byte(fmt.Sprintf("% x", buf.Bytes())), 0666)
		fmt.Println("./"+bf.filename, buf.String())
		if err != nil{
			isSuccess = false
			fmt.Println(err.Error())
		}
	}else{
		fmt.Println(err.Error())
	}
	return isSuccess
}

func main() {
	args := os.Args
	action := ""
	filename := ""
	param := ""
	if len(args) > 2{
		action = args[1]
	}
	if len(args) > 3{
		filename = args[2]
	}
	if len(args) > 4{
		param = args[3]
	}
	switch action {
	case "write":	// 读取文件
		bf := BTv1(filename)
		// --f=name
		content := "文件写入测试，二进制"
		if len(param) > 4 && param[0:3] == "--f="{
			if isS,Ctt := getFileContent(param[4:]);isS{
				content = Ctt
			}
		}
		bf.Write(content)
	case "read":
		bf := BTv1(filename)
		ctt := bf.Read()
		fmt.Println(ctt)
		if len(param) > 0{
			ioutil.WriteFile(param, []byte(ctt), 0666)
		}
	default:
		fmt.Println(`
	二进制问文件读写测试
	. write filename <--f=文件源名称|字段>
	. read  filename <srcname>
		`)
	}
}

func getFileContent(filename string) (bool, string) {
	b1, err := ioutil.ReadFile(filename)
	isSuccess := false
	content := ""
	if err == nil{
		isSuccess = true
		content = string(b1)
	}
	return isSuccess, content
}