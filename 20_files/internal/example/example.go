package example

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Запись в файл. Возвращается ошибка.
//
// Параметры:
//
//	fileName - имя файла.
//	infoWr - информация для записи.
func AppendToFile(fileName, infoWr string) (err error) {

	// Формирование полного пути к файлу.
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Ошибка получения рабочей директории: <%w>\n", err)
	}
	fullFileName := filepath.Join(dir, fileName)

	// Подключение к файлу.
	file, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("Ошибка открытия файла: <%w>", err)
	}
	defer func() {
		err = file.Close()
	}()

	// Запись.
	_, err = file.WriteString(infoWr + "\n")
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл")
	}

	return nil
}

// Перезапись содержимого файла. Возвращается ошибка.
//
// Параметры:
//
//	fileName - имя файла.
//	infoWr- данные для записи.
func OverwriteFileSeek(fileName, infoWr string) (err error) {

	// Формирование полного пути к файлу.
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Ошибка получения рабочей директории: <%w>\n", err)
	}
	fullFileName := filepath.Join(dir, fileName)

	// Открытие файла.
	f, err := os.OpenFile(fullFileName, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	f.Seek(3, io.SeekStart) // Установка указателя чтения/записи в начало файла c смещение позиции

	// Запись в файл
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "%s\n", infoWr)
	w.Flush() // Передача содержимого буфера на диск

	return nil
}

// Добавление новой строки в указанный номер строки. Возвращается ошибка.
//
// Параметры:
//
//	fileName - имя файла.
//	infoWr- данные для записи.
//	numbStr - номер строки для вставки.
func InsertNewStringByNumb(fileName, infoWr string, numbStr int) (err error) {

	// Формирование полного пути.
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Ошибка при получении директории: <%w>", err)
	}
	fullFileName := filepath.Join(dir, fileName)

	// Подключение к файлу.
	file, err := os.OpenFile(fullFileName, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к файлу: <%w>", err)
	}
	defer func() {
		_ = file.Close()
	}()

	// Чтение файла.
	scan := bufio.NewScanner(file)
	rStr := make([]string, 0)
	for scan.Scan() {
		rStr = append(rStr, scan.Text())
	}
	if err := scan.Err(); err != nil {
		return fmt.Errorf("Ошибка сканера: <%w>", err)
	}

	// Проверка вхождения номера строки.
	cntStr := len(rStr)
	if numbStr >= cntStr {
		return ErrOverNumbStr
	}

	// Формирование данных для записи.
	wStr := make([]string, 0)
	for i, v := range rStr {
		if i == numbStr {
			wStr = append(wStr, infoWr)
		}
		wStr = append(wStr, v)
	}

	// Переход в начало файла и обрезание его
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("Ошибка при перемещении указателя файла: <%w>", err)
	}
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("Ошибка при обрезке файла: <%w>", err)
	}

	// Запись данных в файл
	writer := bufio.NewWriter(file)
	for _, line := range wStr {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("Ошибка записи строки в файл: <%w>", err)
		}
	}
	if err := writer.Flush(); err != nil {
		file.Close()
		return fmt.Errorf("Ошибка сохранения данных: <%w>", err)
	}

	return nil

}
