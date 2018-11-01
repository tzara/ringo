package main

import (
	"encoding/json"
	"fmt"
)

func j2m() {
	doc := `
		{
			"name": "maria",
			"age" : 10
		}
	`
	var data map[string]interface{}
	json.Unmarshal([]byte(doc), &data)

	fmt.Println(data["name"], data["age"])
}

func m2j() {
	data := make(map[string]interface{})
	data["name"] = "maria1"
	data["age"] = 10

	doc, _ := json.Marshal(data)
	fmt.Println(string(doc))

}

type Author struct {
	Name  string
	Email string
}

type Comment struct {
	Id      uint64
	Author  Author
	Content string
}

type Article struct {
	Id        uint64
	Title     string
	Author    Author
	Content   string
	Recomends []string
	Comments  []Comment
}

func main() {
	j2m()
	m2j()

	doc := `
	[{
		"Id": 1,
		"Title" : "Hello, World",
		"Author" : {
			"Name" : "Maria",
			"Email" : "maria@abt.com"
		},
		"Content" : "Hello ~",
		"Recommends" : [
			"John",
			"Andrew"
		],
		"Comments" : [{
		"Id": 1,
		"Author" : {
			"Name" : "Anfria",
			"Email" : "anf@abt.com"
		},
		"Content" : "hello Maria"
		}]
	}]
	`

	var data []Article
	json.Unmarshal([]byte(doc), &data)
	fmt.Println(data)
}
