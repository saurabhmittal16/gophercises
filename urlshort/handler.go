package urlshort

import (
	"encoding/json"
	"log"
	"net/http"
)

// URLMapper is a struct to encapsulate Path and URL
type URLMapper struct {
	path string
	url  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, 301)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
func JSONHandler(JSONBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathMap := BuildMapFromJSON(JSONBytes)
	return MapHandler(pathMap, fallback), nil
}

// BuildMapFromJSON builds a map from JSON
func BuildMapFromJSON(JSONBytes []byte) map[string]string {
	temp := make(map[string]string)

	err := json.Unmarshal(JSONBytes, &temp)
	if err != nil {
		log.Fatal(err)
	}
	return temp
}
