package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	".."
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/nsafai": "http://nicolaisafai.com",
		"/g":     "http://google.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

// 	// Build the YAMLHandler using the mapHandler as the
// 	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `
// 	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
// 	if err != nil {
// 		panic(err)
// 	}

	jsonLinks := `[{
		"path":"/cool",
		"url":"http://cool.com"
}, {
		"path":"/new",
		"url":"http://new.com"
}]`

	jsonHandler, err := urlshort.JSONHandler([]byte(jsonLinks), mapHandler)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/links", allLinks)
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the URL Shortener API. For more info, visit https://github.com/nsafai/Go-URL-Shortener")
}

func allLinks(w http.ResponseWriter, r * http.Request) {
	jsonLinks := `[{
		"path":"/cool",
		"url":"http://cool.com"
}, {
		"path":"/new",
		"url":"http://new.com"
}]`
	fmt.Fprintln(w, jsonLinks)
}
