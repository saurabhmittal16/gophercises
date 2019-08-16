// go inbuilt package - database/sql can be used with any SQL provider
// but a driver is required for every specific database like MySQL, Postgres

package main

import (
	"database/sql"
	"fmt"

	// unused import, just importing the package calls the init function
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "saurabh"
	password = "1234"
	dbname   = "gophercises_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	// open connection to postgres
	// db, err := sql.Open("postgres", psqlInfo)
	// must(err)

	// create a db if doesn't exist
	// err = resetDB(db, dbname)
	// must(err)
	// db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)

	must(createPhoneTable(db))
	id, err := insertTable(db, "1234567890")
	must(err)
	fmt.Println("id =", id)

	db.Close()
}

func insertTable(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createPhoneTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)
	`
	_, err := db.Exec(statement)
	return err
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}

	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}

	return nil
}
