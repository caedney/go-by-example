# Regular Expressions

Go offers built-in support for [regular expressions](https://en.wikipedia.org/wiki/Regular_expression). Here are some examples of common regexp-related tasks in Go.

This tests whether a pattern matches a string.

```go
match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
fmt.Println(match) // true
```

Above we used a string pattern directly, but for other regexp tasks you’ll need to `Compile` an optimized `Regexp` struct.

```go
r, _ := regexp.Compile("p([a-z]+)ch")
```

Many methods are available on these structs. Here’s a match test like we saw earlier.

```go
fmt.Println(r.MatchString("peach")) // true
```

This finds the match for the regexp.

```go
fmt.Println(r.FindString("peach punch")) // peach
```

This also finds the first match but returns the start and end indexes for the match instead of the matching text.

```go
fmt.Println("idx:", r.FindStringIndex("peach punch")) // idx: [0 5]
```

The `Submatch` variants include information about both the whole-pattern matches and the sub-matches within those matches. For example this will return information for both `p([a-z]+)ch` and `([a-z]+)`.

```go
fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]
```

Similarly this will return information about the indexes of matches and sub-matches.

```go
fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]
```

The `All` variants of these functions apply to all matches in the input, not just the first. For example to find all matches for a regexp.

```go
fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]
```

These `All` variants are available for the other functions we saw above as well.

```go
fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
```

Providing a non-negative integer as the second argument to these functions will limit the number of matches.

```go
fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]
```

Our examples above had string arguments and used names like `MatchString`. We can also provide `[]byte` arguments and drop `String` from the function name.

```go
fmt.Println(r.Match([]byte("peach"))) // true
```

When creating global variables with regular expressions you can use the `MustCompile` variation of `Compile`. `MustCompile` panics instead of returning an error, which makes it safer to use for global variables.

```go
r = regexp.MustCompile("p([a-z]+)ch")
fmt.Println("regexp:", r) // regexp: p([a-z]+)ch
```

The `regexp` package can also be used to replace subsets of strings with other values.

```go
fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>
```

The `Func` variant allows you to transform matched text with a given function.

```go
in := []byte("a peach")
out := r.ReplaceAllFunc(in, bytes.ToUpper)
fmt.Println(string(out)) // a PEACH
```

For a complete reference on Go regular expressions check the [regexp](https://pkg.go.dev/regexp) package docs.

```sh
go run regular-expressions/main.go
# true
# true
# peach
# idx: [0 5]
# [peach ea]
# [0 5 1 3]
# [peach punch pinch]
# all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
# [peach punch]
# true
# regexp: p([a-z]+)ch
# a <fruit>
# a PEACH
```
