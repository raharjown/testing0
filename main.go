package main

import (
	"fmt"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// runtime.GOMAXPROCS(2)

	go print(5, "halo")
	go print(5, "apa kabar")

	fmt.Println("asdasdas")
	time.Sleep(1000000000)
}
