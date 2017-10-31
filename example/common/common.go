// 共享库或参数(全局变量或函数)
package common

import (
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var root_path string

func GetRootPath() string {
	if len(root_path) == 0 {
		args := os.Args
		args_index1 := args[0]
		args_index1 = strings.Replace(args_index1, "\\", "/", -1)
		tmpArr := strings.Split(args_index1, "/")
		path := strings.Join(tmpArr[0:len(tmpArr)-1], "/") + "/"
		root_path = path
	}
	return root_path
}

//读入文件
func GetContent(name string) string {
	ctt, err := ioutil.ReadFile(name)
	if err != nil {
		return name + "文件读取发生错误: " + err.Error()
	}
	cttStr := string(ctt)
	return cttStr
}

//写文件/可新建，主要是实现文件覆盖~重写
func PutContent(name, content string) error {
	err := ioutil.WriteFile(name, []byte(content), 0x644)
	return err
}

// 文件复制
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return -1, err
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return -1, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// 是否为目录
func IsDir(p string) bool {
	f, err := os.Stat(p)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// 多目录制造
func MkDirs(pathname string) error {
	if IsDir(pathname) {
		return nil
	}
	// 检测是否有文件，如有文件则进行过滤
	if path.Ext(pathname) != "" {
		pathname, _ = path.Split(pathname)
	}
	return os.MkdirAll(pathname, 0x644)
}

// 时间格式对应表转换表 小写为普通前零导式，大写反之
func Date(format string) string {
	formatTmp := map[string]string{
		"y": "2006",
		"m": "01",
		"d": "02",
		"h": "15",
		"i": "04",
		"s": "05",
		"Y": "06",
		"M": "1",
		"D": "2",
		"H": "3",
		"O": "03",
		"I": "4",
		"S": "5",
	}
	for k, v := range formatTmp {
		format = strings.Replace(format, k, v, -1)
	}
	return time.Now().Format(format)
}

// 生成随机数
func RandInt() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100000)
	return strconv.Itoa(num)
}
