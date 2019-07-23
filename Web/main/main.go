package main

import (
	"fmt"
	// "net/http"
	"time"
)

// func sayHelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() // 解析参数
// 	f.Println(r.Form)
// 	f.Println("path", r.URL.Path)
// 	f.Println("scheme", r.URL.Scheme)

// 	for k, v := range r.Form {
// 		f.Println("key", k, "value", v)
// 	}
// 	f.Fprintf(w, "{name: \"JedenZhan\"}")
// }

// func main() {
// 	http.HandleFunc("/", sayHelloName)
// 	err := http.ListenAndServe(":9090", nil)
// 	if err != nil {
// 		f.Println("err", err)
// 	}
// }

// type MyMux struct {
// }

// func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		sayHelloName(w, r)
// 		return
// 	}
// 	http.NotFound(w, r)
// 	return
// }

// func sayHelloName(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello MyRoute")
// }

// func main() {
// 	mux := &MyMux{}
// 	http.ListenAndServe(":9090", mux)
// }

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("%d ", i)
	}
}

func main() {
	go numbers()

}
