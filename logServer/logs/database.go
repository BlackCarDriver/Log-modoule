package logs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//pointer of database
var db *sql.DB

func testconnect() {
	var err error
	//if don't add 'sllmod=disable' there may export error "pq: SSL is not enabled on the server"
	confstr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"129.204.193.192",
		5432,
		"logmoduleuser",
		"itispassword",
		"logmodule",
	)
	db, err = sql.Open("postgres", confstr)
	err = db.Ping()
	if err!= nil {
		panic(err)
	}
	Println("Database Connect Scuess!")
}
