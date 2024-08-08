package gin_

import (
	"fmt"
	"net/http"
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
	err := r.Run(":80")
	if err != nil {
		fmt.Println("run error", err)
		return
	}
	fmt.Println("A")
}

func TestA2(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello")
	})
	err := r.Run(":80")
	if err != nil {
		return
	} else {
		return
	}
}
