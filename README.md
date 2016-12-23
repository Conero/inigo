start-20160927-go的文件Zip-UnZip工具
>> Joshua Yang
    项目来源： PHP原始ZIP压缩工具无效；例如报错
        + Fatal error: Class '命名空间\ZipArchive' not found in ...
    项目分解：命令式分布式Zip管理工具；可在服务器上运行
>>
##更新日志：
+ 2016年10月7日 星期五
    flag.String更换为os.Args包参数解析
    1.
    func (X *XHelper) command() string {
        z := flag.String("zip", "help", "输入有效的目录地址将获取到zip压缩名称") // 命名指针
        fmt.Println(os.Args)
        server := flag.String("server", "8080", "端口号")
        flag.Parse()
        _, err := os.Stat(*z)
        if err == nil { //文件存在
            return X.zip(*z)
        }
        _, err2 := os.Stat(*server)
        if err2 == nil {
            go X.myServer(*server)
        } else {
            fmt.Println(err2.Error())
        }
        fmt.Println(*server)
        return ""
    }
    2.
##版本更新记录号
+ V1.1.0=20161007
    新增:内置文件服务系统
    新增：内置普通http服务器
    新增：将flag解析命令参数改用os.Args包解析
    新增：实现zip文件压缩

    其他： 命令式多服务器并发执行
    go代码脚本  
        func.go  公共函数支持
        page,go  内置服务器器脚本
+ V1.1.5=20161002
    说明： 根据项目规划编写代码
        进展>>
            面向对象式编程支持；更变原来的方式
+ V1.0.0=20160927
    说明： 项目功能规划>>
                        定位远程服务定位；
                        远程Zip压缩控制
                        服务器化图形界面
          go代码脚本
            main.go         2016-09-27  入口文件
            XHelper.go      2016-09-27  XHelper对象/命令化数据分发            type XHelper struct 
            page.go         2016-10-07  普通服务器默认页面以及处理对象         type Page struct 
            func.go         2016-10-07  公共调用函数 
            config.go       2016-10-11  公共配置处理对象                      type Conf struct 
