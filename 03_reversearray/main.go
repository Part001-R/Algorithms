package main

import "fmt"

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := reverse(arr)

	fmt.Println("Исходный массив :", arr)
	fmt.Println("Результат работы:", res)

}

func reverse(a []int) []int {

	arr := make([]int, len(a), cap(a))
	copy(arr, a)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {

		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
