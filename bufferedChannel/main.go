package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	go sendMessage(ch, "Hello from channel 1")
	go sendMessage(ch, "Hello from channel 2")

	msg1 := <-ch
	msg2 := <-ch
	fmt.Println("msg1 :", msg1)
	fmt.Println("msg2 :", msg2)

}

func sendMessage(ch chan<- string, str string) {
	ch <- str
}
