package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("file.txt")
	check(err)
	defer file.Close()

	bear := make([]byte, 5)
	str, err := file.Read(bear) // convert content to a 'string'
	check(err)

	fmt.Println(str) // print the content as a 'string'
}
