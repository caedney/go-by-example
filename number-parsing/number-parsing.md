# Number Parsing

Parsing numbers from strings is a basic but common task in many programs; here’s how to do it in Go.

The built-in package `strconv` provides the number parsing.

```go
import (
    "fmt"
    "strconv"
)
```

With `ParseFloat`, this 64 tells how many bits of precision to parse.

```go
f, _ := strconv.ParseFloat("1.234", 64)
fmt.Println(f) // 1.234
```

For `ParseInt`, the `0` means infer the base from the string. `64` requires that the result fit in 64 bits.

```go
i, _ := strconv.ParseInt("123", 0, 64)
fmt.Println(i) // 123
```

`ParseInt` will recognize hex-formatted numbers.

```go
d, _ := strconv.ParseInt("0x1c8", 0, 64)
fmt.Println(d) // 456
```

A `ParseUint` is also available.

```go
u, _ := strconv.ParseUint("789", 0, 64)
fmt.Println(u) // 789
```

`Atoi` is a convenience function for basic base-10 `int` parsing.

```go
k, _ := strconv.Atoi("135")
fmt.Println(k) // 135
```

Parse functions return an error on bad input.

```go
_, e := strconv.Atoi("wat")
fmt.Println(e) // strconv.Atoi: parsing "wat": invalid syntax
```

```sh
go run number-parsing/main.go
# 1.234
# 123
# 456
# 789
# 135
# strconv.Atoi: parsing "wat": invalid syntax
```
