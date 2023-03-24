// Faça um algoritmo que leia três números reais e mostre-os em ordem crescente.
package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)
	fmt.Print("Escolha mais um número.")
	fmt.Scan(&y)
	fmt.Print("Escolha um terceiro número.")
	fmt.Scan(&z)

	if x > y && x > z && y > z {
		fmt.Println(z, y, x)
	} else if y > x && x > z && y > z {
		fmt.Println(z, x, y)
	} else if x > y && x > z && y < z {
		fmt.Println(y, z, x)
	} else if y > x && x < z && y > z {
		fmt.Println(x, z, y)
	} else if z > x && z > y && x > y {
		fmt.Println(y, x, z)
	} else if z > x && z > y && x < y {
		fmt.Println(x, y, z)
	}
}
 
