package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href='/home'> CLick here to start a request`))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	span := opentracing.StartSpan("/home")
	defer span.Finish()

	w.Write([]byte("Request Start<br/>"))
	go func() {
		http.Get("http://localhost:8080/async")
	}()
	_, err := http.Get("http://localhost:8080/service")
	if err != nil {
		ext.Error.Set(span, true)
		// span.LogEventWithPayLoad("GET service error", err)
	}
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // 模拟超时
	w.Write([]byte("Request Done"))
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	http.Get("http://localhost:8080/db")
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
}

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/async", serviceHandler)
	mux.HandleFunc("/db", dbHandler)
	mux.HandleFunc("/service", serviceHandler)
	fmt.Printf("go to http://localhost:%d/home to start a request", port)
	log.Fatal(http.ListenAndServe(addr, mux))

}
