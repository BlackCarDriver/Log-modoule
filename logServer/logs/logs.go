package logs

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
	"os"
	"io/ioutil"
	"strconv"
)

const(
	// is the path saving the logs files
	logs_root= `./logs/logsfile/`
	// you can set mew_log = 0 when you are testing,
	// then it will cover the oldest flode to save the logsfiles 
	//note that new_log should be 0 or 1 !
	new_log = 1
)
 
var(
	error_log_name = "error.log"
	warn_log_name = "warning.log"
	info_log_name = "info.log"
)

var errlog *log.Logger
var warnlog *log.Logger
var infolog *log.Logger


//the following constant are the paramater used in Log() style
const (
	Err = 400
	Warn = 200
	Q_err =-400
	Q_warn =-200
)

func init() {
	//connect to database
	testconnect()

	log_floder := createfloder()
	errorp, err := os.Create(log_floder + error_log_name)
	warnp, _ := os.Create(log_floder + warn_log_name)
	infop, _ := os.Create(log_floder + info_log_name)
	if err != nil  {
		panic("can not create logsfile!")
	}
	//if close the *File,the log can't be writen in file!
	errlog = log.New(errorp, "", 3)
	warnlog = log.New(warnp, "", 3)
	infolog = log.New(infop, "", 3)
}

//return the user file location, like "main.go"
func getpath()string{
	_,file,_,ok := runtime.Caller(2)
	if ok==false{
		return "???"
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}

//set the format of pathstring
func formatPath(path string) string{
	return  "[" + path + "] : "
}

//the raw ouput is like :[a,b,c]
//and after it function handle, it will be : a  b  c
func formatInterface(prefix string, any ...interface{})string{
	formatstr := strings.Repeat("%v  ", len(any))
	prefix += fmt.Sprintf(formatstr, any...)
	return prefix
}

//create a new floder to save the logs file, return the path of new floder
func createfloder()string{
	rd, err := ioutil.ReadDir(logs_root)
	if err != nil {
		fmt.Println("Can not read direcotry !")
		panic(err)
	}
	filenum := new_log
	for _, fi := range rd {
		if fi.IsDir(){
			filenum++
		}
	}
 	datestr := time.Now().Format("2006-01-02")
	tempPath := logs_root + datestr + `#` + strconv.Itoa(filenum) + `\`
	if new_log == 0 {
		err := os.RemoveAll(tempPath)
		if err!= nil {
			Log(Err,"Remove director Fall : ",err)
		}
	}
	err = os.Mkdir(tempPath, os.ModeDir)
	if err!=nil{
		fmt.Println("Can not make directory !")
		panic(err)
	}
	return tempPath
}

//write the message into logfile, style control the way you record the log,
//style use Err and Q_err will record the logs into error.log, 
//style use Warn and Q_warn will record the logs into warning.log,
//note that Q_warn and Q_err will not output message on consloe
func Log(style int, msg ...interface{}){
	path := getpath()
	path = formatPath(path)
	s := formatInterface(path, msg...)
	if style > 0 {
		fmt.Println(s)
	}else{
		style = 0-style
	} 
	switch style {
	case Err :
		errlog.Println(s)
	case Warn:
		warnlog.Println(s)
	}
}

//compare to fmt.Println, it Println will display the caller and record intto logfile
func Println(any ...interface{}) {
	path := getpath()
	path = formatPath(path)
	s := formatInterface(path , any...)
	fmt.Println(s)
	infolog.Println(s)
}

