package main

import (
	"net/http"

	"./models"
	"gopkg.in/gin-gonic/gin.v1"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api/login", func(c *gin.Context) {
		// c.JSON(200, Login{"yunfeng"})
	})
	router.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "hello")
	})
	router.POST("/api", func(c *gin.Context) {

	})
	return router
}

func main() {
	models.Run()
	router := gin.Default()
	initRouter()
	router.Run(":9003")
}
