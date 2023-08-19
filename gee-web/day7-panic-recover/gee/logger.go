package gee

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(c *Context) {
		// start timer
		t := time.Now()
		// 处理请求
		c.Next()

		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
