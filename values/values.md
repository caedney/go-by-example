# Values

Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

Strings, which can be added together with +.

```go
fmt.Println("go" + "lang") // golang
```

Integers and floats.

```go
fmt.Println("1+1 =", 1+1) // 1+1 = 2
fmt.Println("7.0/3.0 =", 7.0/3.0) // 7.0/3.0 = 2.3333333333333335
```

Booleans, with boolean operators as you’d expect.

```go
fmt.Println(true && false) // false
fmt.Println(true || false) // true
fmt.Println(!true) // false
```

## Go Run

```sh
go run values/main.go
# golang
# 1+1 = 2
# 7.0/3.0 = 2.3333333333333335
# false
# true
# false
```
