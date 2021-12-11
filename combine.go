package comments

import "fmt"

func PostMessage(){
	id:=GetID()
	password:=GetPassword()
	name:=QueryRowForUserName(id,password)
	text:=GetText()
	InsertMessage(name,id,text)
}

func MessagesOrComments(){
	//var flag int
	//fmt.Println("Messages(1) or Comments(0)?")
	//fmt.Scanln(&flag)
	//if flag==1{
	//	CreateTableForMessage()
	//
	fmt.Println("Post(1) or Check(0)?")
	var b int
	fmt.Scanln(&b)
	if b==1{
		PostMessage()
	}else{
		CheckMessage()
	}
	fmt.Println("Need to change?Yes(1)/No(0)")
	var c int
	fmt.Scanln(&c)
	if c==1{
		DeleteOrUpdate()
	}else{
		CheckMessage()
	}
}