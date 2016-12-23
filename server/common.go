// common 公用函数
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
	工具
	2016年6月21日
	单行语句多行=》  ``
	"reflect" 反射
*/
//接口
//type JSON interface{} //bug# type JSON does not support indexing
type JSON map[string]interface{} //bug# 多重JSON依然存在问题

func now(sel string) string {
	retVal := ""
	t := time.Now()
	tstr := t.String()
	tslice := strings.Split(tstr, " ") //time切块
	dstr := tslice[0]
	tistr := strings.Split(tslice[1], ".")[0]
	yslice := strings.Split(dstr, "-") //year切块
	//fmt.Println(tslice)
	/*
		for i, v := range tslice {
			fmt.Println(i, v)
		}
	*/
	/*
		if sel == "T" {
			retVal = dstr + " " + tistr
		} else if sel == "Y" { //年
			retVal = strconv.Itoa(t.Year())
		} else if sel == "M" { //月
			retVal = yslice[1]
		} else if sel == "D" { //日
			retVal = yslice[2]
		} else if sel == "t" { //时间
			retVal = tistr
		} else if sel == "h" { //小时
			retVal = strconv.Itoa(t.Hour())
		} else if sel == "m" {
			retVal = strconv.Itoa(t.Minute())
		} else if sel == "s" {
			retVal = strconv.Itoa(t.Second())
		} else {
			retVal = dstr
		}
	*/
	switch sel {
	case "T":
		retVal = dstr + " " + tistr
	case "Y":
		retVal = strconv.Itoa(t.Year())
	case "M":
		retVal = yslice[1]
	case "D":
		retVal = yslice[2]
	case "t":
		retVal = tistr
	case "h":
		retVal = strconv.Itoa(t.Hour())
	case "m":
		retVal = strconv.Itoa(t.Minute())
	case "s":
		retVal = strconv.Itoa(t.Second())
	case "f":
		retVal = tstr
	default:
		retVal = dstr
	}
	/*
		无效
		const layout = "2016-06-21"
		fmt.Println(t.Format(layout))
	*/
	return retVal
}

//返回秒用于计算程序用时,参数为0时返回当前的毫秒，否则返回计算后的秒差
func sec(start float64) float64 {
	t := time.Now()
	ns := float64(t.Nanosecond())
	ms := ns / math.Pow10(6) //1ms = 10^6ns
	if start == 0 {
		return ms
	}
	ds := (ms - start) / math.Pow10(3)
	ds = round(ds, 5)
	return ds
}

//字符串方法处理float等长数据 规定位数
func round(num float64, b int) float64 {
	if b == 0 {
		return float64(int(num))
	}
	n2t := int(num * math.Pow10(b))    //num转换数
	base := int(num * math.Pow10(b+1)) //四舍五入的最后一位数
	base = int(math.Abs(float64(base - n2t*10)))
	if base > 5 {
		n2t += 1
	}
	num = float64(int(num)) + float64(n2t)/float64(math.Pow10(b))
	return num
}

//json=>string 返回string
func json_encode() {

}

//array=> json 返回json 对象  原生解码十分困难/可借助第三方库
func json_decode(str string) JSON {
	var feek map[string]interface{}
	var jtmp = []byte(str)
	err := json.Unmarshal(jtmp, &feek)
	if err != nil {
		log := now("f") + `:
		json_decode解析json失败
		`
		write("./runtime/error.log", log)
	}
	return feek
}

//多重解码失败时，将再次自动转化为json
func jsua(v interface{}) JSON {
	vcld, _ := json.Marshal(v)
	str := string(vcld)
	return json_decode(str)

}

//解码格式为home=$pss2300a=:9888 单解码
func str2json(str string) JSON {
	tmp := strings.Replace(strings.Replace(str, "$", "\",\"", -1), "=", "\":\"", -1)
	tmp = "{\"" + tmp + "\"}"
	fmt.Println(tmp)
	return json_decode(tmp)
}

//失败~ 实例 https://go-zh.org/pkg/encoding/json/#Decoder
func json_decodex(str string) {
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(str))
	fmt.Println(dec)
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

//文件/目录检测
func is_dir(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		return false
	} else {
		return true
	}
}

//生成文件/目录(可检测)?文件未测试是否可行
func mkdir(name string) bool {
	if is_dir(name) == false {
		err := os.Mkdir(name, os.ModeDir)
		if err != nil {
			return false
		}
	}
	return true
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

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

//打开系统链接
func open_url(uri string) error {
	fmt.Println(runtime.GOOS)
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

/*
func jsua_decode(str string) {
	//var feek map[string]jbase
	feek := make(map[string]interface{})
	var jtmp = []byte(str)
	err := json.Unmarshal(jtmp, &feek)
	if err != nil {
		log := now("f") + `:
		json_decode解析json失败
		`
		write("./runtime/error.log", log)
	}
	fmt.Println("jsua_decode>>>")
	fmt.Println(feek)
	//类型判断
	//fmt.Println(feek["data"].(type))// 仅仅用在switch语句中
	fmt.Println(reflect.TypeOf(feek["data"]))
	v, _ := json.Marshal(feek["data"])
	fmt.Println(string(v))
	fmt.Println("<<<jsua_decode")
	var tmp jsua
	return tmp
}
*/
/*
func json_encode(str string, feek map[string]interface{}) JSON {
	var jtmp = []byte(str)
	var J JSON
	json.Unmarshal(jtmp, &J)
	//var m = J // 示例 map[Dial:baidu.com:80]
	fmt.Println(J)
	//fmt.Println("Emmi", str)
	//fmt.Println(J.Dial)
	//fmt.Println(m["Dial"])
	fmt.Println(J["Dial"])
	fmt.Println(J["week"])
	return J
}
*/
func tool_guaqi() {
	fmt.Println("tool.go挂起,包没有使用却引入时会报错")
}
