# Atomic Counters

The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the `sync/atomic` package for _atomic counters_ accessed by multiple goroutines.

We’ll use an atomic integer type to represent our (always-positive) counter.

```go
var ops atomic.Uint64
```

A WaitGroup will help us wait for all goroutines to finish their work.

```go
var wg sync.WaitGroup
```

We’ll start 50 goroutines that each increment the counter exactly 1000 times. To atomically increment the counter we use `Add`.

```go
for i := 0; i < 50; i++ {
    wg.Add(1)
    go func() {
        for c := 0; c < 1000; c++ {
            ops.Add(1)
        }
        wg.Done()
    }()
}
```

Wait until all the goroutines are done.

```go
wg.Wait()
```

Here no goroutines are writing to ‘ops’, but using `Load` it’s safe to atomically read a value even while other goroutines are (atomically) updating it.

```go
fmt.Println("ops:", ops.Load())
```

We expect to get exactly 50,000 operations. Had we used a non-atomic integer and incremented it with `ops++`, we’d likely get a different number, changing between runs, because the goroutines would interfere with each other. Moreover, we’d get data race failures when running with the `-race` flag.

```sh
go run atomic-counters/main.go
# ops: 50000
```
