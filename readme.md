# go ini 文件解析库

* AUTHOR    = Joshua Conero
* VERSION   = 0.1.0        
* NAME      = go ini 文件解析库 
* START     = 2017-01-19   
* COPYRIGHT = @Conero      


## 来源
- 2017年3月18日 星期六
- 从原 console 仓库独立出来

### 全解析模型
- 对应函数： FullParse
- 数据模型： 树形图
- 与go结合对象： 基础 Map 

<br /> 		Map	-> [string]interface{}
<br /> 			Map	-> [string]string						string
<br /> 			Map	-> [string]Array						array
<br /> 			Map	-> Map[string]interface{}
- 设想法： 
1. 	生成单个 Map 值, 再于基级结合				主要问题： 多级Map自己无法与基基对象相链接
2.	先解析所有数据对象，再根据级别组合


## 严重错误记录

### 2017年3月22日 星期三
- 逐行读取进入死循环: 在 ini 文件最后一行为注释行/获取持续空行时 程序进入死循环
- 原函数代码如下

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
					line = Render(line)
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
				if err != nil {
					break
				}
			}
			c.seekCount = seekCount
			c.lineCount = lineCount
			c.endtime = time.Now()
		}
		status = 1
	}
	return status, paserData
}

- 修改方式： 逐行读取时优先检测错误性

## 版本更新以及说明

###  V 0.1.4 	(2017年3月22日 星期三)
*  逐行读取进入死循环: 在 ini 文件最后一行为注释行/获取持续空行时 程序进入死循环

###  V 0.1.6 	(2017年3月23日 星期四)
* 移除全局变量 parseData 等变量，使库支持"多进程"调用， 即对象与全局变量分离
* 其他程序修复
* (类似 / alpha 版)

### V 0.2.1 (2017年3月29日 星期三)

* 版本优化，将解析符号变为可设置常量(内嵌代码中)
* 结构优化，将代码移入 src 中
* 重复服务版本号的意义： V a.b.c => 
			a 有标志性改变时增加版本号			版本号增加
			b 功能性代码改变 					版本号增加		
			c 代码提交带git/且有记录是 		  版本号增加
* 实现作用域功能
* 新增 ini -> json 字符换改变
* 新增版本对应号

### V 0.2.2 (2017年3月30日 星期四)
* 新增全局的 ToJson 函数用于实现 map[string]interface{} -> Json 字符串
* Conf.ToJson 具体实现
* Conf.Parse 解析程序调试以及优化， 实现二级作用域对象生成， 以及 {} 相关语法的测试以及生成
* 对应引入的 common 库 进行标准化

### V 0.2.3 (2017年3月31日 星期四)
* 全级解析分析模型
<br /> 	![全级解析分析模型](./doc/full-parse_model_analysis.png)
* 代码优化
* 新增全解析功能，保留远普通解析法
* 全解析规范化处理以及模型分析
* 代码代实现

### V 0.2.4 (2017年4月5日 星期三)
* full-parse 全解析设计与实现；根据设计的方法优化程序，以及相应的测试
* 功能待实现

### V 0.2.5 (2017年4月6日 星期四)
* full-parse 全解析设计与实现： 设计模型更变，降低维度； 即改变原来的二维模型为一维模型，重构代码.
* 功能待实现

### V 0.2.6 （2017年4月8日 星期六)
* full-parse 全解析设计与实现： 代码调试~以及功能待实现 / 二级解析成功； 多级需继续实现
* 未能实现想要的结果，全解析实现失败
* ToJson 修复，适应在full-parse中新增的自定义类型(等待全局性统一)

### V 0.2.7 (2017年4月10日 星期一)
#### full-parse 全解析设计与实现： 
* 代码调试，以及功能实现(进行中)；相对应上一版本，解析过程中多级键名基本与实际相对应， 但是出现值覆盖于丢失对bug
* 现在根据最佳值命中的法则来实现值与级数匹配

### V 0.2.8 (2017年4月13日 星期四)
#### full-parse 全解析设计与实现： 
* 代码修复， 可实现简单的多级 全解析 如:
	<br/> 
	<br/> 	lily = {
	<br/> 		test = {
	<br/> 			test = {
	<br/> 				test = {
	<br/> 					test = {
	<br/> 						like = Emma
	<br/> 						toll = Emma
	<br/> 						lastname = wang
	<br/> 					}
	<br/> 				}
	<br/> 			}
	<br/> 		}
	<br/> 	}
	<br/> 
	<br/>
	"lily":{
		"test":{
			"test":{
				"test":{
					"test":{
						"toll":"Emma",
						"lastname":"wang",
						"like":"Emma"
					}
				}
			}
		}
	}
	</br>
* 新增 func (c *Conf) ValueType(dt map[string]interface{}, key string) (bool, string, interface{})  获取值类型
* vMapLinkArr full-parse 内部方法 Map-Arr 链接器设计，未测试	

### V 0.2.9 (2017年4月18日 星期二)
* 全解析解析模型改造，将原来的 VMap 改造为 vMap 将 stackMap
* 新模型实现全解析功能
* 代码优化
* 下版本将分支为 0.3.0  

### V 0.3.0 (2017年4月24日 星期一)
* 第三版本规划以及相关设计， 重构测试用例。0.2.x 版本分离
* alpha 版
* 以及其他优化

### V 1.0.0 (2017年10月28日 星期六)
* 重写规划和设计 golang-ini
* [详情](./ini-goV1.x.md)

