package main

import (
	"fmt"
	"net/http"
)

// https://go.dev/src/net/http/transport.go
func main() {
	url := "http://localhost:3000"
	var netClient = &http.Client{}
	response, _ := netClient.Get(url)
	fmt.Println(response)
}
