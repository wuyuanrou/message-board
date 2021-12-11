package account

import "fmt"

func GetID()int{
	var id int
	fmt.Scanln(&id)
	return id
}

func GetName()string{
	var name string
	fmt.Scanln(&name)
	return name
}

func GetPassword()int {
	var password int
	fmt.Scanln(&password)
	return password
}

type user struct{
	id int
	name string
	password int
}
