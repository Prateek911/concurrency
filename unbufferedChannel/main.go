package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendMessage(ch)
	msg := <-ch
	fmt.Println("message received :", msg)

}

func sendMessage(ch chan<- string) {
	ch <- "Hello from the other side"
}
