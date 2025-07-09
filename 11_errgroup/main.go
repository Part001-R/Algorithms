package main

import (
	"context"
	"fmt"

	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	// Первая горутина
	g.Go(func() error {
		fmt.Println("Запущена первая задача")
		ta := time.After(2 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Первая задача завершена по контексту")
			return ctx.Err()
		case <-ta:
			fmt.Println("Первая задача завершена по таймеру")
			return nil
		}
	})

	// Вторая горутина (с ошибкой)
	g.Go(func() error {
		fmt.Println("Запущена вторая задача")
		time.Sleep(1 * time.Second)
		fmt.Println("Вторая задача вызывает ошибку")
		return fmt.Errorf("ошибка во второй задаче") // ошибка
	})

	// Третья горутина
	g.Go(func() error {
		fmt.Println("Запущена третья задача")
		ta := time.After(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Третья задача завершена по контексту")
			return ctx.Err()
		case <-ta:
			fmt.Println("Третья задача завершена по таймеру")
			return nil
		}
	})

	// Ждем завершения всех задач
	if err := g.Wait(); err != nil {
		fmt.Printf("Обнаружена ошибка: %v\n", err)
	}
}
