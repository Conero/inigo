# ini-go Version 1.x
- 2017年10月28日 星期六
- Joshua Conero

## 包目录
- ini
    - parse.go          文件解析文件
    - variable.go       包变量
## 分支
- rong      简单(ini-文件解析器)， 符合标准 ini 文件(命名含义-纪念作者的第一个女朋友SuRong@2017)
    - 简单化
    - 不支持嵌套
    - 支持 ini section 节
- ini-go    支持多重ini-文件阅读器， ini 文件扩展
    - 在 rong 的基础上进行扩展
    - 基础 ini-go 0.x 版本的语法风格
    - 可实现值嵌套

## 结构体
```go

// ini 结构体
type Ini struct {
	FileName string
}

```    
- Ini 构造函数
    - Open(name string) *Ini

- Ini 基本方法
    - 