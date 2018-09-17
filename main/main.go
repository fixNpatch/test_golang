package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("main/file.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(hex.Dump(b))

	fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'
}
