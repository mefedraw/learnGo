package arrays

import "fmt"

func Task4() {
	var m = make(map[string]int)

	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
карта, объявленная через var, по умолчанию имеет значение nil.
Это значит, что она не инициализирована,
и в неё нельзя сразу записывать данные
*/
