# Recursion

Go supports [recursive functions](<https://en.wikipedia.org/wiki/Recursion_(computer_science)>). Here’s a classic example.

This `fact` function calls itself until it reaches the base case of `fact(0)`.

```go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
```

Closures can also be recursive, but this requires the closure to be declared with a typed `var` explicitly before it’s defined.

```go
var fib func(n int) int
```

Since `fib` was previously declared in `main`, Go knows which function to call with `fib` here.

```go
fib = func(n int) int {
    if n < 2 {
        return n
    }

    return fib(n-1) + fib(n-2)
}
```

## Go Run

```sh
go run recursion/main.go
# 5040
# 13
```
