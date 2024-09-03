# URL Parsing

URLs provide a [uniform way to locate resources](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/). Here’s how to parse URLs in Go.

We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.

```go
s := "postgres://user:pass@host.com:5432/path?k=v#f"
```

Parse the URL and ensure there are no errors.

```go
u, err := url.Parse(s)
if err != nil {
    panic(err)
}
```

Accessing the scheme is straightforward.

```go
fmt.Println(u.Scheme) // postgres
```

`User` contains all authentication info; call `Username` and `Password` on this for individual values.

```go
fmt.Println(u.User) // user:pass
fmt.Println(u.User.Username()) // user
p, _ := u.User.Password()
fmt.Println(p) // pass
```

The `Host` contains both the hostname and the port, if present. Use `SplitHostPort` to extract them.

```go
fmt.Println(u.Host) // host.com:5432
host, port, _ := net.SplitHostPort(u.Host)
fmt.Println(host) // host.com
fmt.Println(port) // 5432
```

Here we extract the path and the fragment after the #.

```go
fmt.Println(u.Path) // /path
fmt.Println(u.Fragment) // f
```

To get query params in a string of `k=v` format, use `RawQuery`. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into `[0]` if you only want the first value.

```go
fmt.Println(u.RawQuery) // k=v
m, _ := url.ParseQuery(u.RawQuery)
fmt.Println(m) // map[k:[v]]
fmt.Println(m["k"][0]) // v
```

Running our URL parsing program shows all the different pieces that we extracted.

```sh
go run url-parsing/main.go
# postgres
# user:pass
# user
# pass
# host.com:5432
# host.com
# 5432
# /path
# f
# k=v
# map[k:[v]]
# v
```
