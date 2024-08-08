package gin_

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestA(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200,
			gin.H{
				"hello": "world",
			})
	})
	r.GET("/exit", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"exit": "ok",
		})
		os.Exit(1)
	})
	err := r.Run()
	if err != nil {
		fmt.Println("run error", err)
		return
	}
	fmt.Println("A")
}
