package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 测试请求
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// rest风格请求
	router.GET("/rest/:id", func(c *gin.Context) {
		id, flag := c.Params.Get("id")
		if !flag {
			id = "未找到"
		}
		c.JSON(200, gin.H{
			"message": id,
		})
	})

	// get请求
	router.GET("/get", func(c *gin.Context) {
		// c.Query("name")：c.Request.URL.Query().Get("name") 的简写
		name := c.Query("name")
		age := c.DefaultQuery("age", "18")
		c.JSON(200, gin.H{
			"message": name + " " + age,
		})
	})

	// form_post请求
	router.POST("/form_post", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":   "posted",
			"username": username,
			"password": password,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
