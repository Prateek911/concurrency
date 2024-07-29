package main

import "fmt"

func main() {
	//bi-directional channel
	ch := make(chan string)
	go sendMessage(ch)
	fmt.Println(<-ch)
}

// sender channel
func sendMessage(ch chan<- string) {
	ch <- "Hello from the otter side"
}
