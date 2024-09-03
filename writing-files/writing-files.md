# Writing Files

Writing files in Go follows similar patterns to the ones we saw earlier for reading.

To start, here’s how to dump a string (or just bytes) into a file.

```go
d1 := []byte("hello\ngo\n")
err := os.WriteFile("writing-files/data1.txt", d1, 0644)
check(err)
```

For more granular writes, open a file for writing.

```go
f, err := os.Create("writing-files/data2.txt")
check(err)
```

It’s idiomatic to defer a `Close` immediately after opening a file.

```go
defer f.Close()
```

You can `Write` byte slices as you’d expect.

```go
d2 := []byte{115, 111, 109, 101, 10}
n2, err := f.Write(d2)
check(err)
fmt.Printf("wrote %d bytes\n", n2)
```

A `WriteString` is also available.

```go
n3, err := f.WriteString("writes\n")
check(err)
fmt.Printf("wrote %d bytes\n", n3)
```

Issue a `Sync` to flush writes to stable storage.

```go
f.Sync()
```

`bufio` provides buffered writers in addition to the buffered readers we saw earlier.

```go
w := bufio.NewWriter(f)
n4, err := w.WriteString("buffered\n")
check(err)
fmt.Printf("wrote %d bytes\n", n4)
```

Use `Flush` to ensure all buffered operations have been applied to the underlying writer.

```go
w.Flush()
```

```sh
go run writing-files/main.go

cat writing-files/data1
# hello
# go

cat writing-files/data2
# some
# writes
# buffered
```
