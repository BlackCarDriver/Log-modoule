package logs

import(
	"net/http"
	"fmt"
	"os"
)

func Test(){
	// test_db()
	test_path()
	test_syslog()
	os.Exit(0)
}


func test_server(){
	mux := http.NewServeMux()
	mux.HandleFunc("/log/getlogtext", SendLogText)
	mux.HandleFunc("/log/getlogpage", SendLogList)
	mux.HandleFunc("/log/logcategory", Category)
    mux.HandleFunc("/log/logdisplay", Display)
	server := &http.Server{
		Addr : 			"localhost:8090",
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("begin to listen!")
	err := server.ListenAndServe()
	if err!=nil {
		fmt.Println(err)
	}	
}

func test_syslog(){
	Println("test_syslog")
	Log(Err,"Testting record error")
	Log(Warn, "Testing record warning")
	Log(Q_err,"Testting record error quietly")
	Log(Q_warn, "Testing record warning quietly")
}

func test_path(){
	pwd , _ := os.Getwd()
	fmt.Println("WorkPath :",pwd)
}

func test_db(){
	err := db.Ping()
	if err != nil {
		fmt.Println("Database connect fall! (logs)")
	}else{
		fmt.Println("database connect scuess!! (logs)")
	}
}