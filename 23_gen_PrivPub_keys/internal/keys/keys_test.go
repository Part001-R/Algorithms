package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

// Тест генерации ключей.
func TestGenerateKeys(t *testing.T) {

	t.Run("Успешное создание", func(t *testing.T) {

		priv, pub := GenerateKeys()
		require.NotNil(t, priv, "Нет указателя на приватный ключ")
		require.NotNil(t, pub, "Нет указателя на публичный ключ")
	})

	t.Run("Шифрование и расшифрование", func(t *testing.T) {
		priv, pub := GenerateKeys()

		msg := []byte("Секретное сообщение")
		hash := sha256.New()

		excText, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
		require.NoError(t, err, "Ошибка при шифровании сообщения")

		decrText, err := rsa.DecryptOAEP(hash, rand.Reader, priv, excText, nil)
		require.NoError(t, err, "Ошибка при расшифровании сообщения")

		require.Equal(t, msg, decrText, "Расшифрованное сообщение должно совпадать с исходным")
	})
}
