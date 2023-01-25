# Go HTTP Client

- <https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779>

- <https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/>

- <https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/custom-http.html>

1. Dummy http client using default transport (`client` folder)

- <https://go.dev/src/net/http/transport.go>

    ```shell
    ❯ time go run client/main.go
    &{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[0] Date:[Mon, 23 Jan 2023 12:39:52 GMT]] {} 0 [] false false map[] 0x1400017e000 <nil>}
    go run client/main.go  0.20s user 0.18s system 0% cpu 1:40.37 total
    ```

2. Reasonable timeouts (`client-timeout` folder)`

    ```shell
    time go run client-timeout/main.go
    <nil>
    go run client-timeout/main.go  0.23s user 0.24s system 4% cpu 10.909 total
    ```

3. Custom client (`client-custom` folder)

    ```shell
    ➜ time go run client-ctx/main.go
    <nil>
    go run client-ctx/main.go  1.48s user 0.39s system 30% cpu 6.167 total
    ```

4. Active cancellation (`client-ctx` folder)`

    ```shell
    time go run client-ctx/main.go
    ```
