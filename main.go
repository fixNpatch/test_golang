package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("file.txt")
	check(err)
	data := string(file)
	var answer string
	for i := 0; i < len(data); i++ {
		if data[i] == ';' {
			answer += "\n"
		} else if data[i] == ' ' {
			continue
		} else {
			answer += string(data[i])
		}
	}

	file2, err := os.Create("file2.txt")
	check(err)
	defer file2.Close()

	_, err = file2.WriteString(answer)
	check(err)

	file3, err := ioutil.ReadFile("file2.txt")
	dataTest := string(file3)

	fmt.Print(dataTest)

}
