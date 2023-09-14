package data

import "fmt"

//var Countries [10]int
//var Countried [] int  -> this is slice
//var Codes map [int]bool  -> bracket is key and external is value type
//after theat can type in

func Init() {
	Countries := [2]int{2, 3}
	for i := 0; i < len(Countries); i++ {
		fmt.Print(Countries[i])

	}
}
func AddandSub(a int, b int) (int, int) {
	return a + b, a - b

}
func PointersProblem(a *int) {
	*a = 2
	fmt.Printf("%v", *a)
}
