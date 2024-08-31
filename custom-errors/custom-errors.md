# Custom Errors

It’s possible to use custom types as `errors` by implementing the `Error()` method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

A custom error type usually has the suffix "Error".

```go
type argError struct {
    arg     int
    message string
}
```

Adding this `Error` method makes `argError` implement the `error` interface.

```go
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}
```

Return our custom error.

```go
func f(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }

    return arg + 3, nil
}
```

`errors.As` is a more advanced version of `errors.Is`. It checks that a given error (or any error in its chain) matches a specific error type and converts to a value of that type, returning `true`. If there’s no match, it returns `false`.

```go
func main() {
    _, err := f(42)
    var ae *argError

    if errors.As(err, &ae) {
        fmt.Println(ae.arg) // 42
        fmt.Println(ae.message) // can't work with it
    } else {
        fmt.Println("err doesn't match argError")
    }
}
```

```sh
go run custom-errors/main.go
# 42
# can't work with it
```
