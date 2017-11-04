# go ini 文件解析库
> - @author Joshua Conero
> - @descrip ini 文件解析器


## 分支
- v0.x 版本
	- [详情](./doc/readme-v0.x.md)
	- 开发周期： @20170119 - 20170424
- v1.x 版本		(开发中)
	- 开始： 20171028 -> 

## v1.x (20171028 - )
> ***go 开发环境：***

- go@1.9.2
- gogland


***目录说明***
- pkg		程序包脚本
	- ini	用于复杂 ini 文件结构解析， 支持map/array ， 以及map/array复合解析。自定义 ini 脚本语法，采用结构树分析法
		- @20171028
		- [详情](./doc/pkg-ini.md)
	- rong  用于简单 ini 文件解析，支持 ini 文件标准
		- @20171028
		- [详情](./doc/pkg-rong.md)
	- running 工具器， 时间运行跟踪
		- @20171103
		- [详情](./pkg/running/README.md)
- bin 		包测试以及调试

### pkg/ini
> - 语法实例
> ***注释***
```ini
; 单行注释

'''
	支持多行注释
	JOSHUA CONERO
	与 python 相似
'''

```


> ***字符串***
```ini

basestring = 基本字符串，支持转移符号如： \= \, 等
nl2br = "
	支持跨行字符
"
nl2br2 = '
	支持跨行字符
'

```

> ***数组***
```ini
; 数组类型 array
array = v1, v2, v3, v4
array2 = {
	v1
	v2
	v3
	v4
}
```

> ***map***
```ini
; map
scope = {
	; 当行注释
	key1 = value1
	key1 = value1
	scopeinner = {
		descrip = 内部嵌套结构
	}
}
```

> ***array/map***
```ini
; array/map
am = {
	map = {
		k1 = v2
		k2 = v2
		k3 = v2
	}
	array = {
		test
		test
		map = tests
	}
}

```

### pkg/rong
> 基本语法
```ini

base = 基本字符串测试

; 支持节
[section]
base = 字符ini标准
other = yu pkg/ini 比较更加轻量级，效率更高。用于简单的 ini 文件解析


```


## v0.x (20170119 - 20170424)
>
- [详情](./doc/readme-v0.x.md)

***目录说明***
- src/
	- 包脚本
- dist/
	- 发布版本