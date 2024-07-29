package main

import (
	"fmt"
	"time"
)

func main() {
	go helloWorld()
	time.Sleep(10 * time.Second)
	goodBye()
}

func helloWorld() {
	fmt.Println("Hello World")
}

func goodBye() {
	fmt.Println("Good Bye")
}
