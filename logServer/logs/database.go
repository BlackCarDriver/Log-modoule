package logs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

//pointer of database
var db *sql.DB

func init(){
   connect1()                    
}

//connect to remote server database
func connect1() {
	var err error
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
		fmt.Println("database connect fall :",err)
		os.Exit(1)
	}
	fmt.Println("Database Connect Scuess!")
}

//connect to localhost database
func connect2(){
	host     := "localhost"
    port     :=  5432
    user     := "Dong"
    password := "87257745"
	dbname   := "management"
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err!= nil {
		fmt.Println("database connect fall :",err)
		os.Exit(1)
	}
	fmt.Println("Database Connect Scuess!")
}