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
  password = "root"
  dbname   = "Scrapydata1"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

//   sqlStatement := `
// INSERT INTO users (age, email, first_name, last_name)
// VALUES ($1, $2, $3, $4)
// RETURNING id`
//   id := 0
//   err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)

//   fmt.Println(id)
//   if err != nil {
//     panic(err)
//   }

rows, err := db.Query("SELECT id, title, amenities,location,price FROM newdata")
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var id int
    var title string
    var amenities string
    var location string
    var price string
    err = rows.Scan(&id, &title, &amenities,&location,&price)
    if err != nil {
      // handle this error
      panic(err)
    }
    // fmt.Println(id, title,amenities,location,price)
    fmt.Println(id, title,amenities,location,price)
  }
  // get any error encountered during iteration
  err = rows.Err()
  if err != nil {
    panic(err)
  }
  
}