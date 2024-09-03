# Testing And Benchmarking

Unit testing is an important part of writing principled Go programs. The `testing` package provides the tools we need to write unit tests and the `go test` command runs tests.

For the sake of demonstration, this code is in package main, but it could be any package. Testing code typically lives in the same package as the code it tests.

```go
package main
```

We’ll be testing this simple implementation of an integer minimum. Typically, the code we’re testing would be in a source file named something like `intutils.go`, and the test file for it would then be named `intutils_test.go`.

```go
func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

A test is created by writing a function with a name beginning with `Test`.

`t.Error*` will report test failures but continue executing the test.` t.Fatal*` will report test failures and stop the test immediately.

```go
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}
```

Writing tests can be repetitive, so it’s idiomatic to use a _table-driven style_, where test inputs and expected outputs are listed in a table and a single loop walks over them and performs the test logic.

```go
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }
```

t.Run enables running “subtests”, one for each table entry. These are shown separately when executing `go test -v`.

```go
    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
```

Benchmark tests typically go in `_test.go` files and are named beginning with Benchmark. The testing runner executes each benchmark function several times, increasing `b.N` on each run until it collects a precise measurement.

Typically the benchmark runs a function we’re benchmarking in a loop `b.N` times.

```go
func BenchmarkIntMin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IntMin(1, 2)
    }
}
```

Run all tests in the current project in verbose mode.

```sh
cd testing-and-benchmarking
go test -v
# === RUN   TestIntMinBasic
# --- PASS: TestIntMinBasic (0.00s)
# === RUN   TestIntMinTableDriven
# === RUN   TestIntMinTableDriven/0,1
# === RUN   TestIntMinTableDriven/1,0
# === RUN   TestIntMinTableDriven/2,-2
# === RUN   TestIntMinTableDriven/0,-1
# === RUN   TestIntMinTableDriven/-1,0
# --- PASS: TestIntMinTableDriven (0.00s)
#     --- PASS: TestIntMinTableDriven/0,1 (0.00s)
#     --- PASS: TestIntMinTableDriven/1,0 (0.00s)
#     --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
#     --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
#     --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
# PASS
# ok      go-by-example/testing-and-benchmarking  0.239s
```

Run all benchmarks in the current project. All tests are run prior to benchmarks. The `bench` flag filters benchmark function names with a regexp.

```sh
go test -bench=.
# goos: darwin
# goarch: arm64
# pkg: go-by-example/testing-and-benchmarking
# cpu: Apple M1 Pro
# BenchmarkIntMin-8       1000000000               0.3214 ns/op
# PASS
# ok      go-by-example/testing-and-benchmarking  0.586s
```
