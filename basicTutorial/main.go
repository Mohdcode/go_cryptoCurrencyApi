package main

import (
	"fmt"
	fileutils "go_tutorial/Fileutils" //this fileutils is standing out becuase i didn't name the package same to dir name but keep consistency in other file in same dir
	"go_tutorial/data"
	"go_tutorial/second"
	structtype "go_tutorial/structType"
	"os"
) //fileutils the name we are giving to our package imported from Fileutils

func main() {
	 ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "one"
    }()
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "two"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println("Received from ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from ch2:", msg2)
    }
}
	print("appi")
	second.Hellothird()
	data.Init()
	println()
	result, _ := data.AddandSub(3, 5)
	print(result)
	println()
	var age int
	fmt.Scanf("%i", age)
	data.PointersProblem(&age)

	path, _ := os.Getwd() //this give the current directory

	str := fileutils.ReadTextFile(path + "/data/file.txt")
	fmt.Println(str)

	// var u1 structtype.User
	u1 := structtype.User{Id: 3, Name: "kamaal"}
	u1.Greet()
	u2:=structtype.NewInstructor(1);
	u2.Greet()
	a1:=structtype.NewInstructorAdmin( 2, "bilal")
	a1.Greet()
	var group [2] structtype.Signable
	group[0]=u1
	group[1]=a1
	for i,use := range group{
		fmt.Printf("%v this is the greet %v" ,i ,use)
	}

	
}
