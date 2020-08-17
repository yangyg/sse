package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.GET("/sse", sse)

	r.LoadHTMLGlob("index.html")

	r.Run(":6788") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func sse(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream") // 规定把报头设置为 text/event-stream
	c.Header("Cache-Control", "no-cache")         // 设置不对页面进行缓存

	timer := time.NewTimer(2 * time.Second)
	c.Stream(func(w io.Writer) bool {
		for {
			select {
			case <-timer.C:
				msg := "msg"
				c.SSEvent("message", msg)
				timer.Reset(2 * time.Second) // 上一个执行完毕重新设置
				return true
			}
		}
	})
}
