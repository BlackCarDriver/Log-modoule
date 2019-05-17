package logs

import (
	"log"
    "fmt"
    _ "github.com/lib/pq"
    //"encoding/json"
)

type opelog struct {
	Index           int    `json:"index"`
	Logid           int    `json:"logid"`
    Types        	string    `json:"type"`
    Operator        string    `json:"operator"`
    Time            string    `json:"time"`
    Operation    	string    `json:"operation"`
}

type pages struct {
	Pages           int     `json:"pages`
}

func Records(operator string, operation string){       //添加日志(登录)
	log.Println("正在添加日志...")
    rows, err := db.Prepare("insert into t_opelog (operator,operation) values($1,$2)")
    checkErr(err)
    _,err = rows.Exec(operator, operation)
    checkErr(err)
    log.Println("日志添加成功！")
}

func Getlogs()[]opelog{
    var Opelog []opelog
    if nowsearch!="" {
        rows, err := db.Query("select logid,type,operation,time,operation from oresql where type=$2",nowsearch) 
        checkErr(err)
        for rows.Next(){
		var opelog opelog
		err = rows.Scan(&opelog.Logid,&opelog.Types,&opelog.Operator,&opelog.Time,&opelog.Operation)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        Opelog=append(Opelog,opelog)
	}
    }else {
        rows, err := db.Query("select logid,type,operation,time,operation from t_opelog") 
        checkErr(err)
        for rows.Next(){
		var opelog opelog
		err = rows.Scan(&opelog.Logid,&opelog.Types,&opelog.Operator,&opelog.Time,&opelog.Operation)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        Opelog=append(Opelog,opelog)
	}
    }
    fmt.Println("正在读取数据库")
    return Opelog
}

// func Totalpages(){
//     p := &pages{}
//     rows, err := db.Query("SELECT COUNT(*) from opelog;") 
//     checkErr(err)
//     for rows.Next(){
// 		var page int
//         err = rows.Scanp(&page)
//         checkErr(err)
//         p.pages = page
//     }
//     rows.Close()
//     data, err := json.Marshal(p)
//     return data
// }