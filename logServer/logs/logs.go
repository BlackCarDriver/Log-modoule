package logs

import (
	//"../models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Records(admin string, content string){       //添加日志(登录)
	log.Println("正在添加日志...")
    rows, err := db.Prepare("insert into t_opelog (admin,content) values($1,$2)")
    checkErr(err)
    _,err = rows.Exec(admin, content)
    checkErr(err)
    log.Println("日志添加成功！")
}

func Getlogs()[]opelog{
    var Opelog []opelog
    if nowsearch!="" {
        rows, err := db.Query("select logid,module,admin,logtime,content from oresql where module=$2",nowsearch) 
        checkErr(err)
        for rows.Next(){
		var opelog opelog
		err = rows.Scan(&opelog.Logid,&opelog.Module,&opelog.Admin,&opelog.Logtime,&opelog.Content)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        Opelog=append(Opelog,opelog)
	}
    }else {
        rows, err := db.Query("select logid,module,admin,logtime,content from t_opelog") 
        checkErr(err)
        for rows.Next(){
		var opelog opelog
		err = rows.Scan(&opelog.Logid,&opelog.Logtime,&opelog.Admin,&opelog.Module,&opelog.Content)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        Opelog=append(Opelog,opelog)
	}
    }
    fmt.Println("正在读取数据库")
    return Opelog
}
