/*
docker run -d -p 8080:8080 --rm -v "$PWD":/go/src/hello-world -w /go/src/hello-world golang go run rest.go
*/
package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello Mani, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
