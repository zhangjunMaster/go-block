package model

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)
var (  
    dbhostsip  = "127.0.0.1:3306"//IP地址  
    dbusername = "root"//用户名  
    dbpassword = "123456"//密码  
    dbname     = "block_chain"//表名  
)  
type Block struct {
	id                 string
	Data			   string
	PrevBlockHash      string
	Hash			   string
	TimeStamp		   int
	Nonce			   int 
   }

var Db = &sql.DB{}

func init() {
	Db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block_chain?charset=utf8")  
}
