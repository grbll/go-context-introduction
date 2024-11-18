package main

import (
	"context"
	"fmt"
)

func doSomething(c context.Context) {
	fmt.Printf("I said hello to %v!", c.Value(true))
}
func main() {
	c := context.Background()

	c = context.WithValue(c, true, "hans")
	doSomething(c)
}
