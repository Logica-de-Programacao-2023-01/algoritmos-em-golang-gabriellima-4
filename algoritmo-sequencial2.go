// Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.
package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Selecione um número inteiro.")
	fmt.Scan(&numero)
	fmt.Println("O dobro do seu número escolhido é:", numero*2)
	fmt.Println("O triplo do seu número é:", numero*3)
	fmt.Println("O quádruplo do seu número equivale a:", numero*4)

}
