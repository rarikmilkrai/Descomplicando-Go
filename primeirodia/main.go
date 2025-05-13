package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func dividir(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func diminuir(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%d", &num2)

	fmt.Printf("Soma: %d\n", soma(num1, num2))
	fmt.Printf("Divisão: %d\n", dividir(num1, num2))
	fmt.Printf("Subtração: %d\n", diminuir(num1, num2))
	fmt.Printf("Multiplicação: %d\n", multiplicar(num1, num2))
}
