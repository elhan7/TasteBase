package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Instruction string   `json:"instruction"`
}

var recipes = map[int]Recipe{
	1: {
		ID:          1,
		Title:       "Легкая паста",
		Ingredients: []string{"Макароны", "Сыр", "Масло"},
		Instruction: "Сварить макароны, добавить масло и посыпать тертым сыром.",
	},
}

func main() {
	r := gin.Default()

	r.GET("/recipes", func(c *gin.Context) {
		var list []Recipe
		for _, v := range recipes {
			list = append(list, v)
		}
		c.JSON(http.StatusOK, list)
	})

	r.POST("/recipes", func(c *gin.Context) {
		var newRecipe Recipe
		if err := c.ShouldBindJSON(&newRecipe); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newRecipe.ID = len(recipes) + 1
		recipes[newRecipe.ID] = newRecipe
		c.JSON(http.StatusCreated, newRecipe)
	})

	r.GET("/recipes/random", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		keys := make([]int, 0, len(recipes))
		for k := range recipes {
			keys = append(keys, k)
		}
		randomID := keys[rand.Intn(len(keys))]
		c.JSON(http.StatusOK, recipes[randomID])
	})

	r.Run(":8080")
}