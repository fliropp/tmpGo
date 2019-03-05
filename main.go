package main

import (
	"database/sql"
	"fmt"
	"log"
	"html"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

)

type OrdinoDataEntry struct {
	Uid         int64
	Score       float64
	Artid       string
	Segment     string
	Publication string
	Source      string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "axx54TWN"
    dbName := "ORDINO"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
		defer db.Close()

    return db
}



func init() {
	fmt.Println("init...")

}

func main() {
	http.HandleFunc("/ordino", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
		db := dbConn()

		insert, err := db.Query("INSERT INTO ORDINO VALUES(1234,'1234567','www.an.no', 0.5, 'billboard')")

		if err != nil {
				panic(err.Error)
		}

		defer insert.Close()

    log.Fatal(http.ListenAndServe(":8080", nil))
}
