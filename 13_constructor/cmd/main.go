package main

import (
	"e/internal/example"
	"fmt"
)

func main() {

	inst := example.New(111, 222)
	fmt.Println("Первый вызов конструктора")
	fmt.Printf("Значение VarA: <%d>\n", inst.GetValA())
	fmt.Printf("Значение VarB: <%d>\n", inst.GetValB())

	fmt.Println()

	inst = example.New(333, 444)
	fmt.Println("Второй вызов конструктора")
	fmt.Printf("Значение VarA: <%d>\n", inst.GetValA())
	fmt.Printf("Значение VarB: <%d>\n", inst.GetValB())

}
