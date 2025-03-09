package goroutines

func Task4() {
	ch := make(chan bool) // сделать канал буф
	ch <- true
	go func() {
		<-ch
	}()
	ch <- true
}
