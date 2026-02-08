package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, sem chan struct{}) {

	// Получение доступа к ресурсу
	sem <- struct{}{} // ожидание свободного места
	defer func() {
		wg.Done()
		<-sem // освобождение
	}()

	// Имитация работы
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d is done\n", id)
}

func main() {

	const maxWorkers = 3
	sem := make(chan struct{}, maxWorkers)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, sem)
	}

	wg.Wait()
}
