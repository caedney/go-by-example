# Recover

Go makes it possible to _recover_ from a panic, by using the `recover` built-in function. A `recover` can stop a `panic` from aborting the program and let it continue with execution instead.

An example of where this can be useful: a server wouldn’t want to crash if one of the client connections exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients. In fact, this is what Go’s `net/http` does by default for HTTP servers.

This function panics.

```go
func mayPanic() {
    panic("a problem")
}
```

`recover` must be called within a deferred function. When the enclosing function panics, the defer will activate and a `recover` call within it will catch the panic.

The return value of `recover` is the error raised in the call to `panic`.

The code after `mayPanic()` will not run, because `mayPanic` panics. The execution of `main` stops at the point of the panic and resumes in the deferred closure.

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()")
}
```

```sh
go run recover/main.go
# Recovered. Error:
#  a problem
```
