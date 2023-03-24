// Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.
package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)

	for i := 1; i <= 10; i++ {
		fmt.Println(x * float64(i))

	}
}
