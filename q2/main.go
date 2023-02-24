// go routine
package main

import (
	"fmt"
	"time"
)

func myFunc() {
	for i := 1; i <= 3; i++ {
		fmt.Println("First come:", i)
		time.Sleep(time.Second * 3)
	}
}

func main() {
	go myFunc()

	for i := 1; i <= 2; i++ {
		fmt.Println("First serve:", i)
		time.Sleep(time.Second * 3)
	}
}
