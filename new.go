package comments

import (
	"database/sql"
	"fmt"
)

func CreateTableForMessage(){
	sqlStr := "create table `messages`(" +
		"`MessageID`bigint(20)not null auto_increment," +
		"`name`varchar(20)default ''," +
		"`id`int," +
		"`text`varchar(20) default ''," +
		"primary key(`MessageID`)" +
		")engine=InnoDB auto_increment=1 default charset=utf8mb4;"
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/comment")
	if err != nil{
		fmt.Println("err:",err)
		return
	}
	ret,errB := db.Prepare(sqlStr)
	if errB!=nil{
		fmt.Println("errB:",errB)
		return
	}
	defer ret.Close()

	row,errC:=ret.Query()
	if errC!= nil{
		fmt.Println("errC",errC)
		return
	}
	defer row.Close()

	for row.Next(){
		var m message
		errD:=row.Scan(&m.MessageID,&m.name,&m.text)
		if errD!=nil{
			fmt.Println("errD:",errD)
			return
		}
		fmt.Println(m.MessageID,m.name,m.text)
		fmt.Println("Create successfully")
	}
}

func InsertMessage(name string,id int,text string){
	sqlStr:="insert into messages(name,id,text)values(?,?,?)"
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/comment")
	if err!=nil{
		fmt.Println("err",err)
		return
	}
	stmt,errB:=db.Exec(sqlStr,name,id,text)
	if errB!=nil{
		fmt.Println("errB",errB)
		return
	}
	MessageID,errC:=stmt.LastInsertId()
	if errC!=nil{
		fmt.Println("insertMessage errC",errC)
		return
	}
	fmt.Println(MessageID)
}