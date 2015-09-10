package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

// Main function, load config.ini and run a http web server.
func main() {

	// 使用pat来做类似sinatra的路由
	router := pat.New()

	// 用于监控服务器状态的入口
	router.Get("/status", http.HandlerFunc(TestStatus))
	// 用于网页URL测试
	router.Get("/url", http.HandlerFunc(MobileUrl))

	// 注册为默认路由
	http.Handle("/", router)
	// HTTP Server
	log.Println("Web Tools Server start on port 8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
