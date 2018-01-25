package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	id       int
	name     string
	location Ltree
}

type Ltree []string

func (p Ltree) Value() (driver.Value, error) {
	var op []string
	for _, x := range p {
		op = append(op, x)
	}
	return strings.Join(op, "."), nil
}

func (p *Ltree) Scan(value interface{}) error {
	if value == nil {
		*p = []string{}
		return nil
	}
	if str, err := driver.String.ConvertValue(value); err == nil {
		full := ""
		for _, x := range str.([]uint8) {
			full += string(x)
		}
		*p = strings.Split(full, ".")
		return nil
	}
	return errors.New("Unable to scan Ltree")
}

func panicIfError(err error, msg string) {
	if err != nil {
		panic(err)
	}
	fmt.Println("LOG : ", msg)
}

func setupDB(db *sql.DB) {
	// create table if not exists
	dropTable := "drop table if exists person;"
	_, err := db.Exec(dropTable)
	panicIfError(err, "Drop table done.")

	createTable := `CREATE TABLE IF NOT EXISTS person(
		id TEXT NOT NULL PRIMARY KEY,
		name text,
		location ltree);`
	_, err = db.Exec(createTable)
	panicIfError(err, "Create table done.")
}

func insertRecord(db *sql.DB, p *person) {
	insertPerson := "insert into person (id, name, location) values (?, ?, ?);"
	_, err := db.Exec(insertPerson, p.id, p.name, p.location)
	panicIfError(err, "Insert done.")
}

func readAllRecords(db *sql.DB) {
	selectQuery := "select * from person;"
	rows, err := db.Query(selectQuery)
	panicIfError(err, "Rows reading done.")
	defer rows.Close()
	for rows.Next() {
		p := new(person)
		e := rows.Scan(&p.id, &p.name, &p.location)
		panicIfError(e, "reading row done.")
		fmt.Println(p)
	}
}
func main() {
	fmt.Println("okay")
	db, err := sql.Open("sqlite3", "./db")
	panicIfError(err, "DB open done.")
	defer db.Close()
	setupDB(db)

	p1 := person{
		id:       1,
		name:     "pranav",
		location: []string{"india", "karnataka", "banglaore"},
	}
	p2 := person{
		id:       2,
		name:     "pranav",
		location: []string{"india", "maharashta", "pune"},
	}
	p3 := person{
		id:       3,
		name:     "pranav",
		location: []string{"india", "maharashta", "mumbai"},
	}
	insertRecord(db, &p1)
	insertRecord(db, &p2)
	insertRecord(db, &p3)
	readAllRecords(db)
}
