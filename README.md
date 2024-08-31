# Go by example

## Getting started

When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a `go.mod` file that tracks the modules that provide those packages. That `go.mod` file stays with your code, including in your source code repository.

`init` initializes and writes a new `go.mod` file in the current directory, in effect creating a new module rooted at the current directory. The `go.mod` file must not already exist.

```
go mod init
```

### Naming a module

When you run `go mod init` to create a module for tracking dependencies, you specify a module path that serves as the module’s name. The module path becomes the import path prefix for packages in the module.

`init` accepts one optional argument, the module path for the new module. If the module path argument is omitted, `init` will attempt to infer the module path using import comments in `.go` files, vendoring tool configuration files (like `Gopkg.lock`), and the current directory (if in GOPATH).

```
go mod init example/hello
```

In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be `github.com/module`. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module.

### Updating Go

Change the version of Go in a project

```sh
go mod edit -go=1.23
```
