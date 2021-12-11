package account

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Find(id int) {
	sqlStr := "select name, password from users where id= ?"
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/users")
	if err != nil {
		fmt.Println(err)
		return
	}
    stmt,errA := db.Prepare(sqlStr)
	if errA !=nil{
		fmt.Println("errA:",errA)
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(id)
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.name,&u.password)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s,记住你的密码噢：%d",u.name,u.password)
	}
}

func Change(id int){
	sqlStr := "update users set password = ? where id = ?"
	db,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/users")
	if err != nil{
		fmt.Println(err)
	}
	stmt,errB:= db.Prepare(sqlStr)
	if errB!=nil{
		fmt.Println("errB:",errB)
	}
	defer stmt.Close()

	password := GetPassword()
	_,errD := stmt.Query(password,id)
	if errD != nil{
		fmt.Println("errD", errD)
	}
	_,errC := stmt.Exec(password,id)
	if errC != nil{
		fmt.Println("errC:", errC)
		return
	}
	fmt.Println("Change successfully!")
}

func Password(){
	var flag int
	r:=gin.Default()
	fmt.Println("Find(1) or Change(0)?")
	fmt.Scanln(&flag)
	if flag==1 {
		id := GetID()
		Find(id)
		r.GET("/", func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"status": "Find!",
			})
		})
		r.Run()
	}else if flag==0 {
		id := GetID()
		Change(id)
		r.GET("/", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"status": "Change",
			})
		})
		r.Run()
	}else{
		fmt.Println("Input 1 or 0")
	}
}
