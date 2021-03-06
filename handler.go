package urlshort

import (
	"net/http"
	// yaml "gopkg.in/yaml.v2"
	"encoding/json"
	"fmt"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// if we can match a path...
		if dest, matchFound := pathsToUrls[path]; matchFound {
			// redirect to it
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		// else...
		fallback.ServeHTTP(w, r)
	}
}

/*********** YAML ***********/
// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
// func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
// 	// 1. Parse the YAML
// 	pathUrls, err := parseYaml(yamlBytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 2. Convert YAML array into map
// 	pathsToUrls := buildMap(pathUrls)
// 	// 3. return a map handler using the map
// 	return MapHandler(pathsToUrls, fallback), nil // returning nil as error because we're confident not to have an error here
// }

// func parseYaml(data []byte) ([]PathUrl, error) {
// 	var pathUrls []PathUrl
// 	err := yaml.Unmarshal(data, &pathUrls)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pathUrls, nil
// }

func buildMap(pathUrls []PathUrl) map[string]string {
	pathsToUrls:= make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

/*********** JSON ***********/

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the JSON
	pathUrls, err := ParseJson(jsonBytes)
	if err != nil {
		return nil, err
	}
	// 2. Convert JSON array into map
	pathsToUrls := buildMap(pathUrls)
	// 3. return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil // returning nil as error because we're confident not to have an error here
}

func ParseJson(data []byte) ([]PathUrl, error) {
	var pathUrls []PathUrl
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		fmt.Println("error parsing json!")
		return nil, err
	}
	// pu := PathUrl{"/test", "http://google.com"}
	return pathUrls, nil
}

type PathUrl struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}
