package main

import (
	"fmt"
	"net/http"
)

// Middleware
func middlewareFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rp := r.URL.Query()
		if len(rp) == 0 {
			http.Error(w, "Плохой запрос - нет параметров запроса", http.StatusBadRequest)
			return
		}
		fmt.Println("Работа Middleware начата")
		next.ServeHTTP(w, r)
		fmt.Println("Работа Middleware завершена")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All parameters present!"))
}

func main() {

	mux := http.NewServeMux()

	handler := http.HandlerFunc(mainHandler)

	mux.Handle("/", middlewareFunc(handler))

	fmt.Println("Server: http://localhost:8080/?user=John")

	_ = http.ListenAndServe(":8080", mux)
}
