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

type Object struct {
	url string
}

func (obj Object) appendChar(x string) string {
	slice := []string{obj.url}
	slice = append(slice, x)
	return obj.url
}

func main() {
	/* READ FILE */
	file, err := ioutil.ReadFile("file.txt")
	check(err)
	data := string(file)

	/* INIT STRUCTURE */
	var str string
	var counter = 0
	var array []Object
	var buffer Object

	/* PARSE DATA TO STRUCTURE */
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case ';':
			{
				buffer.url = str
				array = append(array, buffer)
				str = ""
				counter++
			}
		case ' ':
			{ /*continue*/
			}
		default:
			str += string(data[i])
		}
	}
	buffer.url = str
	array = append(array, buffer)
	str = ""

	/* PRINT ALL ELEMENTS OF STRUCTURE*/
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i].url)
	}

	/* PUSHING DATA TO TEMP*/
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
	/* PUSHING TEMP TO NEW FILE*/
	file2, err := os.Create("file2.txt")
	check(err)
	defer file2.Close()
	_, err = file2.WriteString(answer)
	check(err)

}
