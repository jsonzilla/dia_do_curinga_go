package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLeapYear(year int) bool {
	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}

func fixDay(year int, day int) int {
	leap := leapYearInt(year)
	if day > 60-leap {
		return day - 60
	}
	return day + 305
}

func fixDayVerifyLeapYear(year int, day int) int {
	return fixDay(year, day)
}

func fixYear(year int) int {
	if year < 1790 {
		return 1790 - year
	}

	return year - 1790
}

func suitYear(year int) int {
	return (fixYear(year) / 13) % 4
}

func cardYear(year int) int {
	return fixYear(year) % 13
}

func leapYearInt(year int) int {
	if isLeapYear(year) {
		return 1
	}
	return 0
}

func seasons(day int, year int) int {
	leap := leapYearInt(year)
	if day <= (62 - leap) {
		return 1
	}
	if day <= (154 - leap) {
		return 2
	}
	if day <= (247 - leap) {
		return 3
	}
	if day <= (338 - leap) {
		return 0
	}
	if day <= (367 - leap) {
		return 1
	}
	return 1
}

func cardMonth(day int) int {
	return (day / 28) % 13
}

func suitWeek(day int) int {
	return ((day / 7) / 13) % 4
}

func cardWeek(day int) int {
	return (day / 7) % 13
}

func suitDay(day int) int {
	if day == 0 {
		return 4
	}
	return ((day - 1) / 13) % 4
}

func cardDay(day int) int {
	if day == 0 {
		return 13
	}
	return (day - 1) % 13
}

func feb(day int, year int) bool {
	if day <= (28 + leapYearInt(year)) {
		return true
	}
	return false
}

func validDate(day int, month int, year int) bool {
	if day < 1 || day > 31 || year == 0 || month < 1 || month > 12 {
		return false
	}
	if (month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) && day <= 31 {
		return true
	}
	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		return true
	}
	if month == 2 {
		return feb(day, year)
	}

	return false
}

func dayOfYear(day int, month int, year int) int {
	if !validDate(day, month, year) {
		return 0
	}
	return countDays(day, month, year)
}

func countDays(day int, month int, year int) int {
	leap := leapYearInt(year)
	switch month {
	case 1:
		return day
	case 2:
		return day + 31
	case 3:
		return day + 59 + leap
	case 4:
		return day + 90 + leap
	case 5:
		return day + 120 + leap
	case 6:
		return day + 151 + leap
	case 7:
		return day + 181 + leap
	case 8:
		return day + 212 + leap
	case 9:
		return day + 243 + leap
	case 10:
		return day + 273 + leap
	case 11:
		return day + 304 + leap
	case 12:
		return day + 334 + leap
	}
	return 0
}

func SimpleVersion(day int, month int, year int) string {
	if !validDate(day, month, year) {
		return ""
	}

	cards := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "Jo", "Jd"}
	suits := [...]string{"O", "P", "C", "E", ""}
	days := fixDayVerifyLeapYear(year, dayOfYear(day, month, year))

	sDay := cards[cardDay(days)] + suits[suitDay(days)]
	sWeek := cards[cardWeek(days)] + suits[suitWeek(days)]
	sMonth := cards[cardMonth(days)] + suits[seasons(day, year)]
	sYear := cards[cardYear(year)] + suits[suitYear(year)]
	output := sDay + sWeek + sMonth + sYear

	return output
}

func LongVersion(day int, month int, year int) string {
	if !validDate(day, month, year) {
		return ""
	}

	cards := [...]string{"As", "Dois", "Tres", "Quatro", "Cinco",
		"Seis", "Sete", "Oito", "Nove", "Dez",
		"Valete", "Dama", "Rei", "do Curinga"}
	suites := [...]string{" de ouros", " de paus", " de copas", " de espadas"}

	days := fixDayVerifyLeapYear(year, dayOfYear(day, month, year))

	output := "\n\tDia " + cards[cardDay(days)] + suites[suitDay(days)]
	output += "\n\tSemana " + cards[cardWeek(days)] + suites[suitWeek(days)]
	output += "\n\tMes " + cards[cardMonth(days)] + " estacao" + suites[seasons(day, year)]
	output += "\n\tAno " + cards[cardYear(year)] + suites[suitYear(year)]
	output += "\n\t" + strconv.Itoa(day) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(year) + " e dia numero " + strconv.Itoa(days)

	return output
}

func header() {
	fmt.Println("Entre com dia mes e ano (separados por enter) e precione ctrl-c")
	fmt.Println("\n\tEntre com dia mes e ano (separados por espaco):")
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func showOutput(day int, month int, year int) {
	fmt.Println("\n\tCalendario de Paciencia de LongVersion")
	fmt.Println("\t---------------------------------")
	fmt.Println(LongVersion(day, month, year))
	fmt.Println("\n\tSimples -- " + SimpleVersion(day, month, year))
}

type simpleDate struct {
	day int
	month int
	year int
}

func clearInput(input string) simpleDate {
	newInput := strings.Replace(input, "\n", " ", -1)
	args := strings.Split(newInput, " ")
	if len(args) >= 3 {
		day, _ := strconv.Atoi(args[0])
		month, _ := strconv.Atoi(args[1])
		year, _ := strconv.Atoi(args[2])
		return simpleDate{day, month, year}
	}

	return simpleDate{0, 0, 0}
}

func main() {
	header()
	input := readUserInput()
	cleanInput := clearInput(input)
	showOutput(cleanInput.day, cleanInput.month, cleanInput.year)
}
