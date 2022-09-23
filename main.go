package main

import (
	"fmt"
	"time"
)

func launch() {
	// UTF-8  Encoding: 0xF0 0x9F 0x9A 0x80
	// UTF-16 Encoding: 0xD83D 0xDE80
	// UTF-32 Encoding: 0x0001F680
	fmt.Println("\xF0\x9F\x9A\x80")
	// fmt.Println("\uD83D\uDE80") // doesn't compile
	// fmt.Println("\u1F680") //wrong!
	fmt.Println("\U0001F680")
}

func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(100 * time.Millisecond)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}
