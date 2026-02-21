package main

import (
	"e/session/server/internal/controller"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var addrServer = "localhost:50100"

func main() {

	// Конструктор
	instServer := controller.New(addrServer)

	// Роутер
	cr := chi.NewRouter()
	cr.Group(func(r chi.Router) {

		// Миддлвара.
		r.Use(instServer.Middleware)

		// Регистрация пользователя.
		r.Post("/registration", instServer.Registration)
		// Аутентификация пользователя.
		r.Post("/authentication", instServer.Authentication)
		// Получеение информации.
		r.Get("/info", instServer.Info)
	})

	// Запуск сервера.
	srv := &http.Server{
		Addr:    instServer.Addr,
		Handler: cr,
	}

	log.Printf("Сервер запускается на: <%s>\n", instServer.Addr)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка в работе сервера: <%v>\n", err)
	}
}
