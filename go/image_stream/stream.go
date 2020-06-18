package main

import (
    "fmt"
    "log"
	"net/http"
    "io"
    "bytes"
    urlParser "net/url"
    "io/ioutil"
)

func main() {
	url := "http://i.imgur.com/m1UIjW1.jpg"
    // don't worry about errors
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()
    // fmt.Printf("size: %v\n", getSize(response.Body))
    buf := new(bytes.Buffer)
    buf.ReadFrom(response.Body)
    fmt.Printf("type: %v\n", http.DetectContentType(buf.Bytes()))
    url = "https://images-na.ssl-images-amazon.com/images/I/01RmK%2BJ4pJL._AC_UL200_SR200,200_.gif"
    decodedURL, err := urlParser.QueryUnescape(url)
    if err != nil {
        fmt.Printf("Error: %s\n", decodedURL)
    }
    fmt.Printf("decoded_url: %s\n", decodedURL)
    response, e = http.Get(url)
    if e != nil {
        log.Fatal(e)
    }
    body, err := ioutil.ReadAll(response.Body)
    fmt.Printf("decoded_url_type: %v\n", http.DetectContentType(body))
}

func getSize(stream io.Reader) int {
    buf := new(bytes.Buffer)
    buf.ReadFrom(stream)
    return buf.Len()/1024
}
