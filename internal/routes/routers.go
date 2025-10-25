package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pablopasquim/GoPizza/internal/handler"
)

func HandleRequest(r *gin.Engine) { // Function to handle routes
	r.GET("/pizzas", handlers.GetPizza) // Route to get pizzas
	r.GET("/pizzas/:id", handlers.GetPizzaID) // Route to get pizza by ID
	r.POST("/pizzas", handlers.PostPizza) // Route to create a new pizza
	r.PUT("/pizzas/:id", handlers.PutPizza) // Route to update a pizza
	r.DELETE("/pizzas/:id", handlers.DeletePizza) // Route to delete a pizza by ID
}
