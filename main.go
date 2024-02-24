package main

import (
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type game struct {
	Id int `json:"id"`
	Title string `json:"title" binding:"required,max=128,min=2"`
	Genre string `json:"genre" binding:"required"`
	Rating float32 `json:"rating" binding:"required,min=1,max=5"`
}

type ErrorMsg struct {
	Field string `json:"field"`
	Message   string `json:"message"`
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

func postGame(c *gin.Context) {
	var newGame game;

	if err := c.BindJSON(&newGame); err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "errors": out })
		}

		return
	}

	newGame.Id = len(games) + 1
	
	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, games)
}

func main() {
	router := gin.Default()

	router.GET("/games", getGames)
	router.GET("/games/:id", getGameById)
	router.POST("/games", postGame)

	router.Run("localhost:8080")
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Should be greater than " + fe.Param()
	case "max":
		return "Should be less than " + fe.Param()
	}

	return "Unkown error"
}