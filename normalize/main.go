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
	db, err := sql.Open("postgres", psqlInfo)
	must(err)

	// create a db if doesn't exist
	err = resetDB(db, dbname)
	must(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)

	must(db.Ping())
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
