// package main

// import (
//   "database/sql"
//   "fmt"

//   _ "github.com/lib/pq"
// )

// const (
//   host     = "localhost"
//   port     = 5432
//   user     = "postgres"
//   password = "your-password"
//   dbname   = "calhounio_demo"
// )

// func main() {
//   psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//     "password=%s dbname=%s sslmode=disable",
//     host, port, user, password, dbname)
//   db, err := sql.Open("postgres", psqlInfo)
//   if err != nil {
//     panic(err)
//   }
//   defer db.Close()

//   sqlStatement := `
// INSERT INTO users (age, email, first_name, last_name)
// VALUES ($1, $2, $3, $4)
// RETURNING id`
//   id := 0
//   err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
//   if err != nil {
//     panic(err)
//   }
//   fmt.Println("New record ID is:", id)
// }

package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "kriti"
    password = "nkx01"
    dbname   = "go_dummy"
)
 
func main() {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
 
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
 
    defer db.Close()
 
    // insert
    // hardcoded
    insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
    _, e := db.Exec(insertStmt)
    CheckError(e)
 
    // dynamic
    insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
    _, e = db.Exec(insertDynStmt, "Jane", 2)
    CheckError(e)
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

// package main

// import "fmt"

// func main() {
//     fmt.Printf("Hello, World\n")
// }

// go run .