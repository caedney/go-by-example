# Time

Go offers extensive support for times and durations; here are some examples.

```go
p := fmt.Println
```

We’ll start by getting the current time.

```go
now := time.Now()
p(now) // 2024-09-02 16:19:45.362889 +0100 BST m=+0.000052876
```

You can build a `time` struct by providing the year, month, day, etc. Times are always associated with a `Location`, i.e. time zone.

```go
then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
p(then) // 2009-11-17 20:34:58.651387237 +0000 UTC
```

You can extract the various components of the time value as expected.

```go
p(then.Year()) // 2009
p(then.Month()) // November
p(then.Day()) // 17
p(then.Hour()) // 20
p(then.Minute()) // 34
p(then.Second()) // 58
p(then.Nanosecond()) // 651387237
p(then.Location()) // UTC
```

The Monday-Sunday Weekday is also available.

```go
p(then.Weekday()) // Tuesday
```

These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.

```go
p(then.Before(now)) // true
p(then.After(now)) // false
p(then.Equal(now)) // false
```

The `Sub` methods returns a `Duration` representing the interval between two times.

```go
diff := now.Sub(then)
p(diff) // 129666h44m46.711501763s
```

We can compute the length of the duration in various units.

```go
p(diff.Hours()) // 129666.74630875049
p(diff.Minutes()) // 7.780004778525029e+06
p(diff.Seconds()) // 4.668002867115018e+08
p(diff.Nanoseconds()) // 466800286711501763
```

You can use `Add` to advance a time by a given duration, or with a - to move backwards by a duration.

```go
p(then.Add(diff)) // 2024-09-02 15:19:45.362889 +0000 UTC
p(then.Add(-diff)) // 1995-02-02 01:50:11.939885474 +0000 UTC
```

```sh
go run time/main.go
# 2024-09-02 16:19:45.362889 +0100 BST m=+0.000052876
# 2009-11-17 20:34:58.651387237 +0000 UTC
# 2009
# November
# 17
# 20
# 34
# 58
# 651387237
# UTC
# Tuesday
# true
# false
# false
# 129666h44m46.711501763s
# 129666.74630875049
# 7.780004778525029e+06
# 4.668002867115018e+08
# 466800286711501763
# 2024-09-02 15:19:45.362889 +0000 UTC
# 1995-02-02 01:50:11.939885474 +0000 UTC
```
