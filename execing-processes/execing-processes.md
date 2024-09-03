# Exec'ing Processes

In the previous example we looked at spawning external processes. We do this when we need an external process accessible to a running Go process. Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one. To do this we’ll use Go’s implementation of the classic [`exec`](<https://en.wikipedia.org/wiki/Exec_(operating_system)>) function.

For our example we’ll exec `ls`. Go requires an absolute path to the binary we want to execute, so we’ll use `exec.LookPath` to find it (probably /`bin/ls`).

```go
binary, lookErr := exec.LookPath("ls")
if lookErr != nil {
    panic(lookErr)
}
```

`Exec` requires arguments in slice form (as opposed to one big string). We’ll give `ls` a few common arguments. Note that the first argument should be the program name.

```go
args := []string{"ls", "-a", "-l", "-h"}
```

`Exec` also needs a set of environment variables to use. Here we just provide our current environment.

```go
env := os.Environ()
```

Here’s the actual `syscall.Exec` call. If this call is successful, the execution of our process will end here and be replaced by the `/bin/ls -a -l -h` process. If there is an error we’ll get a return value.

```go
execErr := syscall.Exec(binary, args, env)
if execErr != nil {
    panic(execErr)
}
```

When we run our program it is replaced by `ls`.

```sh
go run execing-processes/main.go
# total 32
# drwxr-xr-x   4 craigedney  staff   128B Sep  3 12:59 execing-processes
# ...
```

Note that Go does not offer a classic Unix `fork` function. Usually this isn’t an issue though, since starting goroutines, spawning processes, and exec’ing processes covers most use cases for `fork`.
