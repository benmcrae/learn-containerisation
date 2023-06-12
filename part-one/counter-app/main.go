package main

import (
	"fmt"
	"time"
)

func main() {
	var c int
	for {
		fmt.Printf("Hello, world! My counter value is %d\n", c)
		time.Sleep(1 * time.Second)
		c++
	}
}
