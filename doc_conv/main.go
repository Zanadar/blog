package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Body  string `json:"markdown"`
	Title string `json:"title"`
	Date  string `json:"created_at"`
}
type Posts struct {
	Posts []Post
}

var jsonFile = ".blog_posts.json"

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

	var posts Posts
	err = json.Unmarshal(jf, &posts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", posts)

}
