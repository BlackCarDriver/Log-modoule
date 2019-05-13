package logs

import(
	"net/http"
	"strings"
)

//return the catologe of system logs
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
		Log(Err,"require body is unll!")
		return
	}
	dir := strings.Replace(floder[0],"_","#",1)
	path := dir + `/` + name[0]
	logtext := readlogfile(path)
	writeJson(w,logtext)
}
