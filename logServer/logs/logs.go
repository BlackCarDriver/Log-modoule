package logs

import (
	"net/http"
    "fmt"
    "encoding/json"
    "time"
	"strings"
    "strconv"
)

//==============================================================================

type returndata struct{
    Rowsnumber int `json:"rowsnumber"` ;
    Log        []opelog     `json:"log"`;
}

//数据库查询并返回日志，参数：日志类型,起始下标
//客户端页面每页显示11行数据,所以每次加载日志行数应为11的倍数
func getlogdata(ty string, index int)(data returndata){
    fmt.Println(ty,"   ",index)
    var logsdata []opelog;
    var rowsnumber int
    var err error
    sql1 := `select count(logid) from t_opelog where types=$1`;
    sql2 := `select types,operator,logtime,operation from t_opelog where types=$1 offset $2 limit 110`
    sql3 := `select count(logid) from t_opelog`;
    sql4 := `select types,operator,logtime,operation from t_opelog offset $1 limit 110`
    if ty=="all" {  //全部查询
        row := db.QueryRow(sql3)
        err = row.Scan(&rowsnumber)
        if err!=nil {
            return
        }
        rows, err2 := db.Query(sql4,index)
        if err2!=nil{
            return
        } 
        var i = index+1
        for rows.Next(){
            var temp opelog;
          
            err = rows.Scan(&temp.Types, &temp.Operator, &temp.Time, &temp.Operation)
            if err != nil {
                break
            }
            temp.Index = i
            i+=1
            temp.Time = parsetime( temp.Time )
            logsdata = append(logsdata,temp)
        }
    }else{       //条件查询
        row := db.QueryRow(sql1, ty)
        err := row.Scan(&rowsnumber)
        if err!=nil {
            fmt.Println(err)
        }
        rows, err2 := db.Query(sql2,ty,index)
        if err2!=nil{
            fmt.Println(err2)
        } 
        var i = index+1
        for rows.Next(){
            var temp opelog;
            temp.Index = i
            i+=1
            err2 = rows.Scan(&temp.Types, &temp.Operator, &temp.Time, &temp.Operation)
            if err2 != nil {
                fmt.Println(err2)
                break
            }
            temp.Time = parsetime( temp.Time )
            logsdata = append(logsdata,temp)
        }
    }
    data.Rowsnumber = rowsnumber;
    data.Log = logsdata;
    return data
}

//http://localhost:8090/log/getlog?type=warn&&index=0
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

//格式化时间字符串
func parsetime( timestr string)string{
	format:= "2006-01-02T15:04:05Z"
	time,err := time.Parse(format, timestr)
	if err!=nil {
		fmt.Println(err)
	    return ""
	}
	str := time.String()
	str = strings.Split(str,".")[0]
    return str
}




//==============================================================================


/* recycle bin

func Records(operator string, operation string){       //添加日志(登录)
	log.Println("正在添加日志...")
    rows, err := db.Prepare("insert into t_opelog (operator,operation) values($1,$2)")
    checkErr(err)
    _,err = rows.Exec(operator, operation)
}

type classify struct {
    Search          string    `json:"search"`
}

//参数不要用全局变量来传,多线程时可能会出错

var nowcontent string;
var nowsearch string;


func Search(w http.ResponseWriter, r *http.Request){          //获取查询
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
    defer r.Body.Close()
    con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &classify{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    fmt.Println("客户端输入：")
    fmt.Println("\tsearch:", su.Search)
    nowsearch=su.Search;
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
func Display(w http.ResponseWriter, r *http.Request){        
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	var opelog []opelog
	opelog = Getlogs()
	data, _ := json.Marshal(opelog)
    fmt.Fprintf(w,string(data))
}

*/