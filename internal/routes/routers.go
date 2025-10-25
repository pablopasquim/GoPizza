package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/pablopasquim/GoPizza/internal/handler"
)

func HandleRequest(r *gin.Engine) { // Function to handle routes
	r.GET("/pizzas", handlers.GetPizza) // Route to get pizzas
}
