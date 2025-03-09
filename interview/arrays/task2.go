package arrays

import "fmt"

func Task2() {
	var foo []int
	var bar []int

	foo = append(foo, 1)
	foo = append(foo, 2)
	foo = append(foo, 3)
	bar = append(foo, 4)
	foo = append(foo, 5)

	fmt.Println(foo, bar)
}

/*
[1 2 3 5] [1 2 3 5]


слайс - структура с 3 полями:
указатель на массив данных, длина и вместимость

после трех добавлений у foo длина 3
bar начинает ссылаться на тот же массив что и foo
туда на 4 эл добавляется 4 и длина становится 4
но у foo она остается 3
потом в foo добавляют эл 5 и тк у foo длина 3 то эл 5
встает на 4 позицию -> и слайс bar перезаписывается
*/
