// Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.
package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Qual é o seu peso em kg?")
	fmt.Scan(&peso)
	fmt.Println("O seu peso em libras é:", peso*2.20462)
}
