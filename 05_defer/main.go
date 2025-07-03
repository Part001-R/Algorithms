package main

import "fmt"

func main() {

	v := 10
	defer fmt.Println(v) // 10

	defer func(vv int) {
		fmt.Println(vv) // 10
	}(v)

	v = 11
	defer func(vv int) {
		fmt.Println(vv) // 11
	}(v)

	defer func() {
		fmt.Println(v) //12
	}()

	defer fmt.Println(v) // 11

	v = 12
	fmt.Println(v) // 12
}

// В выводе
/*
	12
	11
	12
	11
	10
	10
*/
