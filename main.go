package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// hello world
func hello(c *gin.Context) {
	somequery := c.Query("somequery")

	if somequery == "" {
		c.JSON(http.StatusOK, gin.H{"message": "hello!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": somequery})
	}
}

// something world
func something(c *gin.Context) {
	something := c.Param("something")
	c.JSON(http.StatusOK, gin.H{"message": something})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", hello)
	router.GET("/:something", something)

	router.Run(":" + port)
}
