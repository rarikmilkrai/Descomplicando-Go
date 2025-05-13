// O primeiro método pra ler 
// arquivos é a gente utilizar uma função 
// chamada ReadFile do pacote os que nós já 
// vimos. Essa função lê todo o conteudo do arquivo!
 

package main



import (  

	"fmt"

	"os"

)



func main() {  

	file, err := os.ReadFile("arquivo.txt")

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println("File:", string(file))

}



// Existem algumas outras formas de ler arquivos 
// em Go - vamos falar de mais duas, usando o 
// pacote `bufio`, podemos ler arquivos de forma
// parcial, o que é extremamente eficiente para 
// arquivos maiores. Pensem que toda leitura / 
// escrita de arquivo pode ser feita em pequenas
//  partes, e isso as vezes vai ser melhor do 
//  que fazer a leitura ou escrita ou de um c
//  aracter por vez ou tudo de uma vez só.



// A segunda forma que vamos ver aqui fazer a l
// eitura de um arquivo vai ser criando um 
// Scanner para esse arquivo. Esse scanner 
// consegue ler de linha em linha!


file, err := os.Open("arquivo.txt")

if err != nil {

	log.Fatal(err)

}

scanner := bufio.NewScanner(file)

for scanner.Scan() {

	log.Println(scanner.Text())

}

err = scanner.Err()

if err != nil {

	log.Fatal(err)

}

// A terceira forma é utilizando um reader e 
// lendo byte a byte.




﻿﻿    file, err := os.Open("arquivo.txt")

if err != nil {

	log.Fatal(err)

}

reader := bufio.NewReader(file)

for {
b, err := reader.ReadByte(b)
	if err != nil {
		fmt.Println("Error reading file:", err)
		break
	}
	fmt.Println(string(b))
}
}

