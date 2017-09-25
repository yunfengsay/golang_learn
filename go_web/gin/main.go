package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

type Login struct {
	name string `json:name`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
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
	router.Run(":9003")
}
