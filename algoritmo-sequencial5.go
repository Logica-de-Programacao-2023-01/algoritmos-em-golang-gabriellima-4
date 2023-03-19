// Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.
package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Quantos anos você tem?")
	fmt.Scan(&idade)
	fmt.Println("A sua idade em dias equivale a:", idade*365)

}
