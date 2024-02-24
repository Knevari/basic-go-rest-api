package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type game struct {
	Id int `json:"id"`
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

func getGameById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "invalid id" })
	}
	
	for _, game := range games {
		if game.Id == id {
			c.IndentedJSON(http.StatusOK, game)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "game not found" })
}

func main() {
	router := gin.Default()

	router.GET("/games", getGames)
	router.GET("/games/:id", getGameById)

	router.Run("localhost:8080")
}