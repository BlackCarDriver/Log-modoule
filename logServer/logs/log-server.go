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
	data := Readloglist()
	writeJson(w, data)
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
func Readloglist()[]cataloge{
	var data []cataloge
	dir,err:= os.Open(logs_root)
	if err!=nil {
		Log(Err,"Can not open logs_root when read logs list : ",err)
		return data
	}
	defer dir.Close()
	fileinfo,_:= dir.Readdir(-1)
	var date_times = make(map[string]int)
	for _,v := range fileinfo {		//traverse the floder
		if v.IsDir() == false {
			return data
		}
		dirname := v.Name()
		dt := strings.Split(dirname, "#")
		if len(dt)!=2 {
			Log(Warn, "Find an illeage directory name in /logsfile : ", dirname)
			continue
		}
		_, prs := date_times[dt[0]]
		if prs == false {	//if still not push in data[]
			var tc cataloge
			tc.Name = dt[0] 
			tc.List = append(tc.List, dt[1])
			data = append(data, tc)
		}else{ //already have a same date in data[]
			for i,v := range data {	//traverse the array
				if v.Name == dt[0] {
					data[i].List = append(data[i].List, dt[1])
					break
				}
			}
		}
		date_times[ dt[0] ] ++	
	}
	Println(data)
	return data
}

/* example of mock data
var	temttext = []string{
		`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file  `,
		`2019/05/11 13:04:07 [main.go] : It will not ouput to warn file quietly!  `,
}
var mockdate = []cataloge{
	{"test1",[]string{"1","2","3"},},
	{"test1",[]string{"1","2","3"},},
}
*/