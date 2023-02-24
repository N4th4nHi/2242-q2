//go channels
package main

import (
    "fmt"
)

func myFunc(ch1 chan int, ch2 chan string) {
    for i := 1; i <= 5; i++ {
        ch1 <- i
    }
    ch2 <- "Done"
    close(ch1)
    close(ch2)
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan string)

    go myFunc(ch1, ch2)

    for {
        select {
        case i, ok := <-ch1:
            if ok {
                fmt.Println("Received from ch1:", i)
            } else {
                fmt.Println("ch1 closed")
                ch1 = nil
            }
        case s, ok := <-ch2:
            if ok {
                fmt.Println("Received from ch2:", s)
            } else {
                fmt.Println("ch2 closed")
                ch2 = nil
            }
        }
        if ch1 == nil && ch2 == nil {
            break
        }
    }
}