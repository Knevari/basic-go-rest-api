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

var games = []game {
	{ Id: 1, Title: "Elden Ring", Genre: "Action RPG", Rating: 5  },
	{ Id: 2, Title: "God of War: Ragnarok", Genre: "Adventure", Rating: 5  },
	{ Id: 3, Title: "Little Nighmares", Genre: "Horror", Rating: 5  },
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games);
}

func main() {
	router := gin.Default()

	router.GET("/games", getGames)

	router.Run("localhost:8080")
}