package main // É o principal pacote do programa

import (
	"fmt"
	"os"
) // pacote de formatação

func main() {
	// Não recebe e nem retorna nada
	exibeIntroducao()
	exibeMenu()
	// fmt.Println("Comando", comando)

	// if comando == 1 {
	// 	fmt.Println("Iniciar monitoramento selecionado")
	// } else if comando == 2 {
	// 	fmt.Println("Exibir logs selecionado")
	// } else if comando == 3 {
	// 	fmt.Println("Saindo...")
	// } else {
	// 	fmt.Println("Comando inválido")
	// }

	switch leComando() {
	case 1:
		fmt.Println("Iniciar monitoramento selecionado")

	case 2:
		fmt.Println("Exibir logs selecionado")

	case 3:
		fmt.Println("Saindo...")
		os.Exit(0)

	default:
		fmt.Println("Comando inválido")
		os.Exit(-1)
	}
}

// fazer o build com go build arquivo.go
// compila e já executa: go run arquivo.go

func exibeIntroducao() {
	nome := "Luiz" // o := sem a palavra reservada var, quer dizer que eu estou a criar a variável e a inferir o tipo nela automaticamente
	versao := 1.2

	// nome = 12 // vai dar erro, por que o tipo inferido para nome foi string
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
