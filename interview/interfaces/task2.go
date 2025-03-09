package interfaces

import "fmt"

func printType(whoami interface{}) {
	// вывести тип переменной
	fmt.Printf("%T\n", whoami)
}

func Task2() {
	printType(42)
	printType("im string")
	printType(true)
}
