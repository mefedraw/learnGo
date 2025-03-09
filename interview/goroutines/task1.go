package goroutines

import (
	"fmt"
	"sync"
)

/*
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	wg.Wait()
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}
*/

func Task1() {
	ch := make(chan int, 3)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	wg.Wait()
	close(ch)
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}

/*
изначально мы видим небуферизированный канал
wg group, три горутинки запускаются и записывают данные в канал
потом считываем из канала данные

чтение начинается после того как три горутины уже заблокированы
а для записи в небуф канал нужно еще чтобы шло чтение в этот момент

я сделал канал буф, чтобы можно было записывать в него без одновременного чтения
и также закрыл канал после того как все горутинки запишут в него данные
что цикл for rage по каналу бесконечно не ждал данных
*/
