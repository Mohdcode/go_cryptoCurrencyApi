package structtype

import "fmt"

type User struct{
	Id int
	Name string 
}
func NewInstructor (u int) User {
	return User{Id:1}
}
func (u User ) Greet(){
	fmt.Println(u.Name )
}
