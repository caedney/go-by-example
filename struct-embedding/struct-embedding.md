# Struct Embedding

Go supports _embedding_ of structs and interfaces to express a more seamless _composition_ of types. This is not to be confused with `//go:embed` which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.

A `container` _embeds_ a base. An embedding looks like a field without a name.

```go
type base struct {
    num int
}

type container struct {
    base
    str string
}
```

When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.

```go
co := container{
    base: base{
        num: 1,
    },
    str: "some name",
}
```

We can access the base’s fields directly on co, e.g. co.num.

```go
fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // co={num: 1, str: some name}
```

Alternatively, we can spell out the full path using the embedded type name.

```go
fmt.Println("also num:", co.base.num) // also num: 1
```

Since `container` embeds `base`, the methods of `base` also become methods of a `container`. Here we invoke a method that was embedded from base directly on `co`.

```go
fmt.Println("describe:", co.describe()) // describe: base with num=1
```

Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a `container` now implements the `describer` interface because it embeds base.

```go
type describer interface {
    describe() string
}

var d describer = co
fmt.Println("describer:", d.describe()) // describer: base with num=1
```

## Go Run

```sh
go run struct-embedding/main.go
# co={num: 1, str: some name}
# also num: 1
# describe: base with num=1
# describer: base with num=1
```
