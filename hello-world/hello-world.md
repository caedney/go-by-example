# Hello, World!

Create a new directory called `hello-world` and add a `main.go` file with the following:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Wh can then run our code to see our greeting.

## Go Run

```sh
go run hello-world/main.go
# Hello, World!
```

## Go Build

Sometimes we’ll want to build our programs into binaries. We can do this using `go build`.

```sh
go build hello-world/main.go
```

We can then execute our code to see our greeting.

```sh
./main
# Hello, World!
```
