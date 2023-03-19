// Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.
package main

import "fmt"

func main() {
	var n1 float64
	var n2 float64
	var n3 float64
	fmt.Print("Escolha um número.")
	fmt.Scan(&n1)
	fmt.Print("Escolha mais um número.")
	fmt.Scan(&n2)
	fmt.Print("Escolha um último número.")
	fmt.Scan(&n3)
	fmt.Println("Para calcular a média ponderada desses valores,"+
		" o primeiro passo é multiplicá-los com seus pesos e somar:", n1*2+n2*3+n3*5)
	fmt.Println("Agora, o segundo passo é calcular a soma dos pesos:", 2+3+5)
	fmt.Println("E como último passo, tem-se que dividir as duas somas:", (n1*2+n2*3+n3*5)/(2+3+5))
}
