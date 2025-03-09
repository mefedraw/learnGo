package ChannelLesson

import (
	"fmt"
	"time"
)

func ChanRun() {
	exit := make(chan int)
	data := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)
}

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("waiting..")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
	close(exit)
}

func say(ch chan string, word string) {
	time.Sleep(3 * time.Second)
	fmt.Println(word)
	ch <- "exit"
}
