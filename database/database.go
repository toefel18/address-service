package main

import (
    "database/sql"
    _ "github.com/lib/pq" // import with _ prefixed means importing solely for it's side-effects
    "log"
)

func main() {

    db, err := sql.Open("postgres", "postgres://postgress:root@localhost/addressservice?sslmode=verify-full")
    if err != nil {
        log.Fatal(err)
    }

    age := 21
    rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)


}