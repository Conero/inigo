# ChangeLog(v2)

> Joshua Conero
>
> 20180819



## v2.0

### v2.0.10/180930

> 设计调整

  ```
  设计: 
  	1. BaseParser -> container		(继承)
  				  -> Parser			(实现)
  				  
  	2. RongParser -> BaseParser     (继承)
  ```

> **package**


- (调整) 将旧版中 *LnReader* 移到新版中

- (优化) *NewParser* 函数的重写

- (+) *新增`container` 抽象容器·，实现对容器中数据的获取以及设置*

- `Parser`


  - (+) *添加方法Section，用于获取有关section的参数*

- `BaseParser`

  - (+) *实现对基本ini文件语法的支持，完成对文件的解析，并且获取到数据*

- `baseFileParse`


  - (+) 添加 *base-ini* 文件的读取与解析




### v2.0.1/180819

- **概述**
  - 删除历史旧代码
  - 不在区分子包， *全部包含在项目下*
- **package**
  - (+) 添加 *parser.go* 设计 ``Parser`` 解析器接口
    - (+) 实现 ``BaseParser`` 基本 ini 文件解析器
  - (+) 添加 StrParser接口，用于字符串的简单解析
  - (+) 搭建``RongParser`` 解析器

### v2.0.0/180819

- 重设计项目架构，优化 git 管理工具
- 删除``v1`` 中于*package* 无关的代码以及文档，由 *ini-go* 更名为 *inigo*
- 保证项目无错误
