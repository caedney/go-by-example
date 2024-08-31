# Variadic Functions

[Variadic functions](https://en.wikipedia.org/wiki/Variadic_function) can be called with any number of trailing arguments. For example, `fmt.Println` is a common variadic function.

Here’s a function that will take an arbitrary number of `ints` as arguments.

```go
func sum(nums ...int) {}
```

Within the function, the type of `nums` is equivalent to `[]int`. We can call `len(nums)`, iterate over it with `range`, etc.

```go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```

Variadic functions can be called in the usual way with individual arguments.

```go
sum(1, 2) // [1 2] 3
sum(1, 2, 3) // [1 2 3] 6
```

If you already have multiple args in a slice, apply them to a variadic function using `func(slice...)` like this.

```go
nums := []int{1, 2, 3, 4}
sum(nums...) // [1 2 3 4] 10
```

## Go Run

```sh
go run variadic-functions/main.go
# [1 2] 3
# [1 2 3] 6
# [1 2 3 4] 10
```

Another key aspect of functions in Go is their ability to form closures, which we’ll look at next.
