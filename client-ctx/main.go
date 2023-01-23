package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

type HTTPClientSettings struct {
	Connect          time.Duration
	ConnKeepAlive    time.Duration
	ExpectContinue   time.Duration
	IdleConn         time.Duration
	MaxAllIdleConns  int
	MaxHostIdleConns int
	ResponseHeader   time.Duration
	TLSHandshake     time.Duration
}

func NewHTTPClientWithSettings(httpSettings HTTPClientSettings) (*http.Client, error) {
	var client http.Client
	tr := &http.Transport{
		ResponseHeaderTimeout: httpSettings.ResponseHeader,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: httpSettings.ConnKeepAlive,
			DualStack: true,
			Timeout:   httpSettings.Connect,
		}).DialContext,
		MaxIdleConns:          httpSettings.MaxAllIdleConns,
		IdleConnTimeout:       httpSettings.IdleConn,
		TLSHandshakeTimeout:   httpSettings.TLSHandshake,
		MaxIdleConnsPerHost:   httpSettings.MaxHostIdleConns,
		ExpectContinueTimeout: httpSettings.ExpectContinue,
	}

	// So client makes HTTP/2 requests
	err := http2.ConfigureTransport(tr)
	if err != nil {
		return &client, err
	}

	return &http.Client{
		Transport: tr,
	}, nil
}

func main() {
	url := "http://localhost:3000"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// ctx, cancelFn := context.WithTimeout(context.TODO(), 3*time.Second)
	// defer cancelFn()
	ctx, cancel := context.WithCancel(context.Background())
	_ = time.AfterFunc(1*time.Second, func() { cancel() })

	httpClient, err := NewHTTPClientWithSettings(HTTPClientSettings{
		Connect:          5 * time.Second,
		ExpectContinue:   1 * time.Second,
		IdleConn:         90 * time.Second,
		ConnKeepAlive:    30 * time.Second,
		MaxAllIdleConns:  100,
		MaxHostIdleConns: 10,
		ResponseHeader:   5 * time.Second,
		TLSHandshake:     5 * time.Second,
	})
	if err != nil {
		fmt.Println("Got an error creating custom HTTP client:")
		fmt.Println(err)
		log.Fatal(err)
	}

	response, _ := httpClient.Do(req.WithContext(ctx))
	fmt.Println(response)
}
