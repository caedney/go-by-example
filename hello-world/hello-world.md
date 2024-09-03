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

If we want to rename the output binary we can use the `-o` flag. The below script will place a `main.go` file in the `hello-world` directory as `hello-world` already exists otherwise it will rename the binary to `hello-world`.

```sh
go build -o hello-world hello-world/main.go
```

We can then execute our code to see our greeting.

```sh
./hello-world/main
# Hello, World!
```
