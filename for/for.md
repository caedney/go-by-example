# For

`for` is Go’s only looping construct. Here are some basic types of `for` loops.

The most basic type, with a single condition.

```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
// 1
// 2
// 3
```

A classic initial/condition/after `for` loop.

```go
for j := 0; j < 3; j++ {
    fmt.Println(j)
}
// 0
// 1
// 2
```

Another way of accomplishing the basic "do this N times" iteration is `range` over an integer.

```go
for i := range 3 {
    fmt.Println("range", i)
}
// range 0
// range 1
// range 2
```

`for` without a condition will loop repeatedly until you `break` out of the loop or `return` from the enclosing function.

```go
for {
    fmt.Println("loop") // loop
    break
}
```

You can also `continue` to the next iteration of the loop.

```go
for n := range 6 {
    if n%2 == 0 {
        continue
    }
    fmt.Println("odd", n)
}
// odd 1
// odd 3
// odd 5
```

```sh
go run for/main.go
# 1
# 2
# 3
# 0
# 1
# 2
# range 0
# range 1
# range 2
# loop
# odd 1
# odd 3
# odd 5
```
