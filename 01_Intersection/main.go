package main

import (
	"fmt"
	"time"
)

func main() {

	data1 := []int{1, 1, 1, 44, 1, 1, 1, 1, 1, 1, 1, 1, 100}
	data2 := []int{1, 4, 1, 45, 56, 100, 20, 44}

	result, timeRun := doIntersection(data1, data2)

	fmt.Println("Slice 1:", data1)
	fmt.Println("Slice 2:", data2)
	fmt.Println("Result :", result)
	fmt.Println("Time execution", timeRun)
}

func doIntersection(a, b []int) ([]int, time.Duration) {

	timeStart := time.Now()

	slA := make([]int, len(a), cap(a))
	slB := make([]int, len(b), cap(b))
	resultSlice := make([]int, 0)

	copy(slA, a)
	copy(slB, b)

	for i := 0; i < len(slA) && len(slB) != 0; i++ {

		seachEl := slA[i]

		for j := 0; j < len(slB); j++ {
			if slB[j] == seachEl {
				resultSlice = append(resultSlice, seachEl)
				slB = deleteElement(slB, j)
			}
		}
	}

	return resultSlice, time.Since(timeStart)
}

func deleteElement(sl []int, id int) []int {

	return append(sl[:id], sl[id+1:]...)

}
