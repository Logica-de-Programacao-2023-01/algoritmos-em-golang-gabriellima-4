// Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.
package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Escolha um número.")
	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {
		if i > numero {
			break
		} else if numero%i == 0 {
			fmt.Println("Os divisores do número escolhido são:", i)
		}
	}
}
