// Faça um algoritmo que leia dois números inteiros e mostre o maior deles.
package main

import (
	"fmt"
)

func main() {
	var n1 int
	var n2 int
	fmt.Print("Escolha um número.")
	fmt.Scan(&n1)
	fmt.Print("Escolha um outro número.")
	fmt.Scan(&n2)

	if n1 > n2 {
		fmt.Println("Seu primeiro número é maior que o segundo.")
	} else if n1 == n2 {
		fmt.Println("Os números escolhidos são iguais.")
	}
	if n1 < n2 {
		fmt.Println("Seu primeiro número é menor que o segundo.")
	}

}

