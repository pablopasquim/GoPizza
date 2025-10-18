package main

import(
	"github.com/pablopasquim/GoPizza/models"
)

func main() {

	customer := models.Customer{Name: "Pablo"}


	order := models.NewOrder(customer)

	order.AddPizza(models.Pizza{Name: "Calabresa", Price: 45.0})
	order.AddPizza(models.Pizza{Name: "Frango Catupiry", Price: 50.0})
	order.AddPizza(models.Pizza{Name: "Mussarela", Price: 40.0})

	order.PrintSummary()

	order.ProcessPayment()

	order.PrintSummary()
}

