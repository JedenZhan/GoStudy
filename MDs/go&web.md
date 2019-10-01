## GO + Web 编程

### 使用Http包建立Web服务

```go
package main
import (
	"fmt"
    "net/http"
    "strings"
    "log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() // 解析参数
    fmt.Println(r.form) // 输出到服务端的信息
    fmt.Println("path路径", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form("url_long"))
    for k, v := range r.Form["url_long"] {
        fmt.Println("key:", k)
        fmt.Println("value:", value)
    }
    fmt.Fprintf(w, "hello Client")
}

func main() {
    http.HandleFunc("/", sayHelloName)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("listenAndserve:", err)
    }
}
```

### Go语言http包运行机制

1. 创建Listen Socket, 监听指定端口, 等待客户端请求到来
2. Listen Socket接受客户端请求, 得到Client Socket, 接下来通过Client Socket与客户端通信
3. 处理客户端请求, 先从Client Socket读取请求头, 如果是POST方法, 还可能要读取客户端提交的数据, 然后交给相应的handler去处理, handler处理完准备好客户端需要的数据, 通过Client Socket写给客户端

### 剖析http包的内容



## 表单

不多废话,直接撸代码

html:

```html
<form action="http://127.0.0.1:9090/login">
	username: <input type="text" name="username">
	password: <input type="password" name="password">
	<input type="submit" value="登录"> 前端不解释
</form>
```

Go:

```go
package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() // 先解析参数
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form["url_long"] {
        fmt.Println("key:", k)
        fmt.Println("value:", value)
    }
    fmt.Fprintf(w, "hello Client")
} // 老代码

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("methods---:", r.methods) // 获取方法
    if r.methods == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username---:", r.Form["username"]) // post方法直接获取
        fmt.Println("password---:", r.Form["password"])
        fmt.Println("------------------华丽丽的分割线---------------------")
    }
}

func main() {
    http.HandleFunc("/", sayHelloName)
    http.HandleFunc("/login", login) // 一个路由一个处理函数
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("listenAndserve:", err)
    }
}
```

## Redis

支持五大数据类型: String, Hash, List, Set(无序), Zset(有序)



### Go 连接 Redis

安装连接库: `go get github.com/garyburd/redigo/redis`

简单的连接和存取:

```go
func main() {
    var r string
    conn, err := redis.Dial("tcp", "127.0.0.1:6379") // 成功的话返回链接实例, 失败返回err
	if err != nil {
		fmt.Println("redis.err", err)
		return // 失败的话退出
	}
    _, err = conn.Do("Set", "字段名", "字段值") // 操作存储值
    if err != nil {
        fmt.Println("操作失败了啊啊啊")
        return
    }
    r, err = redis.String(conn.Do("Get", "字段名")) // 获取值并且转化为字符串类型
    fmt.Println("连接操作成功", r) // 打印出值
}
```
