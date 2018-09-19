package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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
		case '\r':
			{
				break
			}
		default:
			answer += string(data[i])
			answer = strings.TrimSuffix(answer, "\n")
			answer = strings.TrimSuffix(answer, "\r")
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
		data := ReadByIoutil("file.txt")
		array := ParseData(data)
		url = "https://"
		url += array[rand.Intn(len(array))].url
		log.Print(url)
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
		url = ""
		//url = "https://yandex.ru"
		fmt.Fprint(w, AddForm(url))
		return
	}
}

func main() {
	/* Dont pay attention on that. It's a test */
	ReadWriteHandler()
	/* START LOCALHOST SERVER */
	http.HandleFunc("/test", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}
