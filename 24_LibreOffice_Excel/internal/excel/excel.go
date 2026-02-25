package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Чтение вкладки документа с выводом содержимого в терминал.
func ShowSheet(filePath, sheet string) error {

	// Подключение к файлу.
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %v", err)
	}
	defer f.Close()

	// Чтение строк в листе.
	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("Ошибка при получении строк: %v", err)
	}

	for _, row := range rows {
		for _, cell := range row {
			fmt.Printf("%s\t", cell)
		}
		fmt.Println()
	}

	return nil
}

// Чтение указанной ячейки.
func CellValue(filePath, sheet, cell string) (string, error) {

	// Подключение к файлу
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Ошибка при открытии файла: %v", err)
	}
	defer f.Close()

	// Получение значения ячейки
	value, err := f.GetCellValue(sheet, cell)
	if err != nil {
		return "", fmt.Errorf("Ошибка при получении значения ячейки: %v", err)
	}

	return value, nil
}

// Запись в указанную ячейку
func WriteCellValue(filePath, sheet, cell, value string) error {

	// Подключение к файлу
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %v", err)
	}
	defer f.Close()

	// Запись значения в ячейку
	if err := f.SetCellValue(sheet, cell, value); err != nil {
		return fmt.Errorf("Ошибка при записи значения в ячейку: %v", err)
	}

	// Сохраняем изменения
	if err := f.Save(); err != nil {
		return fmt.Errorf("Ошибка при сохранении файла: %v", err)
	}

	return nil
}

// AddSheet добавляет новый лист в Excel документ
func AddSheet(filePath, sheetName string) error {

	// Подключение к файлу
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %v", err)
	}
	defer f.Close()

	// Добавление нового листа
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("Ошибка при создании нового листа: %v", err)
	}

	// Перемещение нового листа на первое место
	f.SetActiveSheet(index)

	// Сохраняем изменения
	if err := f.Save(); err != nil {
		return fmt.Errorf("Ошибка при сохранении файла: %v", err)
	}

	return nil
}
