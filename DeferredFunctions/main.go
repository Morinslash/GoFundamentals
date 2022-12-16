package main

import (
	"database/sql"
	"fmt"
)

func main() {
	//SimpleDefer()
	DatabaseExample()
}

func DatabaseExample() {
	db, _ := sql.Open("driverName", "connection string")
	defer db.Close() // => go executes deffer statements as stack, do this will actually close db connection second

	rows, _ := db.Query("sql query here")
	defer rows.Close() // => go will call this function first, then close db connection
}

func SimpleDefer() {
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")
}
