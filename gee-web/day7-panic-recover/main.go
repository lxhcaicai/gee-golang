package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(context *gee.Context) {
		context.String(http.StatusOK, "Hello Gee\n")
	})

	//索引超出测试范围
	r.GET("/panic", func(context *gee.Context) {
		names := []string{"lxhcaicai"}
		context.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
