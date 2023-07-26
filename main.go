package main

import (
	"fmt"
	"time"
)

func main() {
	HelloGoRoutine()
}

func hello(i int) {
	println("hello" + fmt.Sprint(i))
}

func HelloGoRoutine() {
	for i := 0; i < 10; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}

func CalSquare() {
	src := make(chan int)
	dest := make(chan int)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	for i := range dest {
		println(i)
	}
}
