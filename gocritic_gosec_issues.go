package main

import (
	"database/sql"
	"fmt"
	"os"
)

func foo(arg string) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM foo WHERE name = " + arg)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func bar(arg string) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	query := fmt.Sprintf("select * from foo where name = '%s'", arg)

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func main() {
	foo("foo")
	bar("bar")
}
