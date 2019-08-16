// go inbuilt package - database/sql can be used with any SQL provider
// but a driver is required for every specific database like MySQL, Postgres

package main

import (
	"database/sql"
	"fmt"
	"stark/gophercises/normalize/normalize"

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

	// create phone_numbers table
	// must(createPhoneTable(db))

	// insert testing data in table
	// _, err = insertTable(db, "1234567890")
	// _, err = insertTable(db, "123 456 7891")
	// _, err = insertTable(db, "(123) 456 7892")
	// _, err = insertTable(db, "(123) 456-7893")
	// _, err = insertTable(db, "123-456-7894")
	// _, err = insertTable(db, "123-456-7890")
	// _, err = insertTable(db, "1234567892")
	// _, err = insertTable(db, "(123)456-7892")

	// insert single row in table
	// id, err := insertTable(db, "1234567890")
	// must(err)
	// fmt.Println("id =", id)

	// get a single row with id
	// num, err := getPhoneNumber(db, 6)
	// must(err)
	// fmt.Println(num)

	// get all rows
	res, err := getAllNumbers(db)
	must(err)

	for _, p := range res {
		fmt.Println(p.value)

		// normalize number
		normNumber := normalize.RegexNormalize(p.value)

		if normNumber != p.value {
			// check if normalized entry is already in the db
			found, err := findPhone(db, normNumber)
			must(err)

			if found != nil {
				fmt.Println("Need to delete this entry")
				err = deleteNumber(db, p.id)
				must(err)
			} else {
				fmt.Println("Need to update this entry")
				p.value = normNumber
				err = updateNumber(db, p)
				must(err)
			}
		} else {
			fmt.Println("No change is required")
		}
	}

	db.Close()
}

type phone struct {
	id    int
	value string
}

func updateNumber(db *sql.DB, p phone) error {
	query := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(query, p.id, p.value)
	return err
}

func deleteNumber(db *sql.DB, id int) error {
	query := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(query, id)
	return err
}

func findPhone(db *sql.DB, num string) (*phone, error) {
	var p phone
	err := db.QueryRow("SELECT id, value FROM phone_numbers WHERE value=$1", num).Scan(&p.id, &p.value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func getAllNumbers(db *sql.DB) ([]phone, error) {
	var res []phone
	query := `SELECT id, value FROM phone_numbers`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p phone
		if err = rows.Scan(&p.id, &p.value); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func getPhoneNumber(db *sql.DB, id int) (string, error) {
	var value string
	query := `SELECT value FROM phone_numbers WHERE id=$1`
	err := db.QueryRow(query, id).Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
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
