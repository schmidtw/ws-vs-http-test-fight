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

func send(buf io.Reader, client *http.Client) {
	req, err := http.NewRequest("POST", "http://localhost:8080/", buf)
	resp, err := client.Do(req)

	if nil == err {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}

func send_many(buf io.Reader, client *http.Client) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		send(buf, client)
	}
}

func main() {

	payload := "Hello, world."
	reader := bytes.NewBufferString(payload)

	client := &http.Client{}

	wg.Add(1000)
	start := time.Now()
	for i := 0; i < 1000; i++ {
		go send_many(reader, client)
	}

	wg.Wait()
	delta := time.Since(start)

	fmt.Printf("Delta: %s\n", delta.String())
}
