package inigo

// @Date：   2018/9/30 0030 11:04
// @Author:  Joshua Conero
// @Name:    抽象容器

// 参数容器
type Container struct {
	Data map[interface{}]interface{}
}

// data 数据检测
func (c *Container) GetData() map[interface{}]interface{} {
	if c.Data == nil {
		c.Data = map[interface{}]interface{}{}
	}
	return c.Data
}

// 获取值
func (c *Container) Get(key string) (bool, interface{}) {
	data := c.GetData()
	value, has := data[key]
	return has, value
}

// 获取值，且含默认值
func (c *Container) GetDef(key string, def interface{}) interface{} {
	return c.Value(key, nil, def)
}

// 是否存在键值
func (c *Container) HasKey(key string) bool {
	data := c.GetData()
	_, has := data[key]
	return has
}

// 容器值得获取/设置
func (c *Container) Value(params ...interface{}) interface{} {
	// key, nil, def
	if len(params) > 2 {
		if has, value := c.Get(params[0].(string)); has {
			return value
		}
		return params[2]
	} else if len(params) > 1 {
		c.Set(params[0].(string), params[1])
	} else if len(params) == 1 {
		if has, value := c.Get(params[0].(string)); has {
			return value
		}
	}
	return nil
}

// 设置容器参数
func (c *Container) Set(key string, value interface{}) *Container {
	c.GetData()
	c.Data[key] = value
	return c
}
