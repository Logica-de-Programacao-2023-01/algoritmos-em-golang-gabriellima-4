// Faça um algoritmo que leia o número de dias trabalhados por um funcionário e o valor da sua diária e calcule o seu salário.
package main

import "fmt"

func main() {
	var dias_trabalhados float64
	var valor_diaria float64
	fmt.Print("Você trabalha por quantos dias num mês?")
	fmt.Scan(&dias_trabalhados)
	fmt.Print("Qual é o valor da sua diária?")
	fmt.Scan(&valor_diaria)
	fmt.Println("O seu salário mensal é:", dias_trabalhados*valor_diaria)

}
