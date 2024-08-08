package gin_

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGroup(*testing.T) {
	r := gin.Default()
	v1 := r.Group("/v1")
	v2 := r.Group("/v2")

	//! v1 的代码
	{
		v1.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			fmt.Print("name from v1", name)
			c.JSON(200, gin.H{
				"name":  name,
				"birth": "v1",
			})
		})
	}

	//! v2 的代码
	{
		v2.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			fmt.Print("name from v2", name)
			c.JSON(200, gin.H{
				"name":  name,
				"birth": "v2",
			})
		})
	}
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
