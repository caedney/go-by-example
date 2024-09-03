# Environment Variables

[Environment variables](https://en.wikipedia.org/wiki/Environment_variable) are a universal mechanism for [conveying configuration information to Unix programs](https://www.12factor.net/config). Let’s look at how to set, get, and list environment variables.

To set a key/value pair, use `os.Setenv`. To get a value for a key, use `os.Getenv`. This will return an empty string if the key isn’t present in the environment.

```go
os.Setenv("FOO", "1")
fmt.Println("FOO:", os.Getenv("FOO"))
fmt.Println("BAR:", os.Getenv("BAR"))
```

Use `os.Environ` to list all key/value pairs in the environment. This returns a slice of strings in the form `KEY=value`. You can `strings.SplitN` them to get the key and value. Here we print all the keys.

```go
fmt.Println()
for _, e := range os.Environ() {
    pair := strings.SplitN(e, "=", 2)
    fmt.Println(pair[0])
}
```

Running the program shows that we pick up the value for FOO that we set in the program, but that BAR is empty.

```sh
go run environment-variables/main.go
# FOO: 1
# BAR:
# ...
```

If we set `BAR` in the environment first, the running program picks that value up.

```sh
BAR=2 go run environment-variables/main.go
# FOO: 1
# BAR: 2
# ...
```
