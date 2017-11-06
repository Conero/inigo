# pgk/ini 语法
- Joshua Coenro
- 2017年11月6日 星期一

## 概述
- 解析属性选项
    - $Conero.INI__AutoInferType        自动类型推断
        - 简写：  $C.INI__AIT
        - 自动类型推断开启， 文件头部设置
        - 经过正则匹配指定的类型
        - 设计/构想日期   @2017年11月6日 星期一
        - 实现版本号： (进行中……)
            - 7          -> int
            - 17.012     -> float
            - 其他        -> string
            - true/false  -> bool                   大小写不敏感

## ini 文本编译为二进制
- 版本设计与构想            