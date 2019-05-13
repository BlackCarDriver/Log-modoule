#logs package  function introduction

2019/5/13 9:14:44 

**目录说明**
   
│  log-server.go --服务端路由函数   
│  log-test.go   	--测试函数   
│  logs.go   --   	--系统日志函数   
│  orelog.go      --操作记录函数  
│   
└─logsfile   --日志目录   
····└─2019-05-11#1		-- 日志文件夹


##重要字段说明
[ ./logs.go ]   
logs_root  ： 设置系统日志的目录   
new_log    ： 因为每次启动都会创建一个文件夹，测试时频繁运行不方便，new_log设置
成0可以避免创建新文件夹。非测试使用应使 new_log = 1 。


##函数说明

Println(...interface{})   ---用于记录调试输出，数据记录在info.log文件   
Log(style int,  any interface{})   --系统日志记录，   

**Example**

//记录普通调试输出，数据记录在info.log文件   
**Println(“error",err)**

//记录危害较大的日志,保存位置：error.log   
**Log(logs.Err, "Testting record error")**   

//记录一般的错误日志，保存位置：warning.log   
**Log(logs.Warn,  "Testing record warning")**      
   
//使用Q_err, Q_warn，不会将日志输出到控制台   
Log(logs.Q_err, "Testting record error quietly")   
Log(logs.Q_warn,  "Testing record warning quietly")  

//记录增删改的操作日志，admin为操作者的用户名，module为操作的模块，logsql为对数据库的操作语句 
Records(admin, module, logsql)


## 注意事项
logs 包里面用到数据库连接，数据库变更时需要修改log-database.go 内参数。