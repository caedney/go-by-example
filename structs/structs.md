# Structs

Go’s _structs_ are typed collections of fields. They’re useful for grouping data together to form records.

This `person` struct type has name and age fields.

```go
type person struct {
    name string
    age  int
}
```

`newPerson` constructs a new person struct with the given name.

```go
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}
```

Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.

This syntax creates a new struct.

```go
person{"Bob", 20} // {Bob 20}
```

You can name the fields when initializing a struct.

```go
person{name: "Alice", age: 30} // {Alice 30}
```

Omitted fields will be zero-valued.

```go
person{name: "Fred"} // {Fred 0}
```

An `&` prefix yields a pointer to the struct.

```go
&person{name: "Ann", age: 40} // &{Ann 40}
```

It’s idiomatic to encapsulate new struct creation in constructor functions

```go
newPerson("Jon") // &{Jon 42}
```

Access struct fields with a dot.

```go
s := person{name: "Sean", age: 50}
fmt.Println(s.name) // Sean
```

You can also use dots with struct pointers - the pointers are automatically dereferenced.

```go
sp := &s
fmt.Println(sp.age) // 50
```

Structs are mutable.

```go
sp.age = 51
fmt.Println(sp.age) // 51
```

If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for [table-driven tests](https://gobyexample.com/testing-and-benchmarking).

```go
dog := struct {
    name   string
    isGood bool
}{
    "Rex", true,
}

fmt.Println(dog)
```

## Go Run

```sh
go run structs/main.go
# {Bob 20}
# {Alice 30}
# {Fred 0}
# &{Ann 40}
# &{Jon 42}
# Sean
# 50
# 51
# {Rex true}
```
