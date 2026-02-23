package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	original := "Foo"

	// Кодирование.
	encoded := base64.StdEncoding.EncodeToString([]byte(original))

	// Декодирование.
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	decoded := string(decodedBytes)

	// Проверка.
	if original != decoded {
		log.Printf("Нет соответствия\n")
		return
	}

	log.Println("Работа завершена успешно.")
}
