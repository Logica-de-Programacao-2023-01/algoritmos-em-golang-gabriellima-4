
// Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.
package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	var numero3 int
	fmt.Print("Escolha um número.")
	fmt.Scan(&numero1)
	fmt.Print("Escolha um segundo número.")
	fmt.Scan(&numero2)
	fmt.Print("Escolha um terceiro e último número.")
	fmt.Scan(&numero3)
	fmt.Println("A soma dos três números escolhidos é igual a:", numero1+numero2+numero3)

}
