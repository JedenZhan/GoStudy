关于Vscode报错的解决(mac or linux)

打开vscode_config, 复制三个脚本到$GOPATH/src下, 先增加执行权限(chmod +r ./*.sh)

```shell
bash install_plugin.sh
```


## GO 语法小结

- var和const进行变量和常量的声明
- package和import为包管理和导入包
- func进行函数定义
- return用于函数返回值,多个返回值用逗号隔开,并且要声明返回类型
- defer用于函数结束前执行的最后一行代码
- go用于并行
- select用于选择不同的通讯
- interface定义接口
- struct定义抽象数据类型
- 流程控制
- chan用于channel通讯
- map声明map类型数据
- range读取slice, map, channel数据



## Menu

- [基本(数据类型等)](./Mds/base.md)
- [函数(go一等公民)](./MDs/function.md)
- [协程和管道(goroutine&channel)](./MDs/goroutine&channel.md)
- [go+web编程](./MDs/go&web.md)























 