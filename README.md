# Log-modoule


2019/5/10 12:38:28 

It progress in a part of another bigger progress, the mainly function of it is recorder the log and display the 
log in the program. 


**logPage**   
logpage is the directory to save the code in angular, 

**logServer**   
logServer is the directory saving go code to build and run the server. the log file will save in /logServer/logs/***|*.log


##logServer function

**function :**
####Println( ...interface{} ) : void
//ouput the message to the console and inoflog.log   

#### Log(style int, msg ...interface{}):void
Log() write the message into logsfile,   
style control the way you record the log,   
if style =  Err or Q_err , Log() will record the logs into error.log,    
if style = Warn or Q_warn, Log() will record the logs into warning.log,   
note that Q_warn and Q_err will not output message on consloe


