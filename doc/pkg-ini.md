# ini / pkg - (20171028)


## CHANGELOG - 更新日志

### v1.1.0 / 20171111
- (优化) parse.go
    - func (I *Ini) GetString(key string) string   方法支持多级“.”键值访问对象，利用 func (I *Ini) Get(key string) (bool, interface{}) 原型 @20171110
    - 支持在有效行后注释文本如： *** key = value ;或# 注释文本 *** 
- (优化) 文件格式生成器
    - renamed:    pkg/ini/json.go -> pkg/ini/creator.go
    - 将原来的函数，利用新增的 *** Creator *** 结构体实现
    - json 生成器支持多级格式化字符串， 删除原来的 json 字符中的转义处理
    - (新增) map -> ini 文件格式化处理器， 支持多级嵌套
- (新增)  Container.go    新增数据结构抽象化
    - 关联文件生成器“ini”/“json”文件格式解析器
    - 关联ini文件解析器，用于对解析的数据进行处理以及更新操作
    - 提供简单的增删改查操作    
- 其他
    - 本次更新标记为大版本更新： 
        - v1.1.x:   + ini 文件生成器实现， 解析器优化
        - v1.0.x:   + ini 文件解析实现

### v1.0.7 / 20171107
- (优化) BBA
    - 删除 BBA中的 BRCKey 属性， 直接使用 BranchRuning 机制实现
    - 新增 获取当前指向的 键值(原BRCKey) 的方法
        - func (bba *BBA) getCurrentKeyMap() (string, map[string]interface{}) 
    - 尝试使用 BranchRuning 机制实现***字符串***跨行解析
    - ***type=BOTH***优化，使其更加符号实际逻辑， 也即是不打乱： map/string 中的顺序
- (优化) parse.go
    - 修复文件最后一行无法遍历的原因(可寻找更好的方法优化)
    - func (I *Ini) Get(key string) (bool, interface{}) 
        - 实现“.”操作多级解析文件
- 其他
    - 新增 php/ini 算法跨速测试文本
    - 当前版本基本实现了ini文件的解析(可以用用其他的项目中)
        - 字符串跨行解析
        - 多级 map 遍历
        - 字符串跨行解析
    - 计划实现
        - 二维数据解析
        - array 中有 map/以及string 类型的解析

### v1.0.6 / 20171104
- (新增)
    - 支持***字符串跨行***， 可用于 []string 内的值
    - BBAnalyze.go master 分支内暂时使用可关闭调试器

- (优化) BBA
    - 删除 BBAnalyze.go  原来的分析模型，用新的方法替代
    - variable.go 中添加与模型解析相关的***正则表达式***
    - 暂时引入 ***OnlyTestPutOut***常用，仅仅用在开发脚本中(测试条件下)， 正式版本将删除
    - liner.go 添加与字符串跨行先关的方法

- (优化) BBAnalyze 解析过程中，多级 map 出错的的问题    @1.0.5 版本测试时出现
    - golang 中切片子切片时，出错
```go
    // 8bit
    test := []int{1, 2, 5, 7, 6, 8, 9, 5}
    fmt.Println(test)

    // 取子切片
    bit := 2            // 为子字符串的长度，而非偏移量()
    test = test[0: bit]
```

### v1.0.5 / 20171104
- (优化) BBA
    - BBAnalyze.go 内部引入用于调试的格式文件输出
    - BBA 模型更新，根据***@v1.0.4***版本规划
    - 经过测试时，二级/多级 map 与设计(预想)结果不一致 -- 暂时无法察觉到问题所在

### v1.0.4 / 20171103
- (新增) pkg/running 包 用于程序运行时的时间统计(公共)
- (优化) BBA
    - 模型重构，用于适应更加复杂的 ini 结构文件
        - 数组中包含 map， 以及多重类型 (以前的设计构思将无法适应)
        - map 中的复杂类型的
    - 当然通常情况下复杂的类型比较少，因此在考虑的设计的时候不能过去倚重特例(！！)
    - 支持多行字符串解析
    - 支持二位数组(复杂)解析

### v1.0.3 / 20171102
- (优化) BBA 
    - func (bba *BBA) MiltiLineToArray(value string) *BBA
    - 添加 ValueKey/ValueArray 属性用于实现 数组分行写法的解析
```ini
    key = {
        key = value
        test = {
            v1
            v2
            v3
            v5,
            v8,
            v10
        }
    }
```    
- (优化) json.go
    - func jsonTransform(sJson string, isReconvert bool) string    逻辑修复

### v1.0.2 / 20171101
- (新增) Ini.DataQueue 转化为 json 字符串方法
    - 新增文件 json.go
    - 创建全局 func ToJsonStr(queue map[string]interface{}) string  函数
    - Ini 函数 func (I *Ini) ToJsonString() string 引用 全局函数
- (新增) Liner
    - func (line *Liner) isMultiComment(sLine string) bool 添加多行注释处理    
- (优化) BBA 优化
    - BBA 实现简单的 ini -> map 的数据处理(单级)
    - CBaseKeys -> CBaseValues  一一对应
    - 基本使用流程
        - func BBAnalyze() *BBA  创建 BBA 对象
        - 行遍历处理
            - func (bba *BBA) UpdateBaseKey(bKey string) 更新基键
            - func (bba *BBA) PushQueue(key string, value interface{}) *BBA 推送值
            - func (bba *BBA) CommitQueue() bool    删除基键
        - BBA.DataQueue 获取模型处理的结果集
- func (line Liner) singleObject(cLine string) interface 当行对象解析实现
- 与 v0.x 版本比较
    - BBAnalyze.go 接管行解析模型处理   
    - 值更加注重原始样式，不在自动消除字符充的空格
    - 支持解析的功能更加全面


### v1.0.1 / 20171031
- 引入“基枝模型”分析方法解析多重 ini 结构
    - (新增) BBAnalyze.go  数据处理引擎
- (新增) liner.go 用于处理字符，而改善 v0.x版本中字符解析仅仅集中于某一方法
- (新增) test.go 包内测试应用，可测试包内私有方法

### v1.0.0 / 20171028
- ini 包 基本架构
- 在v0.x的基础上快速构建 go-ini 解析库
- 实现基本的 ini 文本解析
- 基本保持与***rong-ini***代码一致