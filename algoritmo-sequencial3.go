// Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).
package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	fmt.Print("Qual é o seu peso?")
	fmt.Scan(&peso)
	fmt.Print("Qual é a sua altura?")
	fmt.Scan(&altura)
	fmt.Println("De acordo com suas informações, o seu IMC é:", peso/(altura*altura))

}
