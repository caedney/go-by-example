# Command-Line Flags

[Command-line flags](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option) are a common way to specify options for command-line programs. For example, in `wc -l` the `-l` is a command-line flag.

Go provides a `flag` package supporting basic command-line flag parsing. We’ll use this package to implement our example command-line program.

```go
import (
    "flag"
    "fmt"
)
```

Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag `word` with a default value `"foo"` and a short description. This `flag.String` function returns a string pointer (not a string value); we’ll see how to use this pointer below.

```go
wordPtr := flag.String("word", "foo", "a string")
```

This declares `numb` and `fork` flags, using a similar approach to the word flag.

```go
numbPtr := flag.Int("numb", 42, "an int")
forkPtr := flag.Bool("fork", false, "a bool")
```

It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.

```go
var svar string
flag.StringVar(&svar, "svar", "bar", "a string var")
```

Once all flags are declared, call `flag.Parse()` to execute the command-line parsing.

```go
flag.Parse()
```

Here we’ll just dump out the parsed options and any trailing positional arguments. Note that we need to dereference the pointers with e.g. \*wordPtr to get the actual option values.

```go
fmt.Println("word:", *wordPtr)
fmt.Println("numb:", *numbPtr)
fmt.Println("fork:", *forkPtr)
fmt.Println("svar:", svar)
fmt.Println("tail:", flag.Args())
```

To experiment with the command-line flags program it’s best to first compile it and then run the resulting binary directly.

```sh
go build -o command-line-flags command-line-flags/main.go
```

Try out the built program by first giving it values for all flags.

```sh
./command-line-flags/main -word=opt -numb=7 -fork -svar=flag
# word: opt
# numb: 7
# fork: true
# svar: flag
# tail: []
```

Note that if you omit flags they automatically take their default values.

```sh
./command-line-flags/main -word=opt
# word: opt
# numb: 42
# fork: false
# svar: bar
# tail: []
```

Trailing positional arguments can be provided after any flags.

```sh
./command-line-flags/main -word=opt a1 a2 a3
# word: opt
# numb: 42
# fork: false
# svar: bar
# tail: [a1 a2 a3]
```

Note that the `flag` package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).

```sh
./command-line-flags/main -word=opt a1 a2 a3 -numb=7
# word: opt
# numb: 42
# fork: false
# svar: bar
# tail: [a1 a2 a3 -numb=7]
```

If you provide a flag that wasn’t specified to the `flag` package, the program will print an error message and show the help text again.

```sh
./command-line-flags/main -wat
# flag provided but not defined: -wat
# Usage of ./command-line-flags:
# ...
```
