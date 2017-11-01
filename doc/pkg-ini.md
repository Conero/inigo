# ini / pkg - (20171028)


## CHANGELOG - 更新日志

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