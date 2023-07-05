package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}
	stringJson := string(jsonData)
	strings.ReplaceAll(stringJson, "\n", "")

	fmt.Println(stringJson)
}
