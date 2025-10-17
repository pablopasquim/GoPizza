package main

import "fmt"

func main() {
	var nomePizzaria string = "Pizzaria Go"
	instagram, telefone := "@pizzaria_go", 11951
	fmt.Printf("Nome da pizzaria: %s (instagram: %s) - Telefone: %d", nomePizzaria, instagram, telefone)
}

func getPizza(){
	pizzas := make(map[string]int)

	pizzas["calabresa"] = 20.00
	pizzas["4 queijos"] = 30.00
	pizzas["portugues"] = 25.00

	fmt.Println("\nCardápio de pizzas:")
	for sabor, preco := range pizzas{
		fmt.Printf("- %s: %d", sabor, preco)
	}
}
