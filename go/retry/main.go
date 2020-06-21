package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

type RetryCallback func() (err error)

func main() {
	//var signedContent []byte
	var resp *http.Response
	var err error
	url := "https://top100-s3-bucket.s3.ap-southeast-1.amazooonaws.com/images/gDC5ws5VqjthRc4o.jpg"
	var c = func() (err error) {
		resp, err = http.Get(url)
		return
	}
	err = retry(5, 2*time.Second, c)
	if err != nil {
    	log.Println(err)
	}
}

func retry(attempts int, sleep time.Duration, c RetryCallback) (err error) {
    for i := 0; ; i++ {
        err = c()
        if err == nil {
            return
        }

        if i >= (attempts - 1) {
            break
        }

        time.Sleep(sleep)

        log.Println("retrying after error:", err)
    }
    return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
