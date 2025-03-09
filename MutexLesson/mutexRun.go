package MutexLesson

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu  sync.Mutex
	cnt map[string]int
}

func MutexRun() {
	key := "test"

	c := Counter{cnt: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(c.Value(key))
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.cnt[key]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cnt[key]
}

//Для записи используйте Lock и Unlock.
//Они блокируют доступ к данным для всех горутин — как для читающих,
//так и для записывающих.

//Для чтения используйте RLock и RUnlock.
//Эти методы позволяют нескольким горутинам одновременно читать данные,
//но блокируют доступ для записывающих горутин.

type CounterWithRWMutex struct {
	mu  sync.RWMutex
	cnt map[string]int
}

func (cr *CounterWithRWMutex) CountMeAgain() map[string]int {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	return cr.cnt
}

func WgGroupRun() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()

			fmt.Println("goroutine working..", k)
			time.Sleep(time.Millisecond * 300)
		}()
	}

	wg.Wait()
	fmt.Println("all done")
}
