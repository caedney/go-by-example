# Sorting

Go’s `slices` package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.

Sorting functions are generic, and work for any _ordered_ built-in type. For a list of ordered types, see [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).

```go
strs := []string{"c", "a", "b"}
slices.Sort(strs)
fmt.Println("Strings:", strs)
```

An example of sorting `int`s.

```go
ints := []int{7, 2, 4}
slices.Sort(ints)
fmt.Println("Ints:   ", ints)
```

We can also use the `slices` package to check if a slice is already in sorted order.

```go
s := slices.IsSorted(ints)
fmt.Println("Sorted: ", s)
```

```sh
go run sorting/main.go
# Strings: [a b c]
# Ints:    [2 4 7]
# Sorted:  true
```
