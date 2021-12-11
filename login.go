package account

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func QueryRowForLogin(id int, password int){
	sqlStr:= "select name from users where password = ? and id = ?"
	db,err:=sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/users")
	if err != nil{
		fmt.Println("errorLogin",err)
		return
	}else {
		log.Println("连接数据库成功")
	}
	stmt,errC := db.Prepare(sqlStr)
	if errC != nil{
		fmt.Println("errC:", errC)
		return
	}
	defer stmt.Close()
	rows,errB := stmt.Query(id,password)
	if errB != nil{
		fmt.Println("errorLoginB",errB)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.Scan(&u.name)
		if err != nil {
			fmt.Println("errorLoginC",err)
			return
		}
		fmt.Printf("Hello %s", u.name)

		r := gin.Default()
		r.GET("/", func(c *gin.Context){
			//var u user
			if err:= c.ShouldBind(&u); err!=nil{
				c.JSON(http.StatusOK,gin.H{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"name": u.name,
				"status": "Log in successfully!",
			})
		})
		r.Run()
	}
}

func Login(){
    id := GetID()
	password := GetPassword()
	QueryRowForLogin(id,password)
}
