//Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação, de acordo com a seguinte tabela:
//até 9 anos: mirim
//10 a 13 anos: infantil
//14 a 17 anos: juvenil
//maiores de 18 anos: adulto

package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Qual é a sua idade em anos?")
	fmt.Scan(&idade)

	if idade <= 9 {
		fmt.Println("A sua classificação é Mirim.")
	} else if idade >= 10 && idade < 14 {
		fmt.Println("A sua calssificação é Infantil.")
	} else if idade >= 14 && idade < 18 {
		fmt.Println("A sua classificação é Juvenil.")
	} else if idade >= 18 {
		fmt.Println("A sua classificação é Adulto.")
	}
}
