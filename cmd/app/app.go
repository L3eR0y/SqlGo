package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", "./Databases/SqlGoExplore.db")

	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS products (
																	id INTEGER PRIMARY KEY,
																	name TEXT,
																	proteins FLOAT,
																	fats FLOAT,
																	carbohydrates FLOAT
																)`)

	if err == nil {
		statement.Exec()
	} else {
		fmt.Println(err)
	}
}
