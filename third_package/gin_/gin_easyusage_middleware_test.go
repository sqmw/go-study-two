package gin_

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

//? Gin 的中间件，其实本质就是 装饰器模式

func TestMiddleWareUsage(t *testing.T) {
	r := gin.Default()
	r.Use(_middleWare())
	r.GET("/get", func(context *gin.Context) {
		fmt.Println("handle get request")
		context.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})
	if err := r.Run(":80"); err != nil {
		panic(err)
	}

}

// ! 定义 中间件
// \ 返回值才是中间件
func _middleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is middleWare start")
		//! 执行放行
		c.Next()
		fmt.Println("this is middleWare end")
	}
}
