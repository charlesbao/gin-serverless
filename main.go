package main

import (
	"gin-serverless/gateway"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(os.Getenv("mode"))
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Header": c.Request.Header,
			"Query":  c.Request.URL.Query(),
			"IP":     c.ClientIP(),
			"URI":    c.Request.RequestURI,
		})
	})

	if gin.Mode() == "release" {
		gateway.Serve(r)
	} else {
		r.Run(":8080")
	}

}
