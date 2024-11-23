package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(c context.Context) {
	cc, cancelCtx := context.WithTimeout(c, time.Second)
	defer cancelCtx()
	var numChan chan int = make(chan int)

	go doSomethingElse(numChan, cc)
	for i := 0; i < 2000000; i++ {
		select {
		case numChan <- i:
			time.Sleep(10 * time.Millisecond)
		case <-cc.Done():
			break
		}
	}
	cancelCtx()
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething finished.\n")
}

func doSomethingElse(ch <-chan int, c context.Context) {
	for {
		select {
		case <-c.Done():
			if err := c.Err(); err != nil {
				fmt.Printf("Error in context: %v\n", err)
			}
			fmt.Printf("doSomethingElse finished due to closed context.\n")
			return
		case num := <-ch:
			fmt.Printf("%v was recieved.\n", num)
		}
	}
}
func main() {
	c := context.Background()
	doSomething(c)
}
