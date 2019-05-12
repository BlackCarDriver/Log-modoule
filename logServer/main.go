package main

import(
	"./logs"
	"net/http"
	"fmt"
	"os"
)

func main(){
	//test()
	mux := http.NewServeMux()
	mux.HandleFunc("/", logs.Testnet)
	mux.HandleFunc("/log/getlogtext", logs.SendLogText)
	mux.HandleFunc("/log/getlogpage", logs.SendLogList)
	server := &http.Server{
		Addr : 			"localhost:8090",
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}
	logs.Println("begin to listen!")
	err := server.ListenAndServe()
	if err!=nil {
		fmt.Println(err)
	}
}

func test(){
	logs.Readloglist();
	os.Exit(0)
}