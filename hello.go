package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const espera = 20

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		iniciaPrograma()
	}
}

func exibeIntroducao() {
	nome := "Matheus"
	var idade int = 25
	//var versao float32 = 1.1
	var versao = 1.1

	fmt.Println("Olá sr.", nome, "você tem", idade, "anos de idade")
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel versao é:", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciaPrograma() {

	// comando := lerComando()
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo...")
	// } else {
	// 	fmt.Println("Comando digitado nao reconhecido!")
	// }

	switch lerComando() {
	case 1:
		iniciaMonitoramento()
	case 2:
		imprimeLog()
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Comando digitado nao reconhecido!")
		os.Exit(-1)
	}
}

func lerComando() int {
	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	//fmt.Println("O endereço da var comando é:", &comando)
	return comando
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := []string{"https://www.alura.com.br", "https://www.uol.com.br", "https://www.caelum.com.br", "https://random-status-code.herokuapp.com/"}

	sites := buscaSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(espera * time.Second)
	}
}

func testaSite(site string) {
	resposta, erro := http.Get(site)
	if erro != nil {
		fmt.Print("Ocorreu um erro:", erro)
	} else {
		registraLog(site, resposta.StatusCode)
	}
}

func buscaSitesDoArquivo() []string {
	var sites []string
	arquivo, erro := os.Open("sites.txt")
	//arquivo, erro := ioutil.ReadFile("sites.txt")
	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if erro == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status int) {
	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	} else {
		arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- response StatusCode:" + strconv.Itoa(status) + "\n")
	}

	arquivo.Close()

	if status == 200 {
		fmt.Println(site, " - Status OK")
	} else {
		fmt.Println(site, " - Status NOK", status)
	}
}

func imprimeLog() {
	fmt.Println("Exibindo Logs...")

	arquivo, erro := ioutil.ReadFile("log.txt")
	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}
	fmt.Println(string(arquivo))
}
