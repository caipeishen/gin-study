package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 测试请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// rest风格请求
	r.GET("/rest/:id", func(c *gin.Context) {
		id, flag := c.Params.Get("id")
		if !flag {
			id = "未找到"
		}
		c.JSON(200, gin.H{
			"message": id,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
