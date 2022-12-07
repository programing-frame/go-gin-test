package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/test", func(c *gin.Context) {
		// 查询请求URL后面拼接的参数
		id := c.Query("id")
		//查询请求URL后面的参数，如果没有填写默认值
		page := c.DefaultQuery("page", "0")
		// 从表单中查询参数
		name := c.PostForm("name")
		//从取得URL中参数，此处URL中没有name字段
		name = c.Param("name")
		// 从表单中查询参数,，如果没有就获取默认值
		message := c.DefaultPostForm("message", "default")
		// 获取Body值
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		// 获取header指定字段
		headerTest := c.Request.Header.Get("test")
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		fmt.Printf("id: %s; page: %s; name: %s; message: %s; header_test:%s; bodyBytes:%s", id, page, name, message, headerTest, bodyBytes)

		data := map[string]interface{}{
			"data": "ok",
			"code": 200,
		}
		c.JSON(200, data)

	})
	router.Run(":8080")
}
