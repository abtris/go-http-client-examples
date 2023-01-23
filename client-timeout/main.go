package main

import (
	"fmt"
	"net/http"
	"time"
)

// https://go.dev/src/net/http/transport.go
func main() {
	url := "http://localhost:3000"
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, _ := netClient.Get(url)
	fmt.Println(response)
}
