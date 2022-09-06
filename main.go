package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type example struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var examples = []example{
	{ID: 1, Title: "Example 1"},
	{ID: 2, Title: "Example 2"},
	{ID: 3, Title: "Example 3"},
}

func getExamples(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, examples)
}

func main() {
	router := gin.Default()

	router.GET("/", getExamples)
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
