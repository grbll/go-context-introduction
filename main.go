package main

import (
	"context"
	"fmt"
)

func doSomething(c context.Context) {
	fmt.Printf("I said hello to %v!\n", c.Value(true))

	cc := context.WithValue(c, true, "linda")
	doSomethingElse(cc)

	fmt.Printf("I said hello to %v!\n", c.Value(true))
}

func doSomethingElse(c context.Context) {
	fmt.Printf("I said bye to %v!\n", c.Value(true))
}
func main() {
	c := context.Background()

	c = context.WithValue(c, true, "hans")
	doSomething(c)
}
