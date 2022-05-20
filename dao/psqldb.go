// package dao
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	// "github.com/google/uuid"
	"github.com/gofrs/uuid"
    "net/http"
	"log"
	// "os"
	
)

// config postgressql DB
const (
    host     = "localhost"
    port     = 5432
    user     = "kriti"
    password = "nkx01"
    dbname   = "go_dummy"
)


func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
// sales=# CREATE TABLE leads (id INTEGER PRIMARY KEY, name VARCHAR);


func NewPsqlUserDao(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlconn)

	CheckError(err)
    return db, err
}

// func findNextID()

func main() {
	fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
	db, err := NewPsqlUserDao(host, port, user, password, dbname)

    CheckError(err)
 
    defer db.Close()
	
	uuid := uuid.Must(uuid.NewV4()).String() //uuid.New()
	fmt.Printf(uuid)

    // // query_id := `SELECT id FROM user_table;`
	// ids, e := db.Query("SELECT id FROM user_table;")
    // // rows, e := db.Exec(query_id)

	// CheckError(e)

	// // Check for errors from iterating over rows.
	// if ids.Next() == false {
	// 	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	// 	_, e = db.Exec(insertDynStmt, "Jane", 2)
	// 	CheckError(e)
	// }

	// for ids.Next() {
	// 	e := ids.Scan(uuid);
	// 	if e!=nil{
	// 		fmt.Printf(uuid)
	// 	} else {
	// 		fmt.Println("===")
	// 	}
	// }



    // insert
    // hardcoded
    // insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
    // _, e := db.Exec(insertStmt)
    // CheckError(e)

}






 // go mod init sql_inject
 // go run .
//  psql -h localhost -d go_dummy -U kriti -p 5432
// sudo service postgresql restart
// sudo -u postgres psql go_dummy
// https://stackoverflow.com/a/12670521/6660373
