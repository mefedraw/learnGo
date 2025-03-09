package goroutines

import "fmt"

func Task3() {
	ch := make(chan int)
	ch <- 1
	go func() {
		fmt.Println(<-ch)
	}()
}

/*
блокируется главная горутина которая выполняет Task3()
тк ch <- 1 данные в канал записываются а синхронно из них никто не читает
*/
