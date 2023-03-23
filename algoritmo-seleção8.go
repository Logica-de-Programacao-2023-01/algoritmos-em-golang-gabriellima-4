// Faça um algoritmo que leia um número inteiro entre 1 e 7
// e mostre o dia da semana correspondente (1 = domingo, 2 = segunda-feira, etc.).
package main

import "fmt"

func main() {
	var dia int
	fmt.Print("Que dia da semana é hoje?")
	fmt.Scan(&dia)

	if dia == 1 {
		fmt.Println("Hoje é Domingo.")
	} else if dia == 2 {
		fmt.Println("Hoje è Segunda-feira.")
	} else if dia == 3 {
		fmt.Println("Hoje é Terça-feira.")
	} else if dia == 4 {
		fmt.Println("Hoje é Quarta-feira.")
	} else if dia == 5 {
		fmt.Println("Hoje é Quinta-feira.")
	} else if dia == 6 {
		fmt.Println("Hoje é Sexta-feira. Sextou!")
	} else if dia == 7 {
		fmt.Println("Hoje é Sábado. Amém")
	} else if dia > 7 {
		fmt.Println("A semana só tem 7 dias, doidão.")
	}

}
