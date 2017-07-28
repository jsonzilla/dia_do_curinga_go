package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../ddc"
)

func cabecalhoInicial() {
	fmt.Println("Entre com dia mes e ano (separados por enter) e precione ctrl-c")
	fmt.Println("\n\tEntre com dia mes e ano (separados por espaco):")
}

func lerEntradaUsuario() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func exibirDataFrode(dia int, mes int, ano int) {
	fmt.Println("\n\tCalendario de Paciencia de Frode")
	fmt.Println("\t---------------------------------")
	fmt.Println(ddc.Frode(dia, mes, ano))
	fmt.Println("\n\tSimples -- " + ddc.FrodeSimples(dia, mes, ano))
}

type dataSimples struct {
	dia int
	mes int
	ano int
}

func limparEntrada(entrada string) dataSimples {
	novaEntrada := strings.Replace(entrada, "\n", " ", -1)
	args := strings.Split(novaEntrada, " ")
	if len(args) >= 3 {
		dia, _ := strconv.Atoi(args[0])
		mes, _ := strconv.Atoi(args[1])
		ano, _ := strconv.Atoi(args[2])
		return ddc.dataSimples{dia, mes, ano}
	}

	return dataSimples{0, 0, 0}
}

func main() {
	cabecalhoInicial()
	entrada := lerEntradaUsuario()
	entradaLimpa := limparEntrada(entrada)
	exibirDataFrode(entradaLimpa.dia, entradaLimpa.mes, entradaLimpa.ano)
}
