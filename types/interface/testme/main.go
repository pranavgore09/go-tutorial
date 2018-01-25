package main

import "database/sql"
import "fmt"
import _ "github.com/lib/pq"

type person struct {
	name     string
	location string // Asia.India.Bangalore.Banerghatta
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "postgres"
	DB_PORT     = "1234"
)

func panicIfError(err error, msg string) {
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
}

func setupDB(db *sql.DB) {
	dropTable := "drop table if exists person"
	_, err := db.Exec(dropTable)
	panicIfError(err, "Drop table done.")

	createTable := `CREATE TABLE IF NOT EXISTS person(
		name text,
		location text);`
	_, err = db.Exec(createTable)
	panicIfError(err, "Create table done.")
}

func insertRecord(db *sql.DB, p person) {
	insertPerson := "insert into person VALUES ( 'hello', 'world');"
	_, err := db.Exec(insertPerson)
	panicIfError(err, "Insert done.")
}

func printAllRecords(db *sql.DB) {
	selectPerson := "select * from person;"
	rows, err := db.Query(selectPerson)
	panicIfError(err, "Rows fetched.")
	defer rows.Close()
	for rows.Next() {
		p := new(person)
		e := rows.Scan(&p.name, &p.location)
		panicIfError(e, "reading row done.")
		fmt.Println(p)
	}
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := sql.Open("postgres", dbinfo)
	panicIfError(err, "")
	defer db.Close()

	setupDB(db)

	p1 := person{
		name:     "Lala",
		location: "Europe.Germany.Heisenberg",
	}
	insertRecord(db, p1)
	printAllRecords(db)
}
