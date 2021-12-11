package comments

import (
	"database/sql"
	"fmt"
)

func DeleteMessage(){
	MessageID:=GetMessageID()
	sqlStr:="delete from messages where MessageID=?"
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/comment")
	if err!=nil{
		fmt.Println("err:",err)
		return
	}
	ret,errB:=db.Prepare(sqlStr)
	if errB!=nil{
		fmt.Println("errB:",errB)
		return
	}
	defer ret.Close()

	stmt,errC:=ret.Exec(MessageID)
	n,errD:=stmt.RowsAffected()
	if errC!=nil{
		fmt.Println("errC",errC)
		return
	}
	if n > 0{
		fmt.Println("delete successfully")
	}else{
		fmt.Println("delete failed:errD",errD)
		return
	}
}

func UpdateMessage(MessageID int,text string){
	sqlStr:="update messages set text=? where MessageID = ?"
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/comment")
	if err!=nil{
		fmt.Println("err",err)
		return
	}
	stmt,errB:=db.Exec(sqlStr,text,MessageID)
	if errB!=nil{
		fmt.Println("errB",errB)
		return
	}
	defer stmt.LastInsertId()
}

func DeleteOrUpdate(){
	fmt.Println("Delete(1) or Update(0)?")
	var flag int
	fmt.Scanln(&flag)
	ShowMyMessageID()
	if flag==1{
		DeleteMessage()
	}else if flag==0{
		MessageID:=GetMessageID()
		NewText:=GetText()
		UpdateMessage(MessageID,NewText)
	}
}
