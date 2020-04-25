package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", all)
	}
}
