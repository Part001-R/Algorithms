package controller

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

// Представление контроллера.
type Cntr struct {
	regHash map[string]string // Хэши регистрации
	Addr    string            // Адрес сервера
	privKey *rsa.PrivateKey   // Приватный ключ
	pubKey  *rsa.PublicKey    // Публичный ключ
}

// Конструктор.
func New(addr string) *Cntr {

	// Генерация ключей
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Ошибка генерации приватного ключа: %v", err)
	}
	publicKey := &privateKey.PublicKey

	// Экземпляр
	return &Cntr{
		regHash: make(map[string]string),
		Addr:    addr,
		privKey: privateKey,
		pubKey:  publicKey,
	}
}
