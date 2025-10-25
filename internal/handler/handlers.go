package handlers
import (
	"github.com/gin-gonic/gin"

	"github.com/pablopasquim/GoPizza/internal/models"

)

func GetPizza(c *gin.Context) { // Handler function to get pizzas
	pizzas := []models.Pizza{ // Sample data
		{ID: 1, Name: "Margherita", Price: 10.99},
		{ID: 2, Name: "Pepperoni", Price: 12.99},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas, // Return the list of pizzas
	})
}
