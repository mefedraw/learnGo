package goroutines

import "fmt"

func Task2() {
	a := 5000
	for i := 0; i < a; i++ {
		go fmt.Println(i)
	}
}

/*
выведутся числа от 0 до 5000
в рандомном порядке и не до конца
*/
