package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string, 2)
	var wg sync.WaitGroup

	str1 := "Hello from channel 1"
	str2 := "Hello from channel 2"

	wg.Add(1)
	go sendMessage(ch, str1, &wg)
	wg.Wait()

	wg.Add(1)
	go sendMessage(ch, str2, &wg)
	wg.Wait()

	msg1 := <-ch
	msg2 := <-ch
	fmt.Println(msg1)
	fmt.Println(msg2)

}

func sendMessage(ch chan<- string, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error")
		}
	}()

	ch <- msg
}
