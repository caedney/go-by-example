# Command-Line Subcommands

Some command-line tools, like the go tool or `git` have many _subcommands_, each with its own set of flags. For example, `go build` and `go get` are two different subcommands of the go tool. The flag package lets us easily define simple subcommands that have their own flags.

We declare a subcommand using the NewFlagSet function, and proceed to define new flags specific for this subcommand.

```go
fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
fooEnable := fooCmd.Bool("enable", false, "enable")
fooName := fooCmd.String("name", "", "name")
```

For a different subcommand we can define different supported flags.

```go
barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
barLevel := barCmd.Int("level", 0, "level")
```

The subcommand is expected as the first argument to the program.

```go
if len(os.Args) < 2 {
    fmt.Println("expected 'foo' or 'bar' subcommands")
    os.Exit(1)
}
```

Check which subcommand is invoked. For every subcommand, we parse its own flags and have access to trailing positional arguments.

```go
switch os.Args[1] {
case "foo":
    fooCmd.Parse(os.Args[2:])
    fmt.Println("subcommand 'foo'")
    fmt.Println("  enable:", *fooEnable)
    fmt.Println("  name:", *fooName)
    fmt.Println("  tail:", fooCmd.Args())
case "bar":
    barCmd.Parse(os.Args[2:])
    fmt.Println("subcommand 'bar'")
    fmt.Println("  level:", *barLevel)
    fmt.Println("  tail:", barCmd.Args())
default:
    fmt.Println("expected 'foo' or 'bar' subcommands")
    os.Exit(1)
}
```

```sh
go build -o command-line-subcommands command-line-subcommands/main.go
```

First invoke the foo subcommand.

```sh
./command-line-subcommands/main foo -enable -name=joe a1 a2
# subcommand 'foo'
#   enable: true
#   name: joe
#   tail: [a1 a2]
```

Now try bar.

```sh
./command-line-subcommands/main bar -level 8 a1
# subcommand 'bar'
#   level: 8
#   tail: [a1]
```

But bar won’t accept foo’s flags.

```sh
./command-line-subcommands/main bar -enable a1
# flag provided but not defined: -enable
# Usage of bar:
#   -level int
#         level
```
