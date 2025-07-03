package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var errMy = errors.New("fault adding book CCC")

func main() {

	dbPtr, err := sql.Open("sqlite", "onetoone.db")
	if err != nil {
		log.Fatalf("fault db connect:{%v}", err)
	}

	err = dbPtr.Ping()
	if err != nil {
		log.Fatalf("fault ping db:{%v}", err)
	}

	err = createTables(dbPtr)
	if err != nil {
		log.Fatalf("fault create tables:{%v}", err)
	}

	err = fillTables(dbPtr)
	if err != nil {
		log.Fatalf("fault fill table:{%v}", err)
	}

	// Проверка на уникальность.
	// Добавление книги с повторным использованием id автора
	err = addBookFault(dbPtr)
	if errors.Is(err, errMy) {
		log.Print("UNIQUE detected")
	} else {
		log.Print("UNIQUE not detected")
	}

	fmt.Println("Done")
}

func createTables(db *sql.DB) error {

	q := `CREATE TABLE IF NOT EXISTS authors(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEST NOT NULL,
		ts TIMASTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (first_name, last_name)
	);`
	_, err := db.Exec(q)
	if err != nil {
		return fmt.Errorf("fault create table users:{%v}", err)
	}

	q = `CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTONCREMENT,
		author_fk INTEGER UNIQUE NOT NULL,
		name TEXT NOT NULL,
		
		CONSTRAINT author_fk_profile FOREIGN KEY (author_fk) REFERENCES authors (id)
		ON UPDATE CASCADE
		ON DELETE CASCADE
	)`
	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("fault create table books:{%v}", err)
	}

	return nil
}

func fillTables(db *sql.DB) error {

	q := `INSERT INTO authors (first_name, last_name) VALUES (?, ?)`

	_, err := db.Exec(q, "A", "1")
	if err != nil {
		return fmt.Errorf("fault adding A 1 uathor:{%v}", err)
	}
	_, err = db.Exec(q, "B", "2")
	if err != nil {
		return fmt.Errorf("fault adding B 2 uathor:{%v}", err)
	}

	q = `INSERT INTO books (author_fk, name) VALUES (?, ?)`

	_, err = db.Exec(q, 1, "AAA")
	if err != nil {
		return fmt.Errorf("fault adding book AAA:{%v}", err)
	}

	_, err = db.Exec(q, 2, "BBB")
	if err != nil {
		return fmt.Errorf("fault adding book BBB:{%v}", err)
	}

	return nil

}

func addBookFault(db *sql.DB) error {

	q := `INSERT INTO books (author_fk, name) VALUES (?, ?)`

	_, err := db.Exec(q, 1, "CCC")
	if err != nil {
		return fmt.Errorf("fault adding:{%w}", errMy)
	}

	return nil
}
