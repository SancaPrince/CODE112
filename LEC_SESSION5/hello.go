package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("ini gorout 1")

	go fmt.Println("ini gorout 2")

	go fmt.Println("ini gorout 3")

	time.Sleep(1 * time.Second)
}