package main // É o principal pacote do programa

import (
	"fmt"
	"reflect"
) // pacote de formatação

func main() {
	// Não recebe e nem retorna nada
	nome := "Luiz" // o := sem a palavra reservada var, quer dizer que eu estou a criar a variável e a inferir o tipo nela automaticamente
	var versao float32 = 1.2
	var idade int = 23

	// nome = 12 // vai dar erro, por que o tipo inferido para nome foi string
	fmt.Println("Olá sr.", nome)
	fmt.Println("Sua idade é:", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo de nome é:", reflect.TypeOf(nome))
}

// fazer o build com go build arquivo.go
// compila e já executa: go run arquivo.go
