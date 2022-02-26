package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"time"
)

type Recipe struct {
	ID 				string		`json:"id"`
	Name        	string   	`json:"name"`
	Tags        	[]string 	`json:"tags"`
	Ingredients 	[]string 	`json:"ingredients"`
	Instructions	[]string	`json:"instructions"`
	PublishedAt		time.Time	`json:"published_at"`
}

func newRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}


func main() {
	router := gin.Default()
	router.POST("/recipes", newRecipeHandler)
	router.Run()
}

var recipes []Recipe
func init() {
	recipes = make([]Recipe, 0)
}

