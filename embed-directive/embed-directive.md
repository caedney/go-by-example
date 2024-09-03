# Embed Directive

`//go:embed` is a [compiler directive](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) that allows programs to include arbitrary files and folders in the Go binary at build time. Read more about the embed directive [here](https://pkg.go.dev/embed).

Import the `embed` package; if you don’t use any exported identifiers from this package, you can do a blank import with `_ "embed"`.

```go
import (
	"embed"
)
```

`embed` directives accept paths relative to the directory containing the Go source file. This directive embeds the contents of the file into the `string` variable immediately following it.

```go
//go:embed folder/single_file.txt
var fileString string
```

Or embed the contents of the file into a `[]byte`.

```go
//go:embed folder/single_file.txt
var fileByte []byte
```

We can also embed multiple files or even folders with wildcards. This uses a variable of the [embed.FS type](https://pkg.go.dev/embed#FS), which implements a simple virtual file system.

```go
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS
```

Print out the contents of `single_file.txt`.

```go
print(fileString)
print(string(fileByte))
```

Retrieve some files from the embedded folder.

```go
content1, _ := folder.ReadFile("folder/file1.hash")
print(string(content1))
content2, _ := folder.ReadFile("folder/file2.hash")
print(string(content2))
```

Use these commands to run the example. (Note: due to limitation on go playground, this example can only be run on your local machine.)

```sh
mkdir embed-directive/folder
echo "hello go" > embed-directive/folder/single_file.txt
echo "123" > embed-directive/folder/file1.hash
echo "456" > embed-directive/folder/file2.hash

go run embed-directive/main.go
# hello go
# hello go
# 123
# 456
```
