# ChangeLog

> Joshua Conero
>
> 2019年6月6日 星期四

## 0.2

### 0.2.0/20190612-alpha

- ini
  - (+) *oIRQueueManger 新增方法 `getCurIni` 和 `getNameList`*
  - (实现) `open`
    - (+) 添加可选参数 `--alias=<别名>`
    - (实现) *不同目录相同文件名可读取 (限制取消)、项目文件可通过 `alias` 加载*



> **todo**

- ini
  - open 命令增强
    - 不同目录相同文件名可读取 (限制取消)
    - 项目文件可通过 `alias` 加载
    - 读取文件时，获取运行信息：*耗时，行数、文件大小、注释行等*
  - about 当前加载资源的信息展示(新增)





## 0.1

### 0.1.0/20190606

- 命令行项目程序搭建
- (+) 添加命令 `$ ini`
  - 实现的交互时的内部命令，如: *open, use, list, get , help*
- (+) *添加系统所需要的默认配置文件 `inigo.ini`*

