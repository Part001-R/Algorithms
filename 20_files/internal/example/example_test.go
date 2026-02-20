package example

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест добавления записи в файл.
func TestAppendToFile(t *testing.T) {

	fileName := "test.txt"
	infoWr := "123"
	err := AppendToFile(fileName, infoWr)
	require.NoErrorf(t, err, "Неожиданная ошибка 1")

	infoWr = "456"
	err = AppendToFile(fileName, infoWr)
	require.NoErrorf(t, err, "Неожиданная ошибка 1")
}

// Перезапись с использованием Seek.
func TestOverwriteFileSeek(t *testing.T) {

	fileName := "test.txt"
	infoWr := "-AAA"
	err := OverwriteFileSeek(fileName, infoWr)
	require.NoErrorf(t, err, "Неожиданная ошибка")
}

// Вставка строки по номеру.
func TestInsertNewStringByNumb(t *testing.T) {

	t.Run("Корректный номер", func(t *testing.T) {

		// Подготовка
		fileName := "test.txt"
		infoWr := "123"
		err := AppendToFile(fileName, infoWr)
		require.NoErrorf(t, err, "Неожиданная ошибка 1")

		infoWr = "456"
		err = AppendToFile(fileName, infoWr)
		require.NoErrorf(t, err, "Неожиданная ошибка 2")

		defer os.Remove(fileName)

		// Тест
		err = InsertNewStringByNumb(fileName, "================", 1)
		require.NoErrorf(t, err, "Неожиданная ошибка вставки строки: <%w>", err)
	})

	t.Run("Превышен номер", func(t *testing.T) {

		// Подготовка
		fileName := "test2.txt"
		infoWr := "123"
		err := AppendToFile(fileName, infoWr)
		require.NoErrorf(t, err, "Неожиданная ошибка 1")

		infoWr = "456"
		err = AppendToFile(fileName, infoWr)
		require.NoErrorf(t, err, "Неожиданная ошибка 2")

		defer os.Remove(fileName)

		// Тест
		err = InsertNewStringByNumb(fileName, "================", 2)
		assert.Equalf(t, ErrOverNumbStr, err, "Нет соответствия ошибки")
	})
}
