package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Создание хеша. Возвращается хэш и ошибка.
//
// Параметры:
//
//	data... - вариативные данные.
func createHash(data ...string) (string, error) {

	// Проверка.
	if len(data) < 2 {
		return "", ErrMissingData
	}

	// Объединение строк.
	var b strings.Builder
	for _, v := range data {
		b.WriteString(v)
	}
	bStr := b.String()

	// Вычисление хэша.
	h := sha256.New()
	_, err := h.Write([]byte(bStr))
	if err != nil {
		return "", fmt.Errorf("Ошибка создания хэша: <%w>", err)
	}
	hBytes := h.Sum(nil)

	return hex.EncodeToString(hBytes), nil
}
