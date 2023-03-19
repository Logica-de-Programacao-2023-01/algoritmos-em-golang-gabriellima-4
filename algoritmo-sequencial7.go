// Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.
package main

import "fmt"

func main() {
	var x int
	fmt.Print("Escolha um número inteiro.")
	fmt.Scan(&x)
	fmt.Println("O antecessor do seu número escolhido é:", x-1)
	fmt.Println("O sucessor do seu número escolhido é:", x+1)

}
