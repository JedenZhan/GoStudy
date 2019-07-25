package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

// func checkErr(err) {
// 	if err != nil {
// 		return "出错了 %r"err
// 	}
// }

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

}

func reg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formData := r.Form
	userName := formData["username"]
	password := formData["password"]
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("连接失败....")
	}
	_, err = conn.Do("Set", "userName", userName)
	_, err = conn.Do("Set", "passWord", password)

	fmt.Println(formData["username"], formData["password"])
	if err != nil {
		fmt.Println("")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// formData := r.Form

}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/reg", reg)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}
}
