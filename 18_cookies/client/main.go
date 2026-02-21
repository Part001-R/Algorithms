package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {

	url := "http://localhost:8080"

	// Запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Первая куки
	cookie1 := &http.Cookie{
		Name:  "example_cookie",
		Value: "cookie_value",
	}
	req.AddCookie(cookie1)

	// Вторая куки
	cookie2 := &http.Cookie{
		Name:    "example_cookie_2",
		Value:   "cookie_value_2",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}
	req.AddCookie(cookie2)

	// Выполнение запроса
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка: статус %d", resp.StatusCode)
	}

	fmt.Println("Запрос выполнен успешно!")
}

// Клиент с прокси
func newClientWithProxy() (*http.Client, error) {
	proxyURL := os.Getenv("HTTP_PROXY")
	if proxyURL == "" {
		return nil, errors.New("HTTP_PROXY environment variable is not set")
	}

	proxyURLParsed, err := url.Parse(proxyURL)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy URL: %w", err)
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURLParsed),
		},
		Timeout: 10 * time.Second,
	}, nil
}
