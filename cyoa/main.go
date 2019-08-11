package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// struct to encapsulate story attributes
type story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {
	// all stories in a map
	stories := getStories()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// take the query after /
		query := r.URL.Path[1:]
		t, _ := template.ParseFiles("template.html")

		// if no query, render intro
		if query == "" || query == "intro" {
			t.Execute(w, stories["intro"])
		} else {
			if dest, ok := stories[query]; ok {
				t.Execute(w, dest)
			} else {
				http.Error(w, "Chapter not found", http.StatusNotFound)
			}
		}
	})

	fmt.Printf("Server running at port: %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getStories reads story from JSON into map
func getStories() map[string]story {
	bytes, err := ioutil.ReadFile("./story.json")
	if err != nil {
		log.Fatal(err)
	}

	res := make(map[string]story)
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
