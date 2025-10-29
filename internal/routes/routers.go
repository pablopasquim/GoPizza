package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pablopasquim/GoPizza/internal/handlers"
)

func HandleRequest(r *gin.Engine) { // Function to handle routes
	r.GET("/pizzas", handlers.GetPizzas) // Route to get pizzas
	r.GET("/pizzas/:id", handlers.GetPizzasByID) // Route to get pizza by ID
	r.POST("/pizzas", handlers.PostPizzas) // Route to create a new pizza
	r.PUT("/pizzas/:id", handlers.UpdatePizzaByID) // Route to update a pizza
	r.DELETE("/pizzas/:id", handlers.DeletePizzaById) // Route to delete a pizza by ID
}
