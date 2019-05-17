package logs

import(
	"net/http"
	"fmt"
)


//测试入口
func Main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/log/getlog", GetLogs)
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

