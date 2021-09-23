package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

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

	// 上传文件
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		// 上传文件到指定的路径
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// 多上传文件
	router.POST("/uploadMulti", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的路径
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
