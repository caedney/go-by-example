# Exit

Use `os.Exit` to immediately exit with a given status.

`defer`s will not be run when using `os.Exit`, so this `fmt.Println` will never be called.

```go
defer fmt.Println("!")
```

Exit with status 3.

```go
 os.Exit(3)
```

Note that unlike e.g. C, Go does not use an integer return value from `main` to indicate exit status. If you’d like to exit with a non-zero status you should use `os.Exit`.

If you run `exit.go` using `go run`, the exit will be picked up by go and printed.

```sh
go run exit/main.go
# exit status 3
```

By building and executing a binary you can see the status in the terminal.

```sh
go build -o exit exit/main.go
./exit/main
echo $?
# 3
```
