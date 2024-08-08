package gin_

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//? 直接重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
