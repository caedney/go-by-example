# Temporary Files And Directories

Throughout program execution, we often want to create data that isn’t needed after the program exits. _Temporary files and directories_ are useful for this purpose since they don’t pollute the file system over time.

The easiest way to create a temporary file is by calling `os.CreateTemp`. It creates a file _and_ opens it for reading and writing. We provide "" as the first argument, so `os.CreateTemp` will create the file in the default location for our OS.

```go
f, err := os.CreateTemp("", "sample")
check(err)
```

Display the name of the temporary file. On Unix-based OSes the directory will likely be `/tmp`. The file name starts with the prefix given as the second argument to `os.CreateTemp` and the rest is chosen automatically to ensure that concurrent calls will always create different file names.

```go
fmt.Println("Temp file name:", f.Name())
```

Clean up the file after we’re done. The OS is likely to clean up temporary files by itself after some time, but it’s good practice to do this explicitly.

```go
defer os.Remove(f.Name())
```

We can write some data to the file.

```go
_, err = f.Write([]byte{1, 2, 3, 4})
check(err)
```

If we intend to write many temporary files, we may prefer to create a temporary _directory_. `os.MkdirTemp`’s arguments are the same as `CreateTemp`’s, but it returns a directory _name_ rather than an open file.

```go
dname, err := os.MkdirTemp("", "sampledir")
check(err)
fmt.Println("Temp dir name:", dname)

defer os.RemoveAll(dname)
```

Now we can synthesize temporary file names by prefixing them with our temporary directory.

```go
fname := filepath.Join(dname, "file1")
err = os.WriteFile(fname, []byte{1, 2}, 0666)
check(err)
```

```sh
go run temporary-files-and-directories/main.go
# Temp file name: /var/folders/pr/5llfzgdj3bg1d2jj_6dm3g700000gn/T/sample1222402219
# Temp dir name: /var/folders/pr/5llfzgdj3bg1d2jj_6dm3g700000gn/T/sampledir3821367140
```
