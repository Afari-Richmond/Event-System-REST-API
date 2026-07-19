package main 



import (
	"database/sql"
	"log"
)

func main () {
	db, err := sql.Open("sqlite3", "./datadb")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}