package goroutines

import "fmt"

func Task5() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func() {
		ch <- true
	}()
	go func() {
		ch2 <- true
	}()
	go func() {
		ch3 <- true
	}()

	select {
	//case <-ch:
	//	fmt.Printf("val from ch")
	//case <-ch2:
	//	fmt.Printf("val from ch2")
	case <-ch3:
		fmt.Printf("val from ch3")
	}
}

/*
3 небуф канала, в них 3 горутинки записывают данные
и потом селект считывает из канала в который быстрее всего пришло
*/
