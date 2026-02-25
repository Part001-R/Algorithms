package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

// Генерация ключей
func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Ошибка генерации приватного ключа: %v", err)
	}
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey
}
