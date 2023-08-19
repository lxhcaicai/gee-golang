package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandleFunc {
	return func(context *gee.Context) {
		t := time.Now()
		// 如果服务出现错误
		context.Fail(500, "Internal Server Error")
		// 计算处理的时间
		log.Printf("[%d] %s in %v for group v2", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(context *gee.Context) {
			context.String(http.StatusOK, "hello %s, you're at %s\n", context.Param("name"), context.Path)
		})
	}
	r.Run(":9999")
}
