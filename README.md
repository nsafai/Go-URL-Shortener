# Go URL Shortener

A [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

## How to run

- Clone repo
- [Install Go](https://golang.org/doc/install)
- Run `go build`
- Run `go run main/main.go`

### Credits
Largely inspired from [this Gophercises tutorial](https://gophercises.com/exercises/urlshort).