# Rate Limiting

[_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting) is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.

First we’ll look at basic rate limiting. Suppose we want to limit our handling of incoming requests. We’ll serve these requests off a channel of the same name.

```go
requests := make(chan int, 5)
for i := 1; i <= 5; i++ {
    requests <- i
}
close(requests)
```

This `limiter` channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.

```go
limiter := time.Tick(200 * time.Millisecond)
```

By blocking on a receive from the `limiter` channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.

```go
for req := range requests {
    <-limiter
    fmt.Println("request", req, time.Now())
}
```

We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. We can accomplish this by buffering our limiter channel. This `burstyLimiter` channel will allow bursts of up to 3 events.

```go
burstyLimiter := make(chan time.Time, 3)
```

Fill up the channel to represent allowed bursting.

```go
for i := 0; i < 3; i++ {
    burstyLimiter <- time.Now()
}
```

Every 200 milliseconds we’ll try to add a new value to `burstyLimiter`, up to its limit of 3.

```go
go func() {
    for t := range time.Tick(200 * time.Millisecond) {
        burstyLimiter <- t
    }
}()
```

Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of `burstyLimiter`.

```go
burstyRequests := make(chan int, 5)
for i := 1; i <= 5; i++ {
    burstyRequests <- i
}
close(burstyRequests)

for req := range burstyRequests {
    <-burstyLimiter
    fmt.Println("request", req, time.Now())
}
```

Running our program we see the first batch of requests handled once every ~200 milliseconds as desired.

For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.

```sh
go run rate-limiting/main.go
# request 1 2024-09-02 11:54:31.265713 +0100 BST m=+0.201167001
# request 2 2024-09-02 11:54:31.464855 +0100 BST m=+0.400307834
# request 3 2024-09-02 11:54:31.665761 +0100 BST m=+0.601212459
# request 4 2024-09-02 11:54:31.864901 +0100 BST m=+0.800351626
# request 5 2024-09-02 11:54:32.068474 +0100 BST m=+1.003923251

# request 1 2024-09-02 11:54:32.068679 +0100 BST m=+1.004127709
# request 2 2024-09-02 11:54:32.068686 +0100 BST m=+1.004134959
# request 3 2024-09-02 11:54:32.06869  +0100 BST m=+1.004138667
# request 4 2024-09-02 11:54:32.269064 +0100 BST m=+1.204511501
# request 5 2024-09-02 11:54:32.46947  +0100 BST m=+1.404916334
```
