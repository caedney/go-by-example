# Methods

Go supports _methods_ defined on struct types.

This `area` method has a _receiver type_ of `*rect`.

```go
func (r *rect) area() int {
    return r.width * r.height
}
```

Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.

```go
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
```

Here we call the 2 methods defined for our struct.

```go
r := rect{width: 10, height: 5}
fmt.Println("area: ", r.area()) // area: 50
fmt.Println("perim:", r.perim()) // perim: 30
```

Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

```go
rp := &r
fmt.Println("area: ", rp.area()) // area: 50
fmt.Println("perim:", rp.perim()) // perim: 30
```

```sh
go run methods/main.go
# area:  50
# perim: 30
# area:  50
# perim: 30
```
