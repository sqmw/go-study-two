package gin_

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGet(t *testing.T) {
	r := gin.Default()
	//? *XXX *这个通配参数只能放在最后一个/之后
	//? 这里将会拦截所有 /a/XXX 这种请求，/*aaa 不做强制要求
	r.GET("/a/:user-info/*aaa", func(c *gin.Context) {
		aa := c.Param("user-info")
		aaa := c.Param("aaa")
		//! 获取不到的，因为没有
		//aaaa := c.Param("aaaa")
		queryP1 := c.Query("aaaa")
		c.JSON(200, gin.H{
			"aa":      aa,
			"aaa":     aaa,
			"queryP1": queryP1,
		})
	})
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}

func TestPost(t *testing.T) {
	r := gin.Default()
	r.POST("/post/:user-info/*aaa", func(c *gin.Context) {
		reqData := make(map[string]string)

		if err := c.BindJSON(&reqData); err != nil {
			panic(err)
		}

		c.JSON(200, reqData)
	})
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
