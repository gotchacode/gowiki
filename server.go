package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}



func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/view", ViewHandler)
	http.ListenAndServe(":8080", nil)
}

func ViewHandler(w http.ResponseWriter, r*http.Request) {
	title := r.URL.Path[len("/view"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div> %s</div>", p.Title, p.Body)
}