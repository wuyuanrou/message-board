package account

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func QueryRowForNewUser(name string, password int){
	db,err:=sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/users")
	if err != nil{
		log.Fatal(err)
	}else {
		log.Println("连接数据库成功")
	}

	InsertStr := "insert into users (name,password) values (?,?)"
	ret,errC := db.Exec(InsertStr,name,password)
	if errC != nil{
		fmt.Println("errorC",errC)
		return
	}
	theID,errB:=ret.LastInsertId()
	if errB != nil{
		fmt.Printf("Get lastinsertid failed:%v",errB)
		return
	}
	fmt.Printf("id is %d",theID)
}


func NewUser(){
	name := GetName()
	password := GetPassword()
	QueryRowForNewUser(name,password)
     r := gin.Default()
	 r.GET("/",func(c *gin.Context){
		 var u user
		 if err := c.ShouldBind(&u); err != nil{
			 c.JSON(http.StatusOK,gin.H{
				 "error": err.Error(),
			 })
			 return
		 }
		 c.JSON(http.StatusOK,gin.H{
			 "username": name,
			 "status": "Create a new account successfully!",
		 })
	 })
	 r.Run()
}

