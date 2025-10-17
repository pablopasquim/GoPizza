package main

import "fmt"

func main() {
	var nomePizzaria string = "Pizzaria Go"
	instagram, telefone := "@pizzaria_go", 11951
	fmt.Printf("Nome da pizzaria: %s (instagram: %s) - Telefone: %d", nomePizzaria, instagram, telefone)
}

func getPizza(){
	pizzas := make(map[string]int)

	pizzas["calabresa"] = 10
	pizzas["4 queijos"] = 30
	pizzas["portugues"] = 40

}
