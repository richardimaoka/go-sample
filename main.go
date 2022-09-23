package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	// UTF-8  Encoding: 0xF0 0x9F 0x9A 0x80
	// UTF-16 Encoding: 0xD83D 0xDE80
	// UTF-32 Encoding: 0x0001F680
	fmt.Print("\xF0\x9F\x9A\x80")
	// fmt.Println("\uD83D\uDE80") // doesn't compile
	// fmt.Println("\u1F680") //wrong!
	fmt.Println("\U0001F680 launched!!!")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read 1 byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("aborted!!")
		return
	}
	launch()

	// tick := time.Tick(100 * time.Millisecond)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	<-tick
	// }

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) //"0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
