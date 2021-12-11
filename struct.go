package comments

import (
	"database/sql"
	"fmt"
)

type message struct {
	id int       `json:"id"`
	name string  `json:"name"`
	MessageID int `json:"message_id"`
	text string   `json:"text"`
}

 func GetID()int{
	 fmt.Println("id?")
    var id int
	 fmt.Scanln(&id)
	return id
 }

 func GetPassword()int{
	 fmt.Println("Password?")
	 var password int
	 fmt.Scanln(&password)
	 return password
 }

 func GetMessageID()int{
	 fmt.Println("MessageID?")
	 var MessageID int
	 fmt.Scanln(&MessageID)
	 return MessageID
 }

func GetText()string{
	fmt.Println("Post your ideas!")
	var text string
	fmt.Scanln(&text)
	return text
}

func QueryRowForUserName(id int,password int)(name string){
	sqlStr:="select name from users where id=? and password=?"
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/users")
	if err!= nil{
		fmt.Println("QueryRowForUserName Failed", err)
	}
	var m message
	errB:=db.QueryRow(sqlStr,id,password).Scan(&m.name)
	if errB!=nil{
		fmt.Println("QueryRowForUserName errB Failed", errB)
	}
	return m.name
}