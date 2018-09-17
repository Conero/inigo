# ChangeLog(v2)

> Joshua Conero
>
> 20180819



## v2.0.x

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

