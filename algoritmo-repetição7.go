// Faça um algoritmo que imprima os números de 1 a 100,
// substituindo os múltiplos de 3 pela palavra "Fizz" e os múltiplos de 5 pela palavra "Buzz".
// Para os números que são múltiplos de ambos, utilize a palavra "FizzBuzz".
package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		fmt.Println(i)
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 && i%3 != 0 {
			fmt.Println("Buzz", i)
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
		}
	}
}
