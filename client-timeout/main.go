package main

import (
	"fmt"
	"net/http"
	"time"
)

// https://go.dev/src/net/http/transport.go
func main() {
	url := "http://localhost:3000"
	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, _ := httpClient.Get(url)
	fmt.Println(response)
}
