# Epoch

A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the [Unix epoch](https://en.wikipedia.org/wiki/Unix_time). Here’s how to do it in Go.

Use `time.Now` with `Unix`, `UnixMilli` or `UnixNano` to get elapsed time since the Unix epoch in seconds, milliseconds or nanoseconds, respectively.

```go
now := time.Now()
fmt.Println(now) // 2024-09-02 16:25:16.681614 +0100 BST m=+0.000073084
fmt.Println(now.Unix()) // 1725290716
fmt.Println(now.UnixMilli()) // 1725290716681
fmt.Println(now.UnixNano()) // 1725290716681614000
```

You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.

```go
fmt.Println(time.Unix(now.Unix(), 0)) // 2024-09-02 16:25:16 +0100 BST
fmt.Println(time.Unix(0, now.UnixNano())) // 2024-09-02 16:25:16.681614 +0100 BST
```

```sh
go run epoch/main.go
```
