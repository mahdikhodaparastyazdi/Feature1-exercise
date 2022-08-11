package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "first_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	insertstmt := `insert into "Employee("Name", "EmpId") values("rohit", 21)`
	_, e := db.Exec(insertstmt)
	CheckError(e)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
