package main

import (
	"fmt"
	"time"
	"runtime"
)

var count = 15

func ping(c chan<- int) {
	for i := 1; i < count; i++ {
		c <- 2*i - 1
	}
}
func pong(c chan<- int) {
	for i := 1; i < count; i++ {
		c <- 2 * i
	}
}

func print(ch <-chan int) {
	for {
		msg := <-ch
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	one()
	two()
	var input string
	fmt.Scanln(&input)
}
func one()  {
	ch := make(chan int)
	go ping(ch)
	go pong(ch)
	go print(ch)
}
func two()  {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 1; i < count; i++ {
			fmt.Println(2 * i)
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 1; i < count; i++ {
			fmt.Println(2*i - 1)
			runtime.Gosched()
		}
	}()
}