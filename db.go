package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
)

func initDB() *sql.DB {
	//get database variables
	DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS := getDBConfig()
	//build connection detail string
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	//try to establish database connection
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		printErr(err, "Unable to establish database connection!")
	}

	err = db.Ping()
	if err != nil {
		printErr(err, "Database isn't reachable!")
	}

	return db
}

func getDBConfig() (DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS string) {
	//retrieve enviromental variables
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")

	return DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS
}

func insertQuote(quote string) (*Quote, error) {
	db := initDB()
	defer db.Close()

	quote = strings.Replace(quote, `'`, `"`, -1)
	quote = strings.Replace(quote, "`", `'`, -1)

	query := fmt.Sprintf(`INSERT INTO quotes("text") VALUES('%s') RETURNING *`, quote)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var text string
		err = rows.Scan(&id, &text)
		if err != nil {
			return nil, err
		}

		return &Quote{ID: id, Text: text}, nil
	}
	err = errors.New("Error occured when inserting")
	return nil, err
}

func retrieveQuotes() ([]*Quote, error) {
	db := initDB()
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM quotes`)
	if err != nil {
		return nil, err
	}

	var quoteList []*Quote

	for rows.Next() {
		var id int
		var text string
		err = rows.Scan(&id, &text)
		if err != nil {
			return nil, err
		}

		quote := Quote{ID: id, Text: text}
		quoteList = append(quoteList, &quote)
	}

	return quoteList, nil

}
