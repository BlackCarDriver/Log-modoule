package main

import(
	"net/http"
	"fmt"
	"./logs"
	//"os"
)


func main(){
	//test()
	//logs.Records("dong","事件","delete youhreat")
	mux := http.NewServeMux()
	mux.HandleFunc("/log/logcategory", logs.Search)
    mux.HandleFunc("/log/logdisplay", logs.Display)
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