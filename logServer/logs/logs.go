package logs

import (
	//"../models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
)


type classify struct {
    Search          string    `json:"search"`
}

var nowcontent string;
var nowsearch string;

func checkErr(err error) {   //报错
    if err != nil {
        log.Println("出错啦!")
        panic(err)
    }
}

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

func Display(w http.ResponseWriter, r *http.Request){         //
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	var opelog []opelog
	opelog = Getlogs()
	data, _ := json.Marshal(opelog)
    fmt.Fprintf(w,string(data))
}
