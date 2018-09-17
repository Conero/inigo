# inigo(go ini 文件解析库)
> - @author Joshua Conero
> - @descrip ini 文件解析器

## 项目管理
- ``master`` 主分支；用户可下载使用
- ``alpha`` 开发数据分支(develop)，程序开发不直接操作``master`` 而由开发该分组再合并到主分支
- ``demo`` 项目实际测试；
- ``document`` 项目文档
- ``v{n}`` 历史版本分支，历史保存

## 设计

- `Parser`		解析器**接口**
  - `BaseParser`   默认*ini* 文件解析器
  - `RongParser`   *rong* ini 文件解析器
- `FileParser` 文件解析器**接口**
- `StrParser` 字符串解析器**接口**



## 分支

- v0.x 版本
	- [详情](./doc/readme-v0.x.md)
	- 开发周期： @20170119 - 20170424
		v1.x 版本		(开发中)
	- 开始： 20171028 -> 
		v2.x (版本)	
    - 通过对 go 语言的学习重新库；v1.x中项目设计多数受其他语言的影响，完全按照go语言的风格。

## v2.x (20180819 - )
### 特性

- 使用新的 *git* 管理方式；见 ``项目管理``
- 程序测试使用go语言提供的 *test* 测试程序
- 删除项目中与库无关的文件夹，转移至分支



> ***go 开发环境：***

- go@1.10
- gogland
