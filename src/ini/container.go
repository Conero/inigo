/* @ini-go V1.x
 * @Joshua Conero
 * @2017年11月11日 星期六
 * @ini 容器
 */
package ini

// 容器结构体
type Container struct {
	DataQueue map[string]interface{}		// 实体数据
	Child *Container						// 子节点，用于嵌套
}

// 生成容器
func MkContainer()  *Container{
	ctn := &Container{
		DataQueue: map[string]interface{}{},
	}
	return ctn
}
// 直接使用 data 生成容器
func DataToContainer(data map[string]interface{}) *Container {
	ctn := &Container{
		DataQueue: data,
	}
	return ctn
}

// 设置一级值
// 支持链式操作
func (ctn *Container) Set(key string, value interface{})  *Container{
	ctn.DataQueue[key] = value
	return ctn
}

// 获取键值
func (ctn *Container) Get(key string)(bool, interface{})  {
	value, has := ctn.DataQueue[key]
	return  has, value
}

// 删除元素
func (ctn *Container) Delete(key string) bool  {
	_, has := ctn.DataQueue[key]
	isSuccess := false
	if has{
		delete(ctn.DataQueue, key)
		isSuccess = true
	}
	return isSuccess
}
// 格式化ini文件
func (ctn *Container) ToIniString() string {
	iniString := ""
	if len(ctn.DataQueue) > 0{
		iniString = MkCreator(ctn.DataQueue).ToIniString()
	}
	return iniString
}
// 格式化json文件
func (ctn *Container) ToJsonString() string {
	iniString := ""
	if len(ctn.DataQueue) > 0{
		iniString = MkCreator(ctn.DataQueue).ToJsonString()
	}
	return iniString
}