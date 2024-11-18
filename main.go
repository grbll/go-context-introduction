package main

import (
	"context"
	"fmt"
)

func doSomething(c context.Context) {
	fmt.Println("Did something!")
}
func main() {
	c := context.TODO()
	doSomething(c)
}
