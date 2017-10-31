# common 库版本说明
- 2017年3月30日 星期四
- Joshua Conero
- version go1.8 *

## 风格
- 函数命名风格  - 骆驼峰命名法
- 私有类小写打头， 公用大写打头(golang 语法规定)

# version - tree
- 三级增版本计数法  
- 0.0.1 - 170330 (版本号 - 日期)
- a.b.c  => a. 大版本更替(具有重大意义) ; b. 小版本更替(大版本下分支更新) ; c. a.b. 同提交下提交递增好; 

## 0.0.1 - 170330
- 包含脚本
    array.go                数组/切片处理类
    common.go               常用函数
    variable.go             公共库私有脚本
    ~ project.go            特定项目脚本下常量