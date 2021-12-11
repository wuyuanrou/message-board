package comments

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckMessage(){
	sqlStr:="select name,text from messages"
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/comment")
	if err!=nil{
		fmt.Println("err:",err)
		return
	}

	stmt,errB:=db.Query(sqlStr)
	if errB!=nil{
		fmt.Println("errB",errB)
		return
	}
	defer stmt.Close()

	for stmt.Next(){
		var m message
		//m:=[]message{m}
		//msB:=append(ms,ms...)
		errC:=stmt.Scan(&m.name,&m.text)
		if errC!= nil{
			fmt.Println("errC",errC)
			return
		}

		r:=gin.Default()
		r.GET("/", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"name": m.name,
				"text": m.text,
			})
		})
	}
}

func ShowMyMessageID(){
	id:=GetID()
	password:=GetPassword()
	name:=QueryRowForUserName(password,id)
	sqlStr:="select MessageID,text from messages where name=? and id=?"
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/comment")
	if err!=nil{
		fmt.Println("No messageID,err",err)
		return
	}
	stmt,errB:=db.Query(sqlStr,name,id)
	if errB!=nil{
		fmt.Println("No messageID,errB",errB)
		return
	}
	defer stmt.Close()
	for stmt.Next(){
		var m message
		err:=stmt.Scan(&m.MessageID,&m.text)
		if err!=nil{
			fmt.Println("Scan failed,err:",err)
			return
		}

		r:=gin.Default()
		r.GET("/",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"Your_MessageID":m,
			})
		})
		r.Run()
	}
}
