package main

import (
	f "fmt"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数
	f.Println(r.Form)
	f.Println("path", r.URL.Path)
	f.Println("scheme", r.URL.Scheme)

	for k, v := range r.Form {
		f.Println("key", k, "value", v)
	}
	f.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		f.Println("err", err)
	}
}
