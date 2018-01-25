// docker run --name test-psql-docker -e POSTGRES_PASSWORD=mysecretpassword -d -p 1234:5432 postgres
// CREATE TABLE person (name text,location text);

// CREATE TABLE personInfo (name text,location ltree);
// CREATE INDEX location_gist_idx ON personInfo USING gist(path);

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type person struct {
	name     string
	location string // Asia.India.Bangalore.Banerghatta
}

const (
	DBUser          = "postgres"
	DBPassword      = "mysecretpassword"
	DBName          = "postgres"
	DBPort          = "1234"
	CreateTableOld  = "CREATE TABLE IF NOT EXISTS person (name text,location text);"
	CreateTable     = "CREATE TABLE IF NOT EXISTS person (name text,location ltree);"
	DropTable       = "DROP TABLE IF Exists person;"
	CreateIndex     = "CREATE INDEX location_gist_idx ON person USING gist(location);"
	CreateExtension = "CREATE EXTENSION IF NOT EXISTS \"ltree\";"
)

func panicIfError(err error, msg string) {
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
}

func setupDB(db *sql.DB) {
	_, err := db.Exec(DropTable)
	panicIfError(err, "Table Dropped.")

	// _, err = db.Exec(CreateTableOld)
	// panicIfError(err, "Table created.")

	_, err = db.Exec(CreateExtension)
	panicIfError(err, "Extension created.")

	_, err = db.Exec(CreateTable)
	panicIfError(err, "Table created.")

	_, err = db.Exec(CreateIndex)
	panicIfError(err, "Index created.")
}

func insertRecord(db *sql.DB, p person) {
	insertPerson := "insert into person VALUES ( $1, $2 );"
	_, err := db.Exec(insertPerson, p.name, p.location)
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
		DBUser, DBPassword, DBName, DBPort)
	db, err := sql.Open("postgres", dbinfo)
	panicIfError(err, "")
	defer db.Close()

	setupDB(db)

	p1 := person{
		name:     "RH dev",
		location: "Europe.Germany.Heisenberg",
	}
	p2 := person{
		name:     "RH dev 2",
		location: "Asia.India.Bangalore",
	}
	p3 := person{
		name:     "RH dev 3",
		location: "Europe.England",
	}
	insertRecord(db, p1)
	insertRecord(db, p2)
	insertRecord(db, p3)
	printAllRecords(db)
}

// type person struct {
// 	name     string
// 	location Path // Asia.India.Bangalore.Banerghatta
// }
// func main() {
// 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		DBUser, DBPassword, DBName, DBPort)
// 	db, err := sql.Open("postgres", dbinfo)
// 	panicIfError(err, "")
// 	defer db.Close()

// 	setupDB(db)

// 	p1 := person{
// 		name:     "RH dev",
// 		location: []string{"Europe", "Germany", "Heisenberg"},
// 	}
// 	p2 := person{
// 		name:     "RH dev 2",
// 		location: []string{"Asia", "India", "Bangalore"},
// 	}
// 	p3 := person{
// 		name:     "RH dev 3",
// 		location: []string{"Europe", "England"},
// 	}
// 	insertRecord(db, p1)
// 	insertRecord(db, p2)
// 	insertRecord(db, p3)
// 	printAllRecords(db)
// }

// // Define own type and implement Valuer and Scanner interfaces

// type Path []string

// func (p Path) Value() (driver.Value, error) {
// 	var op []string
// 	for _, x := range p {
// 		op = append(op, x)
// 	}
// 	return strings.Join(op, "."), nil
// }

// func (p *Path) Scan(value interface{}) error {
// 	if value == nil {
// 		*p = []string{}
// 		return nil
// 	}
// 	if str, err := driver.String.ConvertValue(value); err == nil {
// 		full := ""
// 		for _, x := range str.([]uint8) {
// 			full += string(x)
// 		}
// 		*p = strings.Split(full, ".")
// 		return nil
// 	}
// 	return errors.New("Unable to scan Ltree")
// }
