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
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
			fmt.Println("Первая задача завершена")
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
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
			fmt.Println("Третья задача завершена")
			return nil
		}
	})

	// Ждем завершения всех задач
	if err := g.Wait(); err != nil {
		fmt.Printf("Обнаружена ошибка: %v\n", err)
	}
}
