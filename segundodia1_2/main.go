package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func percorreDiretorio(path string, d fs.DirEntry, err error) error {
	fmt.Println("É diretorio? ", d.IsDir())
	info, _ := d.Info()
	fmt.Println("Name:: ", info.Name())
	fmt.Println("-----")
	return nil
}

func main() {
	filepath.WalkDir(".", percorreDiretorio)
}

// O código percorre o diretório atual e imprime informações sobre cada arquivo e subdiretório.
// O método WalkDir é usado para percorrer diretórios e subdiretórios.
// A função percorreDiretorio é chamada para cada entrada no diretório, onde:
// - Verificamos se a entrada é um diretório usando d.IsDir().
// - Obtemos informações sobre a entrada usando d.Info().
// - Imprimimos o nome da entrada e se é um diretório ou não.
// O resultado é uma lista de arquivos e diretórios no diretório atual, com informações sobre cada um deles.
// O código pode ser executado em um terminal Go, e o resultado será uma lista de arquivos e diretórios no diretório atual, com informações sobre cada um deles.
// O código pode ser executado em um terminal Go, e o resultado será uma lista de arquivos e diretórios no diretório atual, com informações sobre cada um deles.
