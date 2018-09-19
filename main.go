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

type Object struct {
	url string
}

func ReadByIoutil(filePath string) string {

	file, err := ioutil.ReadFile(filePath)
	check(err)
	data := string(file)
	return data
}
func WriteByOs(filePath string, data string) {
	file, err := os.Create(filePath)
	check(err)
	defer file.Close()
	_, err = file.WriteString(data)
	check(err)
}
func ParseData(data string) []Object {
	var counter = 0
	var answer string
	var array []Object
	var buffer Object
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case ';':
			{
				buffer.url = answer
				array = append(array, buffer)
				answer = ""
				counter++
			}
		case ' ':
			{ /*continue*/
			}
		default:
			answer += string(data[i])
		}
	}

	buffer.url = answer
	array = append(array, buffer)
	answer = ""

	return array
}
func ReadWriteHandler() {

	data := ReadByIoutil("file.txt")
	array := ParseData(data)

	var i = 0
	for i < len(array) {
		fmt.Println(array[i].url)
		i++
	}

	/* stringify (not json) */
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
	WriteByOs("file2.txt", answer)

}

func AddForm(url string) string {
	if url == "" {
		url = "http://example.com/"
	}
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body)
	WriteByOs("index.html", content)
	return content
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, AddForm("https://yandex.ru"))
		return
	}
	key := "Placeholder"
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
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
