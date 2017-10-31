/* @Joshua Conero
 * @2017年1月19日 星期四
 * @ini 文件解释器
 */

package ini

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"regexp"
	"strings"
	"time"
)

// 解析以后的数据
var paserData map[string]interface{}

// 当前计算机用户信息
type baseUser struct {
	__DIR__  string // 当前目录
	__USER__ string // 当前用户
	__GID__  string
	__NAME__ string
	__UID__  string
}

// 编译器属性
type Compiler struct {
	filename  string    // 文件名称
	content   string    // 内容
	mtime     time.Time // 用于计算时间差-效率评估量
	endtime   time.Time // 完成时间
	seekCount int       // 遍历次数
	lineCount int       // 行数
}

// 压缩器属性
type Compresser struct {
	filename  string    // 文件名称
	content   string    // 内容
	mtime     time.Time // 用于计算时间差-效率评估量
	endtime   time.Time // 完成时间
	seekCount int       // 遍历次数
	lineCount int       // 行数
}

var currentConf *Conf            // 当前的文件
var currentCompiler Compiler     // 当前编译器对象
var currentCompresser Compresser // 当前压缩器对象

// 当前用户
var currentUser baseUser

// 构造函数-打开配置文件
func Open(fname string) *Conf {
	currentConf = &Conf{
		filename:       fname,
		content:        "",
		mtime:          time.Now(),
		Cdt:            make(map[string]interface{}),
		RenderVariable: make(map[string]string),
		FullComplier:   true,
		CloseRender:    false,
	}
	return currentConf
}

// 构造函数-设置字符串
func Content(vtext string) *Conf {
	currentConf = &Conf{
		filename:       "",
		content:        vtext,
		mtime:          time.Now(),
		Cdt:            make(map[string]interface{}),
		RenderVariable: make(map[string]string),
		FullComplier:   true,
		CloseRender:    false,
	}
	return currentConf
}

// 解析文本
func (c *Conf) Parse() (int, map[string]interface{}) {
	var status int = 0
	paserData = make(map[string]interface{})
	seekCount := 0
	lineCount := 0
	if c.filename != "" {
		c.mtime = time.Now()
		// 逐行读取文件
		fh, err := os.Open(c.filename)
		multiLine := false // 多行注释开启
		if err == nil {
			buf := bufio.NewReader(fh)
			for {
				line, err := buf.ReadString('\n')
				// 程序跳转前检测是否出错，出错直接中断循环，避免还没有检查错误时便继续进入循环(死循环)
				if err != nil {
					break
				}
				line = strings.TrimSpace(line)
				lineCount = lineCount + 1
				// 多行注释开始
				if line == ">>>" {
					multiLine = true
				} else if line == "<<<" { // 注释结束
					multiLine = false
					continue
				}
				// 循环跳出
				if line == "" || strings.Index(line, "#") == 0 || multiLine { // # 字符换跳过
					continue
				}
				key := ""
				// 模式一 key = v1 , v2 , v3	  ; key = v
				eqSigner := regexp.MustCompile(`=+`)
				if eqSigner.Match([]byte(line)) {
					// 删除行之间的所有空格
					kbSinger := regexp.MustCompile(`\s`)
					line = string(kbSinger.ReplaceAll([]byte(line), []byte("")))
					strLen := strings.Index(line, "=")
					key = line[0:strLen]
					line = strings.TrimSpace(line[strLen+1:])
					line = c.Render(line)
					if strings.Index(line, ",") > -1 {
						paserData[key] = strings.Split(line, ",")
					} else {
						paserData[key] = line
					}
				} else {
					// 模式二 key v1 v2 v3; key v
					strLen := strings.Index(line, " ")
					key = line[0:strLen]
					line = strings.TrimSpace(line[strLen+1:])
					line = Render(line)
					// 删除多空格
					mkbSinger := regexp.MustCompile(`[\s]{2:}`)
					if mkbSinger.Match([]byte(line)) {
						line = string(mkbSinger.ReplaceAll([]byte(line), []byte("")))
					}
					if strings.Index(line, " ") > -1 {
						paserData[key] = strings.Split(line, " ")
					} else {
						paserData[key] = line
					}
				}
				seekCount = seekCount + 1
			}
			c.seekCount = seekCount
			c.lineCount = lineCount
			c.endtime = time.Now()
		}
		status = 1
	}
	c.Cdt = paserData
	return status, paserData
}

