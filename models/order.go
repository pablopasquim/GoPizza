package models

import "fmt"

type Order struct {
	OrderNumber string
	Pizzas      []Pizza
	Customer    Customer
	Total       float64
	Paid		bool
}

func NewOrder(customer Customer) *Order {
	return &Order{
		Customer: customer,
		Pizzas:   []Pizza{},
		Paid:     false,
	}
}

func (o *Order) AddPizza(p Pizza) {
	o.Pizzas = append(o.Pizzas, p)
	o.Total += float64(p.Price)
}

func (o *Order) ProcessPayment() {
	fmt.Printf("Processando pagamento de R$ %.2f...\n", o.Total)
	o.Paid = true
	fmt.Println("Pagamento concluído com sucesso ✅")
}

func (o *Order) PrintSummary() {
	fmt.Println("========== RESUMO DO PEDIDO ==========")
	fmt.Println("Cliente:", o.Customer.Name)
	for _, pizza := range o.Pizzas {
		fmt.Printf("- %s (R$ %.2f)\n", pizza.Name, pizza.Price)
	}
	fmt.Printf("Total: R$ %.2f\n", o.Total)
	fmt.Println("Pago:", o.Paid)
	fmt.Println("======================================")
}
