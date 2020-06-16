package main

import (
    "fmt"
    "log"
	"net/http"
    "io"
    "bytes"
)

func main() {
	url := "http://i.imgur.com/m1UIjW1.jpg"
    // don't worry about errors
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()
    fmt.Printf("size: %v\n", getSize(response.Body))
}

func getSize(stream io.Reader) int {
    buf := new(bytes.Buffer)
    buf.ReadFrom(stream)
    return buf.Len()/1024
}
