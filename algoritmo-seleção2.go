// Faça um algoritmo que leia três números inteiros e mostre o menor deles.
package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int
	fmt.Print("Escolha um número.")
	fmt.Scan(&n1)
	fmt.Print("Escolha um outro número.")
	fmt.Scan(&n2)
	fmt.Print("Escolha um terceiro número.")
	fmt.Scan(&n3)

	if n1 < n2 && n1 < n3 {
		fmt.Println("Seu primeiro número é menor que o segundo e o terceiro.")
	} else if n1 == n2 && n1 == n3 {
		fmt.Println("Os números escolhidos são iguais.")
	} else if n2 < n1 && n2 < n3 {
		fmt.Println("Seu segundo número é menor que o primeiro e o terceiro.")
	} else if n3 < n1 && n3 < n2 {
		fmt.Println("Seu terceiro número é menor que o primeiro e o segundo.")
	}

}

