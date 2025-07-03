package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func main() {

	str := "AAA111BBB222CCC"

	res, t := byRune(str)
	fmt.Println("Исходная строка :", str)
	fmt.Println("Результат работы:", res)
	fmt.Println("Время выполнения:", t)
	fmt.Println()

	res, t = byStringsBuilder(str)
	fmt.Println("Исходная строка :", str)
	fmt.Println("Результат работы:", res)
	fmt.Println("Время выполнения:", t)
	fmt.Println()

}

func byRune(s string) (string, time.Duration) {

	ts := time.Now()

	slResult := make([]rune, 0)

	for _, v := range s {

		if unicode.IsLetter(v) {
			slResult = append(slResult, v)
		}
	}

	return string(slResult), time.Since(ts)

}

func byStringsBuilder(s string) (string, time.Duration) {

	ts := time.Now()

	var sb strings.Builder

	for _, v := range s {
		if unicode.IsLetter(v) {
			sb.WriteRune(v)
		}
	}

	return sb.String(), time.Since(ts)

}
