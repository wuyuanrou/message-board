package main

import (
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"message-board/comments"
)

func main(){
	//account.NewUser()
	//account.Login()
	//account.Password()
	comments.MessagesOrComments()
}
