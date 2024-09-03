# File Paths

The `filepath` package provides functions to parse and construct _file paths_ in a way that is portable between operating systems; `dir/file` on Linux vs. `dir\file` on Windows, for example.

`Join` should be used to construct paths in a portable way. It takes any number of arguments and constructs a hierarchical path from them.

```go
p := filepath.Join("dir1", "dir2", "filename")
fmt.Println("p:", p) // p: dir1/dir2/filename
```

You should always use `Join` instead of concatenating /s or \s manually. In addition to providing portability, `Join` will also normalize paths by removing superfluous separators and directory changes.

```go
fmt.Println(filepath.Join("dir1//", "filename")) // dir1/filename
fmt.Println(filepath.Join("dir1/../dir1", "filename")) // dir1/filename
```

`Dir` and `Base` can be used to split a path to the directory and the file. Alternatively, `Split` will return both in the same call.

```go
fmt.Println("Dir(p):", filepath.Dir(p)) // Dir(p): dir1/dir2
fmt.Println("Base(p):", filepath.Base(p)) // Base(p): filename
```

We can check whether a path is absolute.

```go
fmt.Println(filepath.IsAbs("dir/file")) // false
fmt.Println(filepath.IsAbs("/dir/file")) // true
```

Some file names have extensions following a dot. We can split the extension out of such names with `Ext`.

```go
filename := "config.json"
ext := filepath.Ext(filename)
fmt.Println(ext) // .json
```

To find the file’s name with the extension removed, use `strings.TrimSuffix`.

```go
fmt.Println(strings.TrimSuffix(filename, ext)) // config
```

`Rel` finds a relative path between a _base_ and a _target_. It returns an error if the target cannot be made relative to base.

```go
rel, err := filepath.Rel("a/b", "a/b/t/file")
if err != nil {
    panic(err)
}
fmt.Println(rel) // t/file

rel, err = filepath.Rel("a/b", "a/c/t/file")
if err != nil {
    panic(err)
}
fmt.Println(rel) // ../c/t/file
```

```sh
go run file-paths/main.go
# p: dir1/dir2/filename
# dir1/filename
# dir1/filename
# Dir(p): dir1/dir2
# Base(p): filename
# false
# true
# .json
# config
# t/file
# ../c/t/file
```
