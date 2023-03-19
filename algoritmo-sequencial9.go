// Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.
package main

import "fmt"

func main() {
	var preco_produto float64
	fmt.Print("Qual é o preço do produto?")
	fmt.Scan(&preco_produto)
	fmt.Print("O produto está com 10% de desconto no momento.")
	fmt.Println("O preço do produto com desconto é:", preco_produto-preco_produto*10/100)

}
