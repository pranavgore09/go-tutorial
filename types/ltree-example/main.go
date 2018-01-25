package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "postgres"
	DB_PORT     = "1234"
)

type TestInstance struct {
	path string
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	db.QueryRow("INSERT INTO test VALUES('x.y.z');")
	// checkErr(err)

	// fmt.Println("# Updating")
	// stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	// checkErr(err)

	// res, err := stmt.Exec("astaxieupdate", lastInsertId)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect, "rows changed")

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)

	for rows.Next() {
		var t TestInstance
		err = rows.Scan(&t.path)
		checkErr(err)
		fmt.Println("path")
		fmt.Printf("%v\n", t.path)
	}

	// fmt.Println("# Deleting")
	// stmt, err = db.Prepare("delete from userinfo where uid=$1")
	// checkErr(err)

	// res, err = stmt.Exec(lastInsertId)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
