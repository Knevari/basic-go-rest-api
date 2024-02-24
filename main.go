package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type game struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Rating float32 `json:"rating"`
}

func index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{ "message": "Hello from Go" })
}

func main() {
	router := gin.Default()

	router.GET("/", index)

	router.Run("localhost:8080")
}