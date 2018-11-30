// docker run --rm -v "$PWD":/go/src/hello-world -w /go/src/hello-world golang go run hello-world.go
package main

import "fmt"

func main() {
    fmt.Println("hello from go")
}
