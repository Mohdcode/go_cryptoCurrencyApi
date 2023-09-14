package structtype

import "fmt"

type Admin struct {
    User
}

func NewInstructorAdmin(id int, name string) Admin {
    a := Admin{}
    a.Id = id
    a.Name = name
    return a
}

func (a Admin) Greet() {
    fmt.Println(a.Id)
}
