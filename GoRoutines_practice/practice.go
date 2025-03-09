package GoRoutines_practice

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func doRequest(results chan string) {
	results <- fmt.Sprintf("some data #%d", rand.Int()%500)
	results <- fmt.Sprintf("some data #%d", rand.Int()%500)
}

func storeData(results chan string) {
	fmt.Println("before  blocking")
	for {
		select {
		case data, ok := <-results:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Printf("stored data %s\n", data)

		default:
			fmt.Println("no data in channel")
			time.Sleep(1 * time.Second)
		}
	}
}

func ChannelRead() {
	results := make(chan string)
	go storeData(results)
	go doRequest(results)
	fmt.Println("start sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("end sleep")
}
