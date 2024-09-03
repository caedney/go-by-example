# Random Numbers

Go’s `math/rand/v2` package provides [pseudorandom number](https://en.wikipedia.org/wiki/Pseudorandom_number_generator) generation.

For example, `rand.IntN` returns a random `int n, 0 <= n < 100`.

```go
fmt.Print(rand.IntN(100), ",")
fmt.Print(rand.IntN(100))
fmt.Println() // 15,59
```

`rand.Float64` returns a `float64 f`, `0.0 <= f < 1.0`.

```go
fmt.Println(rand.Float64()) // 0.22946240939978058
```

This can be used to generate random floats in other ranges, for example `5.0 <= f' < 10.0`.

```go
fmt.Print((rand.Float64()*5)+5, ",")
fmt.Print((rand.Float64() * 5) + 5)
fmt.Println() // 9.088515467641516,7.475822892893158
```

If you want a known seed, create a new `rand.Source` and pass it into the `New` constructor. `NewPCG` creates a new [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator) source that requires a seed of two `uint64` numbers.

```go
s2 := rand.NewPCG(42, 1024)
r2 := rand.New(s2)
fmt.Print(r2.IntN(100), ",")
fmt.Print(r2.IntN(100))
fmt.Println() // 94,49

s3 := rand.NewPCG(42, 1024)
r3 := rand.New(s3)
fmt.Print(r3.IntN(100), ",")
fmt.Print(r3.IntN(100))
fmt.Println() // 94,49
```

```sh
go run random-numbers/main.go
# 15,59
# 0.22946240939978058
# 9.088515467641516,7.475822892893158
# 94,49
# 94,49
```

See the [math/rand/v2](https://pkg.go.dev/math/rand/v2) package docs for references on other random quantities that Go can provide.
