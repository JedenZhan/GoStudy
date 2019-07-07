package main

import (
	"fmt"
	"log"
	_ "math/rand"
	"net/http"
	"text/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 先解析参数
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form["url_long"] {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
	fmt.Fprintf(w, "hello Client")
} // 老代码

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methods---:", r.Method) // 获取方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("test.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()                                   // 这个是必须的,否则不能获取到数据
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
