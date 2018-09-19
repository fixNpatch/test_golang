package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Object struct {
	url string
}

/*
func ReadByBufio(filePath string) string {

	var data string
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}*/
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
			answer = strings.TrimSuffix(answer, "\n")
			answer = strings.TrimSuffix(answer, "\r")
		}
	}

	buffer.url = answer
	array = append(array, buffer)
	answer = ""

	return array
}

func SaveContent(url string) bool {
	log.Print("Started with " + url)
	filePath := url
	filePath = strings.Trim(filePath, ".com")
	filePath = strings.Trim(filePath, ".ru")
	filePath += ".html"
	log.Print("Filepath to save: " + filePath)
	url = "https://" + url
	log.Print("HTTP Request at " + url)
	resp, err := http.Get(url)
	check(err)
	log.Print("HTTP Receive from " + url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body)
	WriteByOs(filePath, content)
	log.Print("Wrote content " + url)
	return true
}

func main() {
	var wg sync.WaitGroup
	data := ReadByIoutil("file.txt")
	array := ParseData(data)

	for _, element := range array {
		wg.Add(1)
		url := element.url
		go func() {
			defer wg.Done()
			SaveContent(url)
		}()
	}
	wg.Wait()
}
