# BBAnalyze - (BBA 分析模型)
- 2017年11月5日 星期日
- 自定义 ini 文件格式: 语法树分析处理


## BBA-struct
```go

type BBA struct {
	DataQueue         map[string]interface{}   // 数据队列值
	BranchRunningKeys []string                 // 分支运行值 孪生键值缓存
	BranchRunning     []map[string]interface{} // 分支运行值
	BRCIdx            int                      // 分支当前的索引值
	BRCkey            string                   // 分支当前的map键
	MlValue           interface{}              // 多行抽象类型
	Ini               *Ini                     // Ini 对象
}

```

> 说明
- DataQueue 顶级数据结构
- BranchRunningKeys 分支键值列表
- BranchRunning 分支键值结构体

> BranchRunning 标准结构
```ini
    ; 优化模型 - 2017年11月3日 星期五 分析与设计
    BRCIdx = -1 / >-1       键值索引，
    BRCKey = 键值对应的 map 的 key值
    mLString = string/""   多行数组值，
    mLValue = interface{}  多行抽象值
    BranchRunningKeys => []string{}     // BranchRunning 孪生匹配键值     默认： JC__SCOPE   虚拟键值，用于(无键值的)跨行数组
    BranchRunning => [
        {
            type = MAP/ARRAY/BOTH/STRING      ; 类型 map/array/both/string
            onlysa = bool         ; 仅仅是 string-array， (type = ARRAY)
            isInit = bool         ; 初始化标记
            ; value = []interface{} / map[string]interface{} 遗弃，使用两者共同代替
            array = []interface{}
            map   = map[string]interface{}
            string = 主要是实现跨行字符串解析
        }
    ]
```
