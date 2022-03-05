package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	serveMux := http.DefaultServeMux
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Week-02</title>
</head>
<body>
<a href="/header">接收客户端 request，并将 request 中带的 header 写入 response header</a><br/>
<a href="/version">读取当前系统的环境变量中的 VERSION 配置，并写入 response header</a><br/>
<a href="/ip">Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出</a><br/>
<a href="/healthz">当访问 localhost/healthz 时，应返回 200</a>
</body>
</html>
`
		fmt.Fprint(writer, html)
	})
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	serveMux.HandleFunc("/header", func(writer http.ResponseWriter, request *http.Request) {
		header := request.Header
		for key, values := range header {
			for _, value := range values {
				writer.Header().Add(key, value)
				fmt.Printf("%s:%s\n", key, value)
			}
		}
	})
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	serveMux.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		version, ok := os.LookupEnv("VERSION")
		if !ok {
			version = "v1.0.1"
			os.Setenv("VERSION", version)
		}
		fmt.Fprintf(writer, "env VERSION: %s ", version)
	})

	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	serveMux.HandleFunc("/ip", func(writer http.ResponseWriter, request *http.Request) {
		forwardedFor := request.Header.Get("x-forwarded-for")
		if forwardedFor == "" {
			forwardedFor = request.Header.Get("X-Real-IP")
		}
		if forwardedFor == "" {
			forwardedFor = request.RemoteAddr
		}
		fmt.Fprintf(writer, "客户端 IP: %s ", forwardedFor)

	})

	// 当访问 localhost/healthz 时，应返回 200
	serveMux.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		fmt.Fprint(writer, "200")
	})
	err := http.ListenAndServe(":80", serveMux)
	if err != nil {
		log.Fatalln(err)
	}
}
