package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type post struct {
	Body  string `json:"markdown"`
	Title string `json:"title"`
	Date  string `json:"created_at"`
}
type posts struct {
	Posts []post
}

var jsonFile = ".blog_posts.json"

var postTemplate = `
+++
title = "{{.Title}}"
draft = false
date = "{{.Date}}"

+++
{{.Body}}
`

func main() {

	// read all the docs and put them each into a struct
	// - Pull the markdown field
	// - Pull the title field
	// - Pull the created at field
	// output all the structs to markdown files
	jf, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(jf)

	var posts posts
	err = json.Unmarshal(jf, &posts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", posts)
	t := template.Must(template.New("post").Parse(postTemplate))

	for _, p := range posts.Posts {
		err := t.Execute(os.Stdout, p)
		if err != nil {
			fmt.Println("executing template:", err)
		}
	}

}
