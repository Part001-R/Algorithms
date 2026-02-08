package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Принят запрос")

	// Чтение cookies
	cookies := r.Cookies()
	fmt.Printf("Количество куки: <%d>\n", len(cookies))

	for _, cookie := range cookies {
		fmt.Printf("Имя: %s, Значение: %s\n", cookie.Name, cookie.Value)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
