# Strings and Runes

A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8). In other languages, strings are made of “characters”. In Go, the concept of a character is called a `rune` - it’s an integer that represents a Unicode code point. [This Go blog post](https://go.dev/blog/strings) is a good introduction to the topic.

`s` is a `string` assigned a literal value representing the word “hello” in the Thai language. Go string literals are UTF-8 encoded text.

```go
const s = "สวัสดี"
```

Since strings are equivalent to `[]byte`, this will produce the length of the raw bytes stored within.

```go
fmt.Println("Len:", len(s)) // Len: 18
```

Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.

```go
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}
fmt.Println() // e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
```

To count how many runes are in a string, we can use the `utf8` package. Note that the run-time of `RuneCountInString` depends on the size of the string, because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.

```go
fmt.Println("Rune count:", utf8.RuneCountInString(s)) // Rune count: 6
```

A `range` loop handles strings specially and decodes each `rune` along with its offset in the string.

```go
for idx, runeValue := range s {
    fmt.Printf("%#U starts at %d\n", runeValue, idx)
}

// U+0E2A 'ส' starts at 0
// U+0E27 'ว' starts at 3
// U+0E31 'ั' starts at 6
// U+0E2A 'ส' starts at 9
// U+0E14 'ด' starts at 12
// U+0E35 'ี' starts at 15
```

We can achieve the same iteration by using the `utf8.DecodeRuneInString` function explicitly.

```go
for i, w := 0, 0; i < len(s); i += w {
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at %d\n", runeValue, i)
    w = width

    examineRune(runeValue) // This demonstrates passing a rune value to a function.
}
// U+0E2A 'ส' starts at 0
// found so sua
// U+0E27 'ว' starts at 3
// U+0E31 'ั' starts at 6
// U+0E2A 'ส' starts at 9
// found so sua
// U+0E14 'ด' starts at 12
// U+0E35 'ี' starts at 15
```

Values enclosed in single quotes are _rune literals_. We can compare a `rune` value to a `rune` literal directly.

```go
func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
```

```sh
go run strings-and-runes/main.go
# Len: 18
# e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
# Rune count: 6
# U+0E2A 'ส' starts at 0
# U+0E27 'ว' starts at 3
# U+0E31 'ั' starts at 6
# U+0E2A 'ส' starts at 9
# U+0E14 'ด' starts at 12
# U+0E35 'ี' starts at 15

# Using DecodeRuneInString
# U+0E2A 'ส' starts at 0
# found so sua
# U+0E27 'ว' starts at 3
# U+0E31 'ั' starts at 6
# U+0E2A 'ส' starts at 9
# found so sua
# U+0E14 'ด' starts at 12
# U+0E35 'ี' starts at 15
```
