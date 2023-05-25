package main

import (
	"log"
	"net/http"
)

// 我们创建了一个处理器链，通过 http.HandlerFunc 包装 helloHandler 函数，然后将它作为参数传递给中间件函数 loggingMiddleware 和 apiKeyMiddleware，最终生成一个链式处理器。在这个链式处理器中，loggingMiddleware 函数用于记录请求信息，apiKeyMiddleware 函数用于检查 API 密钥，helloHandler 函数用于处理请求

// 定义一个记录请求信息的中间件
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// 定义一个检查 API 密钥的中间件
func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "my-secret-key" {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// 定义一个处理请求的处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		return
	}
}

func main() {
	// 创建一个处理器链
	handler := apiKeyMiddleware(loggingMiddleware(http.HandlerFunc(helloHandler)))

	// 启动 HTTP 服务器
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic("Server Error")
	}
}
