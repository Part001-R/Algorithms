package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {

	dbPtr, err := sql.Open("sqlite", "onetomany.db")
	if err != nil {
		log.Fatal("fault connect DB")
	}
	defer dbPtr.Close()

	err = dbPtr.Ping()
	if err != nil {
		log.Fatal("fault ping DB")
	}

	err = createTables(dbPtr)
	if err != nil {
		log.Fatalf("fault creat tables: {%v}", err)
	}

	err = fillTables(dbPtr)
	if err != nil {
		log.Fatalf("fault fill tables:{%v}", err)
	}

	err = addBook(dbPtr)
	if err != nil {
		log.Printf("fault:{%v}", err)
	} else {
		log.Println("Ok")
	}

	fmt.Println("Done")
}

func createTables(db *sql.DB) error {
	q := `CREATE TABLE IF NOT EXISTS authors(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (first_name, last_name)
	)`

	_, err := db.Exec(q)
	if err != nil {
		return fmt.Errorf("fault create authors table: {%w}", err)
	}

	// authors_fk не UNIQUE, что и реализует схему - один ко многим
	q = `CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		authors_fk INTEGER NOT NULL,  
		name TEXT NOT NULL,  
		ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

		CONSTRAINT authors_fk_c FOREIGN KEY (authors_fk) REFERENCES authors (id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
	)`

	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("fault create books table: {%w}", err)
	}

	return nil
}

func fillTables(db *sql.DB) error {

	q := `INSERT INTO authors (first_name, last_name) VALUES (?, ?)`

	_, err := db.Exec(q, "AAA", "aaa")
	if err != nil {
		return fmt.Errorf("fault adding AAA aaa:{%w}", err)
	}
	_, err = db.Exec(q, "BBB", "bbb")
	if err != nil {
		return fmt.Errorf("fault adding BBB bbb:{%w}", err)
	}

	q = `INSERT INTO books (authors_fk, name) VALUES (?, ?)`

	_, err = db.Exec(q, 1, "AaAaAa")
	if err != nil {
		return fmt.Errorf("fault adding book AaAaAa:{%w}", err)
	}
	_, err = db.Exec(q, 2, "BbBbBb")
	if err != nil {
		return fmt.Errorf("fault adding book BbBbBb:{%w}", err)
	}

	return nil
}

func addBook(db *sql.DB) error {

	q := `INSERT INTO books (authors_fk, name) VALUES (?, ?)`

	_, err := db.Exec(q, 1, "CcCcCc") // Книга с FK уже сесть! Тут реализация один ко многим
	if err != nil {
		return fmt.Errorf("fault adding book CcCcCc:{%w}", err)
	}

	return nil
}
