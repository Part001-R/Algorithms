package main

import (
	"e/internal/excel"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var fileName = "test.xlsx"
var sheetName = "Лист1"

func main() {

	// Формирование полного пути.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка определения рабочей директории: <%v>", err)
	}

	fullName := filepath.Join(dir, fileName)

	//
	// Вывод содержимого вкладки в терминал.
	//
	fmt.Println("Вывод содержимого в терминал ...")

	err = excel.ShowSheet(fullName, sheetName)
	if err != nil {
		log.Printf("Функция excel.ReadSheet, вернула ошибку: <%v>\n", err)
	}
	fmt.Println()

	//
	// Чтение ячейки.
	//
	fmt.Println("Чтение содержимого ячейки ...")

	cell := "A1"
	rxValue, err := excel.CellValue(fullName, sheetName, cell)
	if err != nil {
		log.Printf("Функция excel.GetCellValue, вернула ошибку: <%v>\n", err)
	}

	fmt.Printf("В ячейке <%s>, содержится <%s>\n", cell, rxValue)
	fmt.Println()

	//
	// Запись в ячейку.
	//
	fmt.Println("Запись в ячейку ...")

	cell = "A1"
	val := "AAA"
	err = excel.WriteCellValue(fullName, sheetName, cell, val)
	if err != nil {
		log.Printf("Функция excel.WriteCellValue, вернула ошибку: <%v>\n", err)
	}

	rxValue, err = excel.CellValue(fullName, sheetName, cell)
	if err != nil {
		log.Printf("Функция excel.GetCellValue, вернула ошибку: <%v>\n", err)
	}

	fmt.Printf("В ячейке <%s>, содержится <%s>\n", cell, rxValue)
	fmt.Println()

	//
	// Добавление нового листа.
	//
	fmt.Println("Добавление нового листа ...")

	shName := time.Now().Format("05.000")

	err = excel.AddSheet(fullName, shName)
	if err != nil {
		log.Printf("Функция excel.AddSheet, вернула ошибку: <%v>\n", err)
	}

	fmt.Println()
}
