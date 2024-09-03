# Command-Line Arguments

[Command-line arguments](https://en.wikipedia.org/wiki/Command-line_interface#Arguments) are a common way to parameterize execution of programs. For example, `go run hello.go` uses `run` and `hello.go` arguments to the `go` program.

`os.Args` provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.

```go
argsWithProg := os.Args
argsWithoutProg := os.Args[1:]
```

You can get individual args with normal indexing.

```go
arg := os.Args[3]

fmt.Println(argsWithProg)
fmt.Println(argsWithoutProg)
fmt.Println(arg)
```

To experiment with command-line arguments it’s best to build a binary with `go build` first.

```sh
go build -o command-line-arguments command-line-arguments/main.go
./command-line-arguments/main a b c d
# [./main a b c d]
# [a b c d]
# c
```
