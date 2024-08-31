# Constants

Go supports _constants_ of character, string, boolean, and numeric values.

`const` declares a constant value.

```go
const s string = "constant" // constant
```

A `const` statement can appear anywhere a `var` statement can.

```go
const n = 500000000
```

Constant expressions perform arithmetic with arbitrary precision.

```go
const d = 3e20 / n
fmt.Println(d) // 6e+11
```

A numeric constant has no type until it’s given one, such as by an explicit conversion.

```go
fmt.Println(int64(d)) // 600000000000
```

A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here `math.Sin` expects a `float64`.

```go
fmt.Println(math.Sin(n)) // -0.28470407323754404
```

```sh
go run constants/main.go
# constant
# 6e+11
# 600000000000
# -0.28470407323754404
```
