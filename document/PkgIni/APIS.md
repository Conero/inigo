# APIS 系统接口
- 2017年11月5日 星期日

## APIs

## type Ini

// ini 结构体
type Ini struct {
        Liner                            // 组合继承- 行处理
        FileName  string                 // ini 文件
        DataQueue map[string]interface{} // ini 解析后的数据
        IsSuccess bool                   // ini 文件是否解析成功
        FailMsg   string                 // 错误信息
        File      *File                  // 文件解析信息
}

- func Open(name string) *Ini
- func (I *Ini) Get(key string) (bool, interface{})
- func (I *Ini) GetString(key string) string
- func (I *Ini) HasKey(key string) bool
- func (I *Ini) ToJsonString() string