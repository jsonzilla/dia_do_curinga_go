package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ehAnoBissexto(ano int) bool {
	return (ano%400 == 0) || (ano%4 == 0 && ano%100 != 0)
}

func corrigeDiaFrode(ano int, dia int) int {
	leap := intAnoBissexto(ano)
	if dia > 60-leap {
		return dia - 60
	}
	return dia + 305
}

func corrigeDiaFrodeVerificaBissexto(dia int, ano int) int {
	return corrigeDiaFrode(ano, dia)
}

func corrigeAnoFrode(ano int) int {
	if ano < 1790 {
		return 1790 - ano
	}

	return ano - 1790
}

func naipeAnoFrode(ano int) int {
	return (corrigeAnoFrode(ano) / 13) % 4
}

func cartaAnoFrode(ano int) int {
	return corrigeAnoFrode(ano) % 13
}

func intAnoBissexto(ano int) int {
	if ehAnoBissexto(ano) {
		return 1
	}
	return 0
}

func estacoesFrode(dia int, ano int) int {
	biss := intAnoBissexto(ano)
	if dia <= (62 - biss) {
		return 1
	}
	if dia <= (154 - biss) {
		return 2
	}
	if dia <= (247 - biss) {
		return 3
	}
	if dia <= (338 - biss) {
		return 0
	}
	if dia <= (367 - biss) {
		return 1
	}
	return 1
}

func mesNumeral(dia int) int {
	return (dia / 28) % 13
}

func naipeSemanaFrode(dia int) int {
	return ((dia / 7) / 13) % 4
}

func cartaSemanaFrode(dia int) int {
	return (dia / 7) % 13
}

func naipeDiaFrode(dia int) int {
	if dia == 0 {
		return 4
	}
	return ((dia - 1) / 13) % 4
}

func cartaDiaFrode(dia int) int {
	if dia == 0 {
		return 13
	}
	return (dia - 1) % 13
}

func feb(dia int, ano int) bool {
	if dia <= (28 + intAnoBissexto(ano)) {
		return true
	}
	return false
}

func ehDataValida(dia int, mes int, ano int) bool {
	if dia < 1 || dia > 31 || ano == 0 || mes < 1 || mes > 12 {
		return false
	}
	if (mes == 1 || mes == 3 || mes == 5 || mes == 7 || mes == 8 || mes == 10 || mes == 12) && dia <= 31 {
		return true
	}
	if (mes == 4 || mes == 6 || mes == 9 || mes == 11) && dia > 30 {
		return true
	}
	if mes == 2 {
		return feb(dia, ano)
	}

	return false
}

func diaDoAno(dia int, mes int, ano int) int {
	if !ehDataValida(dia, mes, ano) {
		return 0
	}
	return contaPorMes(dia, mes, ano)
}

func contaPorMes(dia int, mes int, ano int) int {
	biss := intAnoBissexto(ano)
	switch mes {
	case 1:
		return dia
	case 2:
		return dia + 31
	case 3:
		return dia + 59 + biss
	case 4:
		return dia + 90 + biss
	case 5:
		return dia + 120 + biss
	case 6:
		return dia + 151 + biss
	case 7:
		return dia + 181 + biss
	case 8:
		return dia + 212 + biss
	case 9:
		return dia + 243 + biss
	case 10:
		return dia + 273 + biss
	case 11:
		return dia + 304 + biss
	case 12:
		return dia + 334 + biss
	}
	return 0
}

func FrodeSimples(dia int, mes int, ano int) string {
	if !ehDataValida(dia, mes, ano) {
		return ""
	}

	cartas := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "Jo", "Jd"}
	naipes := [...]string{"O", "P", "C", "E", ""}
	numerodedia := corrigeDiaFrodeVerificaBissexto(diaDoAno(dia, mes, ano), ano)

	sDia := cartas[cartaDiaFrode(numerodedia)] + naipes[naipeDiaFrode(numerodedia)]
	sSemana := cartas[cartaSemanaFrode(numerodedia)] + naipes[naipeSemanaFrode(numerodedia)]
	sMes := cartas[mesNumeral(numerodedia)] + naipes[estacoesFrode(dia, ano)]
	sAno := cartas[cartaAnoFrode(ano)] + naipes[naipeAnoFrode(ano)]
	sFrode := sDia + sSemana + sMes + sAno

	return sFrode
}

func Frode(dia int, mes int, ano int) string {
	if !ehDataValida(dia, mes, ano) {
		return ""
	}

	cartas := [...]string{"As", "Dois", "Tres", "Quatro", "Cinco",
		"Seis", "Sete", "Oito", "Nove", "Dez",
		"Valete", "Dama", "Rei", "do Curinga"}
	naipes := [...]string{" de ouros", " de paus", " de copas", " de espadas"}

	numerodedia := corrigeDiaFrodeVerificaBissexto(diaDoAno(dia, mes, ano), ano)

	saida := ""

	if numerodedia < 365 {
		saida += "\n\tDia " + cartas[cartaDiaFrode(numerodedia)] + naipes[naipeDiaFrode(numerodedia)]
	}
	if numerodedia == 365 {
		saida += "\n\tDia " + cartas[13]
	}
	if numerodedia == 366 {
		saida += "\n\tDuplo dia " + cartas[13]
	}

	saida += "\n\tSemana " + cartas[cartaSemanaFrode(numerodedia)] + naipes[naipeSemanaFrode(numerodedia)]
	saida += "\n\tMes " + cartas[mesNumeral(numerodedia)] + " estacao" + naipes[estacoesFrode(dia, ano)]
	saida += "\n\tAno " + cartas[cartaAnoFrode(ano)] + naipes[naipeAnoFrode(ano)]
	saida += "\n\t" + strconv.Itoa(dia) + "/" + strconv.Itoa(mes) + "/" + strconv.Itoa(ano) + " e dia numero " + strconv.Itoa(numerodedia)

	return saida
}

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
	fmt.Println(Frode(dia, mes, ano))
	fmt.Println("\n\tSimples -- " + FrodeSimples(dia, mes, ano))
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
		return dataSimples{dia, mes, ano}
	}

	return dataSimples{0, 0, 0}
}

func main() {
	cabecalhoInicial()
	entrada := lerEntradaUsuario()
	entradaLimpa := limparEntrada(entrada)
	exibirDataFrode(entradaLimpa.dia, entradaLimpa.mes, entradaLimpa.ano)
}
