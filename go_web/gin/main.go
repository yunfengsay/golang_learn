package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

type Login struct {
	name string `json:name`
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/login", func(c *gin.Context) {
		// c.JSON(200, Login{"yunfeng"})
		c.JSON(http.StatusOK, gin.H{
			"name": "yunfeng",
			"age":  "26",
			"block_hole": gin.H{
				"sexed": "none",
			},
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	// router.POST("/person", AddPersonApi)

	// router.GET("/persons", GetPersonsApi)

	// router.GET("/person/:id", GetPersonApi)

	// router.PUT("/person/:id", ModPersonApi)

	// router.DELETE("/person/:id", DelPersonApi)

	return router
}

func main() {
	router := gin.Default()
	initRouter()
	router.Run(":9003")
}
