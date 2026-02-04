package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	sm  = semaphore.NewWeighted(10)
	wg  sync.WaitGroup
	ctx = context.Background()
)

// ANSI escape codes для цветов
const (
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

func main() {
	// Запросы по 10
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reqBy10(i, sm)
	}

	// Запросы по 5
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reqBy5(i, sm)
	}

	// Запросы по 3
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reqBy3(i, sm)
	}

	wg.Wait()
}

// Функция с запросом по 10
func reqBy10(numb int, sem *semaphore.Weighted) {
	defer wg.Done()

	val := 10

	fmt.Printf("%sreqBy10: горутина <%d> пробует захватить <%d>%s\n", Yellow, numb, val, Reset)

	err := sem.Acquire(ctx, int64(val))
	if err != nil {
		log.Fatalf("reqBy10: ошибка в горутине <%d>, при запросе <%d> у семаформа", numb, val)
	}

	fmt.Printf("%sreqBy10: горутина <%d> запросила у семафора <%d>%s\n", Green, numb, val, Reset)
	time.Sleep(2 * time.Second)

	sem.Release(int64(val))
	fmt.Printf("%sreqBy10: горутина <%d> освободила <%d>%s\n", Red, numb, val, Reset)
}

// Функция с запросом по 5
func reqBy5(numb int, sem *semaphore.Weighted) {
	defer wg.Done()

	val := 5

	fmt.Printf("%sreqBy05: горутина <%d> пробует захватить <%d>%s\n", Yellow, numb, val, Reset)

	err := sem.Acquire(ctx, int64(val))
	if err != nil {
		log.Fatalf("reqBy05: ошибка в горутине <%d>, при запросе <%d> у семаформа", numb, val)
	}

	fmt.Printf("%sreqBy05: горутина <%d> запросила у семафора <%d>%s\n", Green, numb, val, Reset)
	time.Sleep(2 * time.Second)

	sem.Release(int64(val))
	fmt.Printf("%sreqBy05: горутина <%d> освободила <%d>%s\n", Red, numb, val, Reset)
}

// Функция с запросом по 3
func reqBy3(numb int, sem *semaphore.Weighted) {
	defer wg.Done()

	val := 3

	fmt.Printf("%sreqBy03: горутина <%d> пробует захватить <%d>%s\n", Yellow, numb, val, Reset)

	err := sem.Acquire(ctx, int64(val))
	if err != nil {
		log.Fatalf("reqBy03: ошибка в горутине <%d>, при запросе <%d> у семаформа", numb, val)
	}

	fmt.Printf("%sreqBy03: горутина <%d> запросила у семафора <%d>%s\n", Green, numb, val, Reset)
	time.Sleep(2 * time.Second)

	sem.Release(int64(val))
	fmt.Printf("%sreqBy03: горутина <%d> освободила <%d>%s\n", Red, numb, val, Reset)
}
