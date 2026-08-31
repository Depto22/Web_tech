package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func processNum(numChan chan int) {
// 	fmt.Println("Processing number ", <-numChan)
// }

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number ", num)
		time.Sleep(time.Second)
	}
}
func main() {
	// messageChan := make(chan string)
	// messageChan <- "ping"
	// msg := <-messageChan
	// fmt.Println(msg)

	// numChan := make(chan int)
	// go processNum(numChan)
	// numChan <- 5
	// time.Sleep(time.Second)

	numChan := make(chan int)
	go processNum(numChan)
	for {
		numChan <- rand.Intn(100)
	}

}
