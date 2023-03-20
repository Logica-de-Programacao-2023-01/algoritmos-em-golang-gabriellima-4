// Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é impar.")
	}

}

