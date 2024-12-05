package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", "./SqlGoExplore.db")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY, name TEXT)")
	// statement.Exec()
	if err == nil {
		statement.Exec()
	} else {
		fmt.Println(err)
	}
}
