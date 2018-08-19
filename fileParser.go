package inigo

// @Date：   2018/8/19 0019 10:57
// @Author:  Joshua Conero
// @Name:    文件解析器


type FileParser interface {
	Line() int
	GetData() map[interface{}]interface{}
}