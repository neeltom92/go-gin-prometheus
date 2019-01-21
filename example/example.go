package main

import (
	"github.com/banzaicloud/go-gin-prometheus"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin", nil)

	p.Use(r, "/metrics")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}
