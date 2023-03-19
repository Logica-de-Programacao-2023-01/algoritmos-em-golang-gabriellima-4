// Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 15%.
package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Qual é o seu salário atual?")
	fmt.Scan(&salario)
	fmt.Print("Você receberá um aumento de 15%. Parabéns!")
	fmt.Println("O seu novo salário será:", salario+salario*15/100)
}
