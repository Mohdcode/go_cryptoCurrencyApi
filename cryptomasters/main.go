package main

import (
	"fmt"

	"Mohcode.com/cryptomasters/api"
)

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Mark this worker as done when it finishes
// 	fmt.Printf("Worker %d is done\n", id)
// }

func main() {
	//var wg sync.WaitGroup

	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1) // Increment the WaitGroup counter
	// 	go worker(i, &wg)
	// }

	rate, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("the rate fro %v is %.2f", rate.Currency, rate.Price)

	} else {
		fmt.Println(err)
	}
	//wg.Wait() // Wait until all workers are done

}
