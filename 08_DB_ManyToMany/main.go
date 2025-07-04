package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {

	dbPtr, err := sql.Open("sqlite", "manytomany.db")
	if err != nil {
		log.Fatalf("fault connect db:{%v}", err)
	}
	err = dbPtr.Ping()
	if err != nil {
		log.Fatalf("fault ping db:{%v}", err)
	}
	err = createTables(dbPtr)
	if err != nil {
		log.Fatalf("fault create tables:{%v}", err)
	}
	err = fillTable(dbPtr)
	if err != nil {
		log.Fatalf("fault fill tables:{%v}", err)
	}

	// Author is exist
	first_name, last_name := "AAA", "aaa"
	sl, err := getBooksByAuthor(dbPtr, first_name, last_name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			log.Printf("empty request result for: {%s %s}", first_name, last_name)
			return
		default:
			log.Fatalf("error exeqution request:{%s}", err)
		}
	}
	fmt.Printf("%v\n", sl)

	// Author is not exist
	first_name, last_name = "FFF", "fff"
	sl, err = getBooksByAuthor(dbPtr, first_name, last_name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			log.Printf("empty request result for: {%s %s}", first_name, last_name)
			return
		default:
			log.Fatalf("error exeqution request:{%s}", err)
		}
	}

	fmt.Printf("%v\n", sl)

}

func createTables(db *sql.DB) error {

	q := `CREATE TABLE IF NOT EXISTS authors(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name VARCHAR(20) NOT NULL,
		last_name VARCHAR(20) NOT NULL,
		ts TIMASTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(q)
	if err != nil {
		return fmt.Errorf("fault create authors table:{%w}", err)
	}

	q = `CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(20) NOT NULL,
		ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("fault create books table:{%w}", err)
	}

	q = `CREATE TABLE IF NOT EXISTS authors_books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author_fk INTEGER NOT NULL,
		book_fk INTEGER NOT NULL,
		ts TIMASTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (author_fk, book_fk),

		CONSTRAINT author_fk_c FOREIGN KEY (author_fk) REFERENCES authors (id)
			ON DELETE CASCADE
			ON UPDATE CASCADE

		CONSTRAINT book_fk_c FOREIGN KEY (book_fk) REFERENCES books (id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	)`
	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("fault create author_book table:{%w}", err)
	}

	return nil
}

func fillTable(db *sql.DB) error {

	q := `INSERT INTO authors (first_name, last_name) VALUES (?, ?)`
	_, err := db.Exec(q, "AAA", "aaa")
	if err != nil {
		return fmt.Errorf("faul add (AAA aaa):{%v}", err)
	}
	_, err = db.Exec(q, "BBB", "bbb")
	if err != nil {
		return fmt.Errorf("faul add (BBB bbb):{%v}", err)
	}
	_, err = db.Exec(q, "CCC", "ccc")
	if err != nil {
		return fmt.Errorf("faul add (CCC ccc):{%v}", err)
	}

	q = `INSERT INTO books (name) VALUES (?)`
	_, err = db.Exec(q, "book_1")
	if err != nil {
		return fmt.Errorf("fault add book (book_1):{%v}", err)
	}
	_, err = db.Exec(q, "book_2")
	if err != nil {
		return fmt.Errorf("fault add book (book_2):{%v}", err)
	}
	_, err = db.Exec(q, "book_3")
	if err != nil {
		return fmt.Errorf("fault add book (book_3):{%v}", err)
	}

	q = `INSERT INTO authors_books (author_fk, book_fk) VALUES (?, ?)`
	_, err = db.Exec(q, 1, 1)
	if err != nil {
		return fmt.Errorf("fault add (1, 1):{%v}", err)
	}
	_, err = db.Exec(q, 1, 2)
	if err != nil {
		return fmt.Errorf("fault add (1, 2):{%v}", err)
	}
	_, err = db.Exec(q, 1, 3)
	if err != nil {
		return fmt.Errorf("fault add (1, 3):{%v}", err)
	}
	_, err = db.Exec(q, 2, 1)
	if err != nil {
		return fmt.Errorf("fault add (2, 1):{%v}", err)
	}
	_, err = db.Exec(q, 3, 1)
	if err != nil {
		return fmt.Errorf("fault add (3, 1):{%v}", err)
	}

	return nil
}

func getBooksByAuthor(db *sql.DB, first_name, last_name string) ([]string, error) {

	sl := make([]string, 0)

	q := `SELECT books.name
          FROM authors
          JOIN authors_books ON authors.id = authors_books.author_fk
          JOIN books ON authors_books.book_fk = books.id
          WHERE authors.first_name = ? AND authors.last_name = ?`

	rows, err := db.Query(q, first_name, last_name)
	if err != nil {
		return nil, fmt.Errorf("fault request:{%w}", err)
	}

	for rows.Next() {

		var num string

		err := rows.Scan(&num)
		if err != nil {
			return nil, fmt.Errorf("fault scan:{%w}", err)
		}

		sl = append(sl, num)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("fault scan 2:{%w}", err)
	}

	return sl, nil
}
