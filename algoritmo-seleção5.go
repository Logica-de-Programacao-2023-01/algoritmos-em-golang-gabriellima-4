// Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.
package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Escolha um número.")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("Este número é múltiplo de 3 e 5 ao mesmo tempo.")
	} else if numero%3 != 0 && numero%5 != 0 {
		fmt.Println("Este número não é múltiplo de 3 e 5 ao mesmo tempo.")
	}
}
