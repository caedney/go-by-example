# Logging

The Go standard library provides straightforward tools for outputting logs from Go programs, with the log package for free-form output and the log/slog package for structured output.

Simply invoking functions like `Println` from the `log` package uses the _standard_ logger, which is already pre-configured for reasonable logging output to `os.Stderr`. Additional methods like `Fatal*` or `Panic*` will exit the program after logging.

```go
log.Println("standard logger") // 2024/09/03 11:43:58 standard logger
```

Loggers can be configured with _flags_ to set their output format. By default, the standard logger has the `log.Ldate` and `log.Ltime` flags set, and these are collected in `log.LstdFlags`. We can change its flags to emit time with microsecond accuracy, for example.

```go
log.SetFlags(log.LstdFlags | log.Lmicroseconds)
log.Println("with micro") // 2024/09/03 11:43:58.530215 with micro
```

It also supports emitting the file name and line from which the log function is called.

```go
log.SetFlags(log.LstdFlags | log.Lshortfile)
log.Println("with file/line") // 2024/09/03 11:43:58 main.go:19: with file/line
```

It may be useful to create a custom logger and pass it around. When creating a new logger, we can set a _prefix_ to distinguish its output from other loggers.

```go
mylog := log.New(os.Stdout, "my:", log.LstdFlags)
mylog.Println("from mylog") // my:2024/09/03 11:43:58 from mylog
```

We can set the prefix on existing loggers (including the standard one) with the `SetPrefix` method.

```go
mylog.SetPrefix("ohmy:")
mylog.Println("from mylog") // ohmy:2024/09/03 11:43:58 from mylog
```

Loggers can have custom output targets; any `io.Writer` works.

```go
var buf bytes.Buffer
buflog := log.New(&buf, "buf:", log.LstdFlags)
```

This call writes the log output into `buf`.

```go
buflog.Println("hello")
```

This will actually show it on standard output.

```go
fmt.Print("from buflog:", buf.String()) // from buflog:buf:2024/09/03 11:43:58 hello
```

The `slog` package provides _structured_ log output. For example, logging in JSON format is straightforward.

```go
jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
myslog := slog.New(jsonHandler)
myslog.Info("hi there") // {"time":"2024-09-03T11:43:58.530252+01:00","level":"INFO","msg":"hi there"}
```

In addition to the message, `slog` output can contain an arbitrary number of key=value pairs.

```go
myslog.Info("hello again", "key", "val", "age", 25) // {"time":"2024-09-03T11:43:58.530257+01:00","level":"INFO","msg":"hello again","key":"val","age":25}
```

Sample output; the date and time emitted will depend on when the example ran.

```sh
go run logging/main.go
# 2024/09/03 11:43:58 standard logger
# 2024/09/03 11:43:58.530215 with micro
# 2024/09/03 11:43:58 main.go:19: with file/line
# my:2024/09/03 11:43:58 from mylog
# ohmy:2024/09/03 11:43:58 from mylog
# from buflog:buf:2024/09/03 11:43:58 hello
# {"time":"2024-09-03T11:43:58.530252+01:00","level":"INFO","msg":"hi there"}
# {"time":"2024-09-03T11:43:58.530257+01:00","level":"INFO","msg":"hello again","key":"val","age":25}
```
