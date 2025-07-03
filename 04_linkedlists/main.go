package main

import "fmt"

type el struct {
	next *el
	num  int
}

func main() {

	base := &el{}

	e := base
	for i := 0; i < 10; i++ {
		e.num = i
		e.next = &el{}

		e = e.next
	}

	fmt.Println("Вывод")
	e = base
	for e.next != nil {
		fmt.Println(e.num)
		e = e.next
	}
}
