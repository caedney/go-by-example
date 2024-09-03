# WaitGroups

To wait for multiple goroutines to finish, we can use a `wait group`.

This is the function we’ll run in every goroutine.

```go
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second) // Sleep to simulate an expensive task.
    fmt.Printf("Worker %d done\n", id)
}
```

This WaitGroup is used to wait for all the goroutines launched here to finish. Note: if a WaitGroup is explicitly passed into functions, it should be done _by pointer_.

```go
var wg sync.WaitGroup
```

Launch several goroutines and increment the WaitGroup counter for each. Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.

```go
for i := 1; i <= 5; i++ {
    wg.Add(1)

    go func() {
        defer wg.Done()
        worker(i)
    }()
}
```

Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.

```go
wg.Wait()
```

Note that this approach has no straightforward way to propagate errors from workers. For more advanced use cases, consider using the [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).

The order of workers starting up and finishing is likely to be different for each invocation.

```sh
go run waitgroups/main.go
# Worker 5 starting
# Worker 3 starting
# Worker 4 starting
# Worker 1 starting
# Worker 2 starting
# Worker 4 done
# Worker 1 done
# Worker 2 done
# Worker 5 done
# Worker 3 done
```