// 设置编译文件名称
func (c *Conf) SetCompileFname(fname string) *Conf {
	currentCompiler.filename = fname
	return c
}

// 是否关闭全编译， 即系统默认 有{var} 但不是系统常量忽略
func (c *Conf) CloseFullComplier(isClose bool) {
	c.FullComplier = !isClose
}

// 是否关闭编译功能(关闭后不再检查变量)
func (c *Conf) IsCloseRender(setClose bool) {
	c.CloseRender = setClose
}

// 编译ini文件
func (c *Conf) Compile() (bool, string) {
	fname := c.filename
	if len(fname) > 0 {
		currentCompiler.mtime = time.Now()
		contentBytes, err := ioutil.ReadFile(fname)
		// 文件不存在
		if err != nil {
			return false, fname + "文件读取失败"
		}
		content := string(contentBytes)
		content = Render(content)
		newFileName := currentCompiler.filename
		if len(newFileName) < 1 {
			ext := path.Ext(fname)
			newFileName = strings.Replace(fname, ext, ".compile"+ext, -1)
		}
		fname = newFileName
		// 文件写入
		isSuccess := ioutil.WriteFile(newFileName, []byte(content), 0x644)
		if isSuccess != nil {
			return false, fname + "文件写入失败"
		}
		currentCompiler.content = content
		currentCompiler.endtime = time.Now()
		return true, fname
	}
	return false, fname
}

// 设置编译文件名称
func (c *Conf) SetCompressFname(fname string) *Conf {
	currentCompresser.filename = fname
	return c
}

// 压缩ini文件
func (c *Conf) Compress() (bool, string) {
	fname := c.filename
	if len(fname) > 0 {
		// 逐行读取文件
		fh, err := os.Open(fname)
		multiLine := false // 多行注释开启
		content := ""
		if err == nil {
			buf := bufio.NewReader(fh)
			for {
				line, err := buf.ReadString('\n')
				if err != nil {
					break
				}
				line = strings.TrimSpace(line)

				// 多行注释开始
				if line == ">>>" {
					multiLine = true
				} else if line == "<<<" { // 注释结束
					multiLine = false
					continue
				}
				// 循环跳出
				if line == "" || strings.Index(line, "#") == 0 || multiLine { // # 字符换跳过
					continue
				}
				content += line
			}
			newFileName := currentCompresser.filename
			if len(newFileName) < 1 {
				ext := path.Ext(fname)
				newFileName = strings.Replace(fname, ext, ".min"+ext, -1)
			}
			fname = newFileName
			// 文件写入
			isSuccess := ioutil.WriteFile(newFileName, []byte(content), 0x644)
			if isSuccess != nil {
				return false, fname + "文件写入失败"
			}
			currentCompresser.content = content
			currentCompresser.endtime = time.Now()
			return true, fname

		}
	}
	return false, fname + "文件读取失败"
}

// 获取遍历次数
func (c *Conf) GetseekCount() int {
	return c.seekCount
}

// 获取编译次数
func (c *Conf) GetlineCount() int {
	return c.lineCount
}

// 键值获取
func (c *Conf) Get(key string) (interface{}, bool) {
	value, has := c.Cdt[key]
	return value, has
}

// interface{} 转 -> string
func (c *Conf) GetString(key string) string {
	v, exist := c.Get(key)
	value := ""
	if exist {
		switch v.(type) {
		case string:
			value = v.(string)
		case []string:
			value = strings.Join(v.([]string), ",")
		}
	}

	return value
}
func (c *Conf) RenderByKeyStr(key, skey, svalue string) string {
	line := c.GetString(key)
	if len(line) == 0 {
		return line
	}
	keysMatch := regexp.MustCompile(`\{[\$a-zA-Z\d_-]+\}`)
	keys := keysMatch.FindStringSubmatch(line)
	for _, v := range keys {
		oKey := ""
		if v == skey {
			oKey = "{" + skey + "}"
		}
		line = strings.Replace(line, oKey, svalue, -1)
	}
	return line
}

// interface{} 转 -> Array
//func (c *Conf) GetArray(key string) array[]string {}
func (c *Conf) GetArray(key string) []string {
	var tmpArray []string
	v, exist := c.Get(key)
	if exist {
		switch Arr := v.(type) {
		case string:
			tmpArray = append(tmpArray, v.(string))
		case []string:
			tmpArray = v.([]string)
		default:
			out(Arr)

		}
	}
	return tmpArray
}
func (c *Conf) SetRenderVals(data map[string]string) {
	c.RenderVariable = data
}

