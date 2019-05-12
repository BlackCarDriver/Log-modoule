package logs
import(
	"net/http"
	"encoding/json"
)

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
}
func writeJson(w http.ResponseWriter, data interface{}) {
	jsondata, _ := json.Marshal(data)
	w.Write(jsondata)
}
func Testnet(w http.ResponseWriter, r *http.Request){
	writeJson(w,"connect scuess!")
}

func SendLogList(w http.ResponseWriter, r *http.Request){
	setHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	page := vars["page"]
	if len(page)==0 {
		return
	}
	mockdate := []cataloge{
		{"test1",[]string{"1","2","3"},},
		{"test1",[]string{"1","2","3"},},
		{"test1",[]string{"1","2","3"},},
		{"test1",[]string{"1","2","3"},},
		{"test1",[]string{"1","2","3"},},
	}
	writeJson(w, mockdate)
}
//log/getlogpage
//return specified log 
func SendLogText(w http.ResponseWriter, r *http.Request){
	setHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	floder := vars["floder"]
	name := vars["name"]
	if len(floder)==0 || len(name)==0 {
		Println("require body is unll!")
		return
	}
	Println(floder[0], "    ", name[0])
	temttext := []string{
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file quietly!  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file quietly!  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file quietly!  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file quietly!  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file  `,
	`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file quietly!  `,
	}
	writeJson(w,temttext)
}

type  cataloge struct{
	Name string  `json:"name"`
	List []string   `json:"list"`
}