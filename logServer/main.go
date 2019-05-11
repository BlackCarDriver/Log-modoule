package main

import(
	"./logs"
)

func main(){
	logs.Println("hahahahah","asdfasdfsdf","fasdfasdfasdfadsf")
	logs.Log(logs.Err, "it will recorde to the error file!")
	logs.Log(logs.Q_err, "It will not ouput to stdoutput!")
	logs.Log(logs.Warn, "It will not ouput to warn file")
	logs.Log(logs.Q_warn, "It will not ouput to warn file quietly!")
	// logs.Test()
}