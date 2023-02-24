//go waitgroups
package main

import (
	"fmt"
	"sync"
)

func myFunc(id int, wg *sync.WaitGroup) {
	fmt.Println("Start", id)
	defer wg.Done()
	fmt.Println("Finish", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go myFunc(i, &wg)
	}

	wg.Wait()
	fmt.Println("End")
}
