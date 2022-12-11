package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("one\n")
	c := make(chan bool) //Create a channel. You need to give a type.
	go testFunction(c)
	fmt.Printf("two\n")
	areWeFinished := <-c //just wait until stg being sent to here. Basically this is blocking. Be careful not to create deadlocks :)
	fmt.Printf("areWeFinished: %v\n", areWeFinished)
}

func testFunction(c chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Checking...\n")
		time.Sleep(1 * time.Second)
	}
	c <- true //if you comment out this line. Go will realize a deadlock and exit.
}
