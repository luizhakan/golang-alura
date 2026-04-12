package main // É o principal pacote do programa

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
) // pacote de formatação e sistema operacional

/*
"net/http" é pacote mais específico à nossa necessidade. Já que nele temos funções para realizar requisições Get e Post.
*/

const monitoramentos = 3
const delay = 5

func main() {
	// Não recebe e nem retorna nada
	exibeIntroducao()
	for {
		// testaCap()
		// exibeNomes()
		exibeMenu()
		// fmt.Println("Comando", comando)

		// nome := devolveNome()
		// fmt.Printf("Tipo de nome %T\n", nome)

		// _, idade := devolveNomeEIdade() // o _ aqui significa que, o primeiro retorno da função não será utilizado. Pode ser em qualquer índice do retorno
		// _, idade := devolveNomeEIdade()
		// fmt.Println(nome)
		// fmt.Println(idade)

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
			iniciarMonitoramento()

		case 2:
			exibirLogs()

		case 3:
			fmt.Println("Saindo...")
			os.Exit(0)

		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
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

func devolveNomeEIdade() (string, int) {
	var nome string = "Luiz"
	var idade int = 23
	return nome, idade
}

func iniciarMonitoramento() {
	fmt.Println("Iniciar monitoramento selecionado")
	// var sites [4]string // array
	// sites[0] = "https://httpbin.org/status/500"
	// sites[1] = "https://httpbin.org/status/401"
	// sites[2] = "https://httpbin.org/status/201"
	// sites[3] = "https://httpbin.org/status/200"
	// sites := []string{"https://httpbin.org/status/500", "https://httpbin.org/status/401", "https://httpbin.org/status/201", "https://httpbin.org/status/200"}
	var sites []string = leSitesDoArquivo()

	fmt.Println(sites)
	// fmt.Println("O meu array tem capacidade para:", cap(sites), "itens")

	// resp, _ := http.Get(sites)

	// switch resp.StatusCode {
	// case 200:
	// 	fmt.Println("Está online")
	// case 404:
	// 	fmt.Println("Está fora do AR!")
	// default:
	// 	fmt.Println(resp.Status)
	// }

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	for indice := 0; indice < monitoramentos; indice++ {
		for indice, site := range sites {
			fmt.Println("\nPosição:", indice)
			fmt.Println("Site:", site)
			// fmt.Println("O site", site, "recebe o status:", resp.Status)
			testaSite(site)
			fmt.Println()
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	switch resp.StatusCode {
	case 200:
		fmt.Println("Está online")
		registraLog(site, true)
	case 404:
		fmt.Println("Está fora do AR!")
		registraLog(site, false)
	default:
		fmt.Println(resp.Status)
		registraLog(site, false)

	}
}

// func exibeNomes() {
// 	nomes := []string{"Jon", "Snow"} // slice
// 	fmt.Println("O meu slice tem:", len(nomes), "itens")
// 	fmt.Println("O meu slice tem capacidade para:", cap(nomes), "itens")

// 	nomes = append(nomes, "Targaryen")
// 	fmt.Println("O meu slice tem:", len(nomes), "itens depois de eu adicionar mais 1")
// 	fmt.Println("O meu slice tem capacidade para:", cap(nomes), "itens depois de eu adicionar mais 1")

// 	fmt.Println(nomes)
// }

// func testaCap() {
// 	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21}
// 	// fmt.Println(cap(pontosPlanningPoker))

// 	pontosPlanningPoker = append(pontosPlanningPoker, 40)
// 	// fmt.Println(cap(pontosPlanningPoker))

// 	pontosPlanningPoker = append(pontosPlanningPoker, 41)
// 	// fmt.Println(cap(pontosPlanningPoker))

// 	pontosPlanningPoker = append(pontosPlanningPoker, 33)
// 	// fmt.Println(cap(pontosPlanningPoker))

// 	// fmt.Println(pontosPlanningPoker)
// }

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, erro := os.Open("sites.txt")

	if erro != nil {
		fmt.Println("Ocorreu um erro", erro)
	} else {
		fmt.Println(arquivo)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, erro := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func exibirLogs() {
	arquivo, erro := os.Open("logs.txt")

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	} else {
		fmt.Println(arquivo)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		if err == io.EOF {
			break
		}

		fmt.Println("linha")
	}
	arquivo.Close()
}
