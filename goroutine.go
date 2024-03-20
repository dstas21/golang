package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan string)

	go say("Hello", ch, ch2)
	fmt.Println(<-ch)
	fmt.Println(<-ch2)

	go counting(ch)
	for a := range ch {
		fmt.Println(a)
	}

	fmt.Println(<-ch) //дефолтное значение int
}

func say(str string, ch chan int, ch2 chan string) {
	fmt.Println(str)
	ch <- 7
	ch2 <- "От goroutine"
}

func counting(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