// 系统常量转换 -> {变量名}
func (c *Conf) Render(linStr string) string {
	if c.CloseRender { // 关闭值渲染
		return linStr
	}
	keysMatch := regexp.MustCompile(`\{[\$a-zA-Z\d_-]+\}`)
	keys := keysMatch.FindAllStringSubmatch(linStr, -1)
	name := ""
	tplKey := ""
	signerMatch := regexp.MustCompile(`\{|\}`)
	for _, v := range keys {
		if len(v) > 0 {
			value := ""
			name = v[0]
			tplKey = string(signerMatch.ReplaceAll([]byte(name), []byte("")))
			switch tplKey {
			case "$dir":
				value = "."
			case "$parentDir":
				value = ".."
			case "__DIR__":
				value = GetBaseDir()
			case "__USER__":
				if len(currentUser.__USER__) < 1 {
					GetBaseDir()
					value = currentUser.__USER__
				} else {
					value = currentUser.__USER__
				}
			case "__GID__":
				if len(currentUser.__GID__) < 1 {
					GetBaseDir()
					value = currentUser.__GID__
				} else {
					value = currentUser.__GID__
				}
			case "__NAME__":
				if len(currentUser.__NAME__) < 1 {
					GetBaseDir()
					value = currentUser.__NAME__
				} else {
					value = currentUser.__NAME__
				}
			case "__UID__":
				if len(currentUser.__UID__) < 1 {
					GetBaseDir()
					value = currentUser.__UID__
				} else {
					value = currentUser.__UID__
				}
			default:
				pv, has := c.Cdt[tplKey]
				if has {
					switch pv.(type) {
					case string:
						value = pv.(string)
					}
				} else if len(c.RenderVariable) > 0 {
					pv1, has1 := c.RenderVariable[tplKey]
					if has1 {
						value = pv1
					}
				} else if c.FullComplier == false {
					return linStr
				}
			}
			linStr = strings.Replace(linStr, name, value, -1)
		}
	}
	return linStr
}

// 系统常量转换 -> {变量名}
func Render(linStr string) string {
	keysMatch := regexp.MustCompile(`\{[\$a-zA-Z\d_-]+\}`)
	keys := keysMatch.FindAllStringSubmatch(linStr, -1)
	name := ""
	tplKey := ""
	signerMatch := regexp.MustCompile(`\{|\}`)
	for _, v := range keys {
		if len(v) > 0 {
			value := ""
			name = v[0]
			tplKey = string(signerMatch.ReplaceAll([]byte(name), []byte("")))
			switch tplKey {
			case "$dir":
				value = "."
			case "$parentDir":
				value = ".."
			case "__DIR__":
				value = GetBaseDir()
			case "__USER__":
				if len(currentUser.__USER__) < 1 {
					GetBaseDir()
					value = currentUser.__USER__
				} else {
					value = currentUser.__USER__
				}
			case "__GID__":
				if len(currentUser.__GID__) < 1 {
					GetBaseDir()
					value = currentUser.__GID__
				} else {
					value = currentUser.__GID__
				}
			case "__NAME__":
				if len(currentUser.__NAME__) < 1 {
					GetBaseDir()
					value = currentUser.__NAME__
				} else {
					value = currentUser.__NAME__
				}
			case "__UID__":
				if len(currentUser.__UID__) < 1 {
					GetBaseDir()
					value = currentUser.__UID__
				} else {
					value = currentUser.__UID__
				}
			default:
				pv, has := paserData[tplKey]
				if has {
					switch pv.(type) {
					case string:
						value = pv.(string)
					}
				}
			}
			linStr = strings.Replace(linStr, name, value, -1)
		}
	}
	return linStr
}

// 获取用于当前路径地址
func GetBaseDir() string {
	if len(currentUser.__DIR__) > 0 {
		return currentUser.__DIR__
	} else {
		user, err := user.Current()
		if err == nil {
			currentUser = baseUser{
				__DIR__:  user.HomeDir,
				__USER__: user.Username,
				__GID__:  user.Gid,
				__NAME__: user.Name,
				__UID__:  user.Uid,
			}
		}
	}
	return currentUser.__DIR__
}
