package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main () {

	http.HandleFunc("/", sayHelloName)	 // 设置访问的路由
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("err", err)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数
	fmt.Println(r.Form) // form表单
	fmt.Println("path", r.URL.Path)
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("value",strings.Join(v, ""))
	}

	fmt.Fprintf(w, "hello astaxie !") // Fprintf？
}
