//Faça um algoritmo que leia a altura e o sexo de uma pessoa e mostre se ela está abaixo,+
//dentro ou acima do peso ideal, utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.

package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	var sexo string
	fmt.Print("Qual é a sua altura (em metros)?")
	fmt.Scan(&altura)
	fmt.Print("Qual é o seu peso em Kg?")
	fmt.Scan(&peso)
	fmt.Print("Qual é o seu sexo?")
	fmt.Scan(&sexo)

	var IMC float64 = peso / (altura * altura)
	fmt.Println(IMC)

	if IMC < 18.49 {
		fmt.Println("Você está abaixo do peso ideal. Classificação: Magreza.")
	} else if IMC > 18.5 && IMC <= 24.99 {
		fmt.Println("Você está dentro da faixa de peso ideal. Classificação: Normal.")
	} else if IMC > 25 && IMC <= 29.99 {
		fmt.Println("Você está acima do peso ideal. Classificação: Sobrepeso.")
	} else if IMC > 30 && IMC <= 34.99 {
		fmt.Println("Você está muito acima do peso ideal. Classificação: Obesidade grau 1.")
	} else if IMC > 35 && IMC <= 39.99 {
		fmt.Println("Você está muito, mas muito acima do peso ideal. Classificação: Obesidade grau 2.")
	} else if IMC > 40 {
		fmt.Println("Você está extremamente acima do peso ideal. Classificação< Obesidade grau 3.")
	}

}
