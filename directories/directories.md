# Directories

Go has several useful functions for working with _directories_ in the file system.

Create a new sub-directory in the current working directory.

```go
err := os.Mkdir("subdir", 0755)
check(err)
```

When creating temporary directories, it’s good practice to defer their removal. `os.RemoveAll` will delete a whole directory tree (similarly to `rm -rf`).

```go
defer os.RemoveAll("subdir")
```

Helper function to create a new empty file.

```go
createEmptyFile := func(name string) {
    d := []byte("")
    check(os.WriteFile(name, d, 0644))
}
createEmptyFile("subdir/file1")
```

We can create a hierarchy of directories, including parents with `MkdirAll`. This is similar to the command-line `mkdir -p`.

```go
err = os.MkdirAll("subdir/parent/child", 0755)
check(err)

createEmptyFile("subdir/parent/file2")
createEmptyFile("subdir/parent/file3")
createEmptyFile("subdir/parent/child/file4")
```

`ReadDir` lists directory contents, returning a slice of `os.DirEntry` objects.

```go
c, err := os.ReadDir("subdir/parent")
check(err)
fmt.Println("Listing subdir/parent")
for _, entry := range c {
    fmt.Println(" ", entry.Name(), entry.IsDir())
}
```

`Chdir` lets us change the current working directory, similarly to `cd`.

```go
err = os.Chdir("subdir/parent/child")
check(err)
```

Now we’ll see the contents of `subdir/parent/child` when listing the _current_ directory.

```go
c, err = os.ReadDir(".")
check(err)
fmt.Println("Listing subdir/parent/child")
for _, entry := range c {
    fmt.Println(" ", entry.Name(), entry.IsDir())
}
```

`cd` back to where we started.

```go
err = os.Chdir("../../..")
check(err)
```

We can also visit a directory _recursively_, including all its sub-directories. `WalkDir` accepts a callback function to handle every file or directory visited.

```go
fmt.Println("Visiting subdir")
err = filepath.WalkDir("subdir", visit)
```

`visit` is called for every file or directory found recursively by `filepath.WalkDir`.

```go
func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}
```

```sh
go run directories/main.go
# Listing subdir/parent
#   child true
#   file2 false
#   file3 false
# Listing subdir/parent/child
#   file4 false
# Visiting subdir
#   subdir true
#   subdir/file1 false
#   subdir/parent true
#   subdir/parent/child true
#   subdir/parent/child/file4 false
#   subdir/parent/file2 false
#   subdir/parent/file3 false
```
