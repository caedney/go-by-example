# HTTP Server

Writing a basic HTTP server is easy using the `net/http` package.

A fundamental concept in `net/http` servers is _handlers_. A handler is an object implementing the `http.Handler` interface. A common way to write a handler is by using the `http.HandlerFunc` adapter on functions with the appropriate signature.

Functions serving as handlers take a `http.ResponseWriter` and a `http.Request` as arguments. The response writer is used to fill in the HTTP response. Here our simple response is just “hello\n”.

```go
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}
```

This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.

```go
func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
```

We register our handlers on server routes using the `http.HandleFunc` convenience function. It sets up the _default router_ in the `net/http` package and takes a function as an argument.

```go
http.HandleFunc("/hello", hello)
http.HandleFunc("/headers", headers)
```

Finally, we call the `ListenAndServe` with the port and a handler. `nil` tells it to use the default router we’ve just set up.

```go
http.ListenAndServe(":8090", nil)
```

Run the server in the background.

```sh
go run http-server/main.go &
# [1] 50790
```

Access the `/hello` route.

```sh
curl localhost:8090/hello
```

Kill the process.

```sh
ps
# 50790 ttys002    0:00.19 go run http-server/main.go
# 50824 ttys002    0:00.01 /var/folders/pr/5llfzgdj3bg1d2jj_6dm3g700000gn/T/go-build2989645194/b001/exe/main
kill -9 50790 50824
```
