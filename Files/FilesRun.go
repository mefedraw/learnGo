package Files

import (
	"fmt"
	"io/ioutil"
)

func FilesRun() {
	writeToFile()
}

func readFile() {
	data, err := ioutil.ReadFile("nginx_logs.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func writeToFile() {
	err := ioutil.WriteFile("test.txt", []byte("hello world"), 0777)
	if err != nil {
		panic(err)
	}
}
