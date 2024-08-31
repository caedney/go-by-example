# Multiple Return Values

Go has built-in support for _multiple return values_. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

The `(int, int)` in this function signature shows that the function returns 2 `ints`.

```go
func vals() (int, int) {
    return 3, 7
}
```

Here we use the 2 different return values from the call with _multiple assignment_.

```go
a, b := vals()
fmt.Println(a) // 3
fmt.Println(b) // 7
```

If you only want a subset of the returned values, use the _blank identifier_ `_`.

```go
_, c := vals()
fmt.Println(c) // 7
```

## Go Run

```sh
go run multiple-return-values/main.go
# 3
# 7
# 7
```

Accepting a variable number of arguments is another nice feature of Go functions; we’ll look at this next.
