# Maps

Maps are Go’s built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array) (sometimes called _hashes_ or _dicts_ in other languages).

To create an empty map, use the builtin `make`: `make(map[key-type]val-type)`.

```go
m := make(map[string]int)
```

Set key/value pairs using typical `name[key] = val` syntax.

```go
m["k1"] = 7
m["k2"] = 13
```

Printing a map with e.g. `fmt.Println` will show all of its key/value pairs.

```go
fmt.Println("map:", m) // map: map[k1:7 k2:13]
```

Get a value for a key with `name[key]`.

```go
v1 := m["k1"]
fmt.Println("v1:", v1) // v1: 7
```

If the key doesn’t exist, the [zero value](https://go.dev/ref/spec#The_zero_value) of the value type is returned.

```go
v3 := m["k3"]
fmt.Println("v3:", v3) // v3: 0
```

The builtin `len` returns the number of key/value pairs when called on a map.

```go
fmt.Println("len:", len(m)) // len: 2
```

The builtin `delete` removes key/value pairs from a map.

```go
delete(m, "k2")
fmt.Println("map:", m) // map: map[k1:7]
```

To remove all key/value pairs from a map, use the builtin `clear` method.

```go
clear(m)
fmt.Println("map:", m) // map: map[]
```

The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like `0` or `""`. Here we didn’t need the value itself, so we ignored it with the _blank identifier_ `_`.

```go
_, prs := m["k2"]
fmt.Println("prs:", prs) // prs: false
```

You can also declare and initialize a new map in the same line with this syntax.

```go
n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n) // map: map[bar:2 foo:1]
```

The `maps` package contains a number of useful utility functions for maps.

```go
n2 := map[string]int{"foo": 1, "bar": 2}
if maps.Equal(n, n2) {
    fmt.Println("n == n2")
}
// n == n2
```

Note that maps appear in the form `map[k:v k:v]` when printed with `fmt.Println`.

```sh
go run maps/main.go
# map: map[k1:7 k2:13]
# v1: 7
# v3: 0
# len: 2
# map: map[k1:7]
# map: map[]
# prs: false
# map: map[bar:2 foo:1]
# n == n2
```
