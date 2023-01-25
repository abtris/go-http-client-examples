package main

import (
	"fmt"
	"net/http"
)

// https://go.dev/src/net/http/transport.go
func main() {
	url := "http://localhost:3000"
	var httpClient = &http.Client{}
	response, _ := httpClient.Get(url)
	fmt.Println(response)
}
