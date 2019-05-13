package logs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//pointer of database
var db *sql.DB

const (                      //数据库登入信息
    host     = "localhost"
    port     =  5432
    user     = "Dong"
    password = "87257745"
    dbname   = "management"
)


//connect to testing database 
func testconnect() {
	var err error
	//if don't add 'sllmod=disable' there may export error "pq: SSL is not enabled on the server"
	// confstr := fmt.Sprintf(
	// 	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	"129.204.193.192",
	// 	5432,
	// 	"logmoduleuser",
	// 	"itispassword",
	// 	"logmodule",
	// )
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err = sql.Open("postgres", psqlInfo)
	//db, err = sql.Open("postgres", confstr)
	err = db.Ping()
	if err!= nil {
		panic(err)
	}
	//pointer of logfile still not init, you can not use logs.Println yet
	fmt.Println("Database Connect Scuess!")
}
