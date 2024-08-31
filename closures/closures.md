# Closures

Go supports [anonymous functions](https://en.wikipedia.org/wiki/Anonymous_function), which can form [closures](<https://en.wikipedia.org/wiki/Closure_(computer_science)>). Anonymous functions are useful when you want to define a function inline without having to name it.

This function `intSeq` returns another function, which we define anonymously in the body of `intSeq`. The returned function _closes over_ the variable `i` to form a closure.

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

We call `intSeq`, assigning the result (a function) to `nextInt`. This function value captures its own i value, which will be updated each time we call `nextInt`.

```go
nextInt := intSeq()
```

See the effect of the closure by calling `nextInt` a few times.

```go
fmt.Println(nextInt()) // 1
fmt.Println(nextInt()) // 2
fmt.Println(nextInt()) // 3
```

To confirm that the state is unique to that particular function, create and test a new one.

```go
newInts := intSeq()
fmt.Println(newInts()) // 1
```

```sh
go run closures/main.go
# 1
# 2
# 3
# 1
```

The last feature of functions we’ll look at for now is recursion.
