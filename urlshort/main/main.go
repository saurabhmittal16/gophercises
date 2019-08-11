package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"stark/gophercises/urlshort"
)

func main() {
	// JSON file containing paths and urls
	JSONbytes := getFileBytes("./paths.json")

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/google":         "https://www.google.com",
		"/fb":             "https://www.facebook.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the JSONHandler using the mapHandler as the fallback
	JSONHandler, err := urlshort.JSONHandler(JSONbytes, mapHandler)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", JSONHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getFileBytes(pathToFile string) []byte {
	bytes, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return nil
	}
	return bytes
}
