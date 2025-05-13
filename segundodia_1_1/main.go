package main

import (
	"fmt"
	"os"
)

func main() {
	txt := []byte("Esse eh o conteudo do meu arquivo novo\n") // Adiciona uma nova linha
	file, err := os.OpenFile("arquivo-novo.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.Write(txt)
	file.WriteString("minha string\n")                // Adiciona uma nova linha
	file.WriteAt([]byte("Texto na posição 80\n"), 80) // Adiciona nova linha na posição 80
}
