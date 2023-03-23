//Faça um algoritmo que leia o salário de um funcionário e+
//calcule o seu novo salário com um aumento de 10% se o salário for menor que R$ 1000,00;+
//ou de 5% se o salário for maior ou igual a R$ 1000,00.

package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Qual é o seu salário atual?")
	fmt.Scan(&salario)

	if salario < 1000 {
		fmt.Println("Você receberá um aumento salarial de 10%. Parabéns."+
			"Seu salário atual agora é:", salario+salario*10/100)
	} else if salario >= 1000 {
		fmt.Println("Você receberá um aumento salarial de 5%. Parabéns."+
			"Seu salário atual agora é:", salario+salario*5/100)
	}
}
