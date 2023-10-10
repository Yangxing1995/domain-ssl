// main.go (updated)

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 设置静态文件夹，用于存放 HTML 文件
	r.Static("/static", "./static")

	// 创建一个路由处理程序，用于加载 HTML 页面
	r.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "static/index.html", gin.H{
		//	"title": "Hello, Go with HTML",
		//})
		data, err := ioutil.ReadFile(filepath.Join("static/", "index.html"))
		if err != nil {
			c.AbortWithError(500, err)
		} else {
			c.Data(200, "", data)
		}
	})

	r.POST("/check", func(c *gin.Context) {
		// 获取前端发送的域名和token参数
		domain := c.PostForm("domain")
		token := c.PostForm("token")

		if domain == "" {

			c.String(http.StatusBadRequest, "domain is empty")
			return
		}

		// 使用curl命令获取数据
		cmd := exec.Command("curl", "-H", fmt.Sprintf("X-CMC-DCV-Host:%s", domain), fmt.Sprintf("https://file.httpsauto.com/.well-known/pki-validation/%s.txt", token))
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.String(http.StatusBadRequest, "check failed")
			return
		}

		c.String(http.StatusOK, string(output))
	})

	r.Run(":8080")

}
