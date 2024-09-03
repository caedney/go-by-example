# String Formatting

Go offers excellent support for string formatting in the `printf` tradition. Here are some examples of common string formatting tasks.

## `Printf`

Go offers several printing "verbs" designed to format general Go values.

```go
p := point{1, 2}
```

### Structs

The `%v` variant prints an instance of our `point` struct.

```go
fmt.Printf("struct1: %v\n", p) // struct1: {1 2}
```

The `%+v` variant will include the struct’s field names.

```go
fmt.Printf("struct2: %+v\n", p) // struct2: {x:1 y:2}
```

The `%#v` variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.

```go
fmt.Printf("struct3: %#v\n", p) // struct3: main.point{x:1, y:2}
```

### Types

The `%T` variant prints the type of a value.

```go
fmt.Printf("type: %T\n", p) // type: main.point
```

### Booleans

The `%t` variant formats booleans.

```go
fmt.Printf("bool: %t\n", true) // bool: true
```

### Integers

There are many options for formatting integers.

The `%d` variant can be used for decimal (base-10) formatting.

```go
fmt.Printf("int: %d\n", 123)  // int: 123
```

The `%b` variant can be used for binary (base-2) formatting.

```go
fmt.Printf("bin: %b\n", 14) // bin: 1110
```

The `%x` variant can be used for hexadecimal (base-16) formatting.

```go
fmt.Printf("hex: %x\n", 456) // hex: 1c8
```

The `%c` variant prints a character corresponding to the given integer.

```go
fmt.Printf("char: %c\n", 33) // char: !
```

### Floats

There are also several formatting options for floats.

The `%f` variant is used for basic decimal formatting.

```go
fmt.Printf("float1: %f\n", 78.9) // float1: 78.900000
```

The `%e` variant is used for scientific notation.

```go
fmt.Printf("float2: %e\n", 123400000.0) // 1.234000e+08
```

The `%E` variant is used for scientific notation.

```go
fmt.Printf("float3: %E\n", 123400000.0) // 1.234000E+08
```

### Strings

The `%s` variant is used for basic string printing.

```go
fmt.Printf("str1: %s\n", "\"string\"") // str1: "string"
```

The `%q` variant is used to double-quote strings as in Go source.

```go
fmt.Printf("str2: %q\n", "\"string\"") // str2: "\"string\""
```

As with integers seen earlier, the `%x` variant renders the string in base-16, with two output characters per byte of input.

```go
fmt.Printf("str3: %x\n", "hex this") // str3: 6865782074686973
```

The `%p` variant is used to print a representation of a pointer.

```go
fmt.Printf("pointer: %p\n", &p) // pointer: 0xc0000ba000
```

### Print Widths

When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.

```go
fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // width1: |    12|   345|
```

You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.

```go
fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) // width2: |  1.20|  3.45|
```

To left-justify, use the `-` flag.

```go
fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // width3: |1.20  |3.45  |
```

You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.

```go
fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // width4: |   foo|     b|
```

To left-justify use the - flag as with numbers.

```go
fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b") // width5: |foo   |b     |
```

## `Sprintf`

So far we’ve seen `Printf`, which prints the formatted string to `os.Stdout`. `Sprintf` formats and returns a string without printing it anywhere.

```go
s := fmt.Sprintf("sprintf: a %s", "string")
fmt.Println(s) // sprintf: a string
```

## `Fprintf`

You can format+print to `io.Writers` other than `os.Stdout` using `Fprintf`.

```go
fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
```

```sh
go run string-formatting/main.go
# struct1: {1 2}
# struct2: {x:1 y:2}
# struct3: main.point{x:1, y:2}
# type: main.point
# bool: true
# int: 123
# bin: 1110
# char: !
# hex: 1c8
# float1: 78.900000
# float2: 1.234000e+08
# float3: 1.234000E+08
# str1: "string"
# str2: "\"string\""
# str3: 6865782074686973
# pointer: 0xc0000ba000
# width1: |    12|   345|
# width2: |  1.20|  3.45|
# width3: |1.20  |3.45  |
# width4: |   foo|     b|
# width5: |foo   |b     |
# sprintf: a string
# io: an error
```
