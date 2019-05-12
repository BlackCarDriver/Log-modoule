package logs
import(
	"net/http"
	"encoding/json"
	"bufio"
	"strings"
	"os"
	"io"
)
//describe the cataloge of logsfiles list
type  cataloge struct{
	Name string  `json:"name"`
	List []string   `json:"list"`
}
func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
}
func writeJson(w http.ResponseWriter, data interface{}) {
	jsondata, _ := json.Marshal(data)
	w.Write(jsondata)
}
//test networkconnection
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
	writeJson(w, mockdate)
}
// (url)/log/getlogpage
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
	dir := strings.Replace(floder[0],"_","#",1)
	path := dir + `/` + name[0]
	Println(path)
	logtext := readlogfile(path)
	writeJson(w,logtext)
}



//read a logfile and save in a string array
//the paramater should be the reference path of logfiles
func readlogfile(path string)[]string{
	path = logs_root + path
	Println(path)
	var data []string
	f, err := os.Open(path)
	if err != nil {
		Log(Err,"Read log fall," , err)
		return data
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			Log(Err,"ReadLine Error,", err)
			return data
		}
		data = append(data, string(b))
	}
	if len(data)==0 {
		data = append(data,"null")
	}
	return data
}
//read the floder of log and return the catologe of logs files
func readloglist()[]cataloge{
	var data []cataloge

	return data
}

var	temttext = []string{
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
var mockdate = []cataloge{
	{"test1",[]string{"1","2","3"},},
	{"test1",[]string{"1","2","3"},},
	{"test1",[]string{"1","2","3"},},
	{"test1",[]string{"1","2","3"},},
	{"test1",[]string{"1","2","3"},},
}