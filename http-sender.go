package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func send(buf io.Reader) {
	resp, err := http.Post("http://localhost:8080/", "text/text", buf)

	if nil == err {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}

func send_many(buf io.Reader) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		send(buf)
	}
}

func main() {

	payload := "Hello, world."
	reader := bytes.NewBufferString(payload)

	wg.Add(1000)
	start := time.Now()
	for i := 0; i < 1000; i++ {
		go send_many(reader)
	}

	wg.Wait()
	delta := time.Since(start)

	fmt.Printf("Delta: %s\n", delta.String())
}
