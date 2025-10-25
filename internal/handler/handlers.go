package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pablopasquim/GoPizza/internal/models"
	"github.com/pablopasquim/GoPizza/internal/data"
)

func GetPizza(c *gin.Context) { // Handler function to get pizzas
	pizzas := data.Pizzas
	c.JSON(200, gin.H{ // Return the list of pizzas
		"pizzas": pizzas,
	})
}

func GetPizzaID(c *gin.Context) {
	pizzaID := c.Param("id") // Get pizza ID from URL parameter
	id, err := strconv.Atoi(pizzaID) // Convert string ID to integer
	if err != nil { // Handle conversion error
		c.JSON(400, gin.H{ // Return bad request response
			"error": "Invalid pizza ID format",
		})
		return
	}

	for _, pizza := range data.Pizzas { // Search for pizza by ID
		if pizza.ID == id { // If found
			c.JSON(200, gin.H{ // Return the found pizza
				"pizza": pizza,
			})
			return
		}
	}

	c.JSON(404, gin.H{ // Return not found response
		"error": "Pizza not found",
	})
}

func PostPizza(c *gin.Context) {
	var newPizza models.Pizza // Declare a variable to hold the new pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil { // Bind JSON to struct
		c.JSON(400, gin.H{
			"error": err.Error(),
			}) // Handle binding error
		return
	}

	c.JSON(201, gin.H{ // Return success response
		"message": "Pizza created successfully", // Success message
		"pizza": newPizza,
		})// Return the created pizza

		data.Pizzas = append(data.Pizzas, newPizza)
}

func PutPizza(c *gin.Context) {
	pizzaID := c.Param("id") // Get pizza ID from URL parameter
	id, err := strconv.Atoi(pizzaID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid pizza ID format",
		})
		return
	}

	var updatedPizza models.Pizza // Declare a variable to hold the updated pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil { // Bind JSON to struct
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	updatedPizza.ID = id // Ensure the ID remains unchanged
	c.JSON(200, gin.H{ // Return success response
		"message": "Pizza updated successfully", // Success message
		"pizza": updatedPizza,
		}) // Return the updated pizza
}

func DeletePizza(c *gin.Context) {
	pizzaID := c.Param("id") // Get pizza ID from URL parameter
	id, err := strconv.Atoi(pizzaID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid pizza ID format",
		})
		return
	}
	c.JSON(200, gin.H{ // Return success response
		"message": "Pizza deleted successfully", // Success message
		"id":  id,
		}) // Return the deleted pizza ID
}
