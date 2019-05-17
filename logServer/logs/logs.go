package logs

import (
	"log"
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
)

type opelog struct {
	Index           string    `json:"index"`
	Logid          string     `json:"logid"`
    Types        	string    `json:"type"`
    Operator        string    `json:"operator"`
    Time            string    `json:"time"`
    Operation    	string    `json:"operation"`
}

//==============================================================================
type returndata struct{
    Rowsnumber int `json:"rowsnumber"` ;
    Log        []opelog     `json:"log"`;
}

//数据库查询并返回日志，参数：日志类型,起始下标
//客户端页面每页显示11行数据,所以每次加载日志行数应为11的倍数
func getlogdata(ty string, index int) returndata{
    var logsdata []opelog;
    var rowsnumber int;
    sql1 := `select count(*) from t_logs where type=$1`;
    sql2 := `select type,operation,time,operation from oresqlwhere type=$1 offset $1 limit 55 `
    row := db.QueryRow(sql1, ty)
    err := row.Scan(&rowsnumber)
    if err!=nil {
        fmt.Println(err)
    }
    rows, err2 := db.Query(sql2,ty,index)
    if err2!=nil{
        fmt.Println(err2)
    } 
    for rows.Next(){
        var temp opelog;
        err2 = rows.Scan(&temp.Types, &temp.Operator, &temp.Time, &temp.Operation)
        if err2 != nil {
            fmt.Println(err2)
            break
        }
        logsdata = append(logsdata,temp)
    }
    var data returndata
    data.Rowsnumber = rowsnumber;
    data.Log = logsdata;
    return data
}

// (url) /log/getlog
//接收请求并返回日志数据
func GetLogs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")            
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") 
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
    tag := vars["type"]
    indexstr := vars["index"]
    if( len(tag)==0 || len(indexstr)==0 ){
        return
    }
    index,err := strconv.Atoi( indexstr[0] )
    if err!=nil{
        fmt.Println(err)
        return
    }
    data := getlogdata(tag[0], index)
    jsondata, _ := json.Marshal(data)
    w.Write(jsondata)
	return
}

//==============================================================================






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
