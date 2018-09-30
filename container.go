package inigo

// @Date：   2018/9/30 0030 11:04
// @Author:  Joshua Conero
// @Name:    抽象容器

type container struct {
	Data map[interface{}]interface{}
}

// data 数据检测
func (c *container) GetData() map[interface{}]interface{} {
	if c.Data == nil {
		c.Data = map[interface{}]interface{}{}
	}
	return c.Data
}

func (c *container) Get(key string) (bool, interface{}) {
	data := c.GetData()
	value, has := data[key]
	return has, value
}

func (c *container) HasKey(key string) bool {
	data := c.GetData()
	_, has := data[key]
	return has
}

func (c *container) Value(params ...interface{}) interface{} {
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

func (c *container) Set(key string, value interface{}) *container {
	c.GetData()
	c.Data[key] = value
	return c
}
