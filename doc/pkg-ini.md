# ini / pkg - (20171028)


## CHANGELOG - 更新日志

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