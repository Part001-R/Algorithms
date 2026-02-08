package main

import (
	"container/list"
	"fmt"
)

type testType struct {
	Name string
	Age  int
}

func main() {

	// Конструтор
	myList := list.New()

	// Добавление записи.
	el1 := myList.PushFront(testType{
		Name: "A",
		Age:  1,
	})

	// Добавление записи перед элементом
	myList.InsertBefore(testType{
		Name: "B",
		Age:  2,
	}, el1)

	// Проход по списку.
	for el := myList.Front(); el != nil; el = el.Next() {
		v := el.Value.(testType)
		fmt.Printf("Name:%s Age:%d\n", v.Name, v.Age)
	}

	// Удаление элемента.
	fmt.Println("Удаление элемента")
	myList.Remove(el1)

	for el := myList.Front(); el != nil; el = el.Next() {
		v := el.Value.(testType)
		fmt.Printf("Name:%s Age:%d\n", v.Name, v.Age)
	}
}
