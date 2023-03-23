// Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles,+
// se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.
package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)
	fmt.Print("Escolha um outro número.")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println(x * y)
	} else if x < 0 && y > 0 {
		fmt.Println(x + y)
	} else if x > 0 && y < 0 {
		fmt.Println(x + y)
	} else if x < 0 && y < 0 {
		fmt.Println(x + y)
	}
}
