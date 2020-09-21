package main

import (
	"fmt"
)

func main() {

	// like async await in nodejs ???
	done := make(chan bool)
	go helloGo(done)

	// uncoment this if not using channel
	// time.Sleep(1 * time.Second)

	<-done
	fmt.Println("ini fungsi main")
}

func helloGo(done chan bool) {
	fmt.Println("Hello Go")
	done <- true
}
