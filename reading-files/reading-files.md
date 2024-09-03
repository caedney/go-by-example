# Reading Files

Reading and writing files are basic tasks needed for many Go programs. First we’ll look at some examples of reading files.

Reading files requires checking most calls for errors. This helper will streamline our error checks below.

```go
func check(e error) {
    if e != nil {
        panic(e)
    }
}
```

Perhaps the most basic file reading task is slurping a file’s entire contents into memory.

```go
dat, err := os.ReadFile("reading-files/data")
check(err)
fmt.Print(string(dat))
```

You’ll often want more control over how and what parts of a file are read. For these tasks, start by Opening a file to obtain an `os.File` value.

```go
f, err := os.Open("reading-files/data")
check(err)
```

Read some bytes from the beginning of the file. Allow up to 5 to be read but also note how many actually were read.

```go
b1 := make([]byte, 5)
n1, err := f.Read(b1)
check(err)
fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
```

You can also `Seek` to a known location in the file and Read from there.

```go
o2, err := f.Seek(6, io.SeekStart)
check(err)
b2 := make([]byte, 2)
n2, err := f.Read(b2)
check(err)
fmt.Printf("%d bytes @ %d: ", n2, o2)
fmt.Printf("%v\n", string(b2[:n2]))
```

Other methods of seeking are relative to the current cursor position,

```go
_, err = f.Seek(8, io.SeekCurrent)
check(err)
```

and relative to the end of the file.

```go
_, err = f.Seek(-8, io.SeekEnd)
check(err)
```

The `io` package provides some functions that may be helpful for file reading. For example, reads like the ones above can be more robustly implemented with `ReadAtLeast`.

```go
o3, err := f.Seek(6, io.SeekStart)
check(err)
b3 := make([]byte, 2)
n3, err := io.ReadAtLeast(f, b3, 2)
check(err)
fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
```

There is no built-in rewind, but `Seek(0, io.SeekStart)` accomplishes this.

```go
_, err = f.Seek(0, io.SeekStart)
check(err)
```

The `bufio` package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.

```go
r4 := bufio.NewReader(f)
b4, err := r4.Peek(5)
check(err)
fmt.Printf("5 bytes: %s\n", string(b4))
```

Close the file when you’re done (usually this would be scheduled immediately after Opening with `defer`).

```go
f.Close()
```

```sh
echo "hello" > reading-files/data
echo "go" >> reading-files/data
go run reading-files/main.go
# hello
# go
# 5 bytes: hello
# 2 bytes @ 6: go
# 2 bytes @ 6: go
# 5 bytes: hello
```
