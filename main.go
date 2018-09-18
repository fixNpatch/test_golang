package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AddForm() string {
	var content string = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
	return content
}

type Object struct {
	url string
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, AddForm())
		return
	}
	key := "Placeholder"
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

func ReadWriteHandler() {
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
	/*for i := 0; i < len(array); i++ {
		fmt.Println(array[i].url)
	}
	*/

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

func main() {
	/* Dont pay attention on that. It's a test */
	ReadWriteHandler()
	/* START LOCALHOST SERVER */
	http.HandleFunc("/add", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}
