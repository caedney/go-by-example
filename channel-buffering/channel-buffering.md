# Channel Buffering

By default channels are _unbuffered_, meaning that they will only accept sends (`chan <-`) if there is a corresponding receive (`<- chan`) ready to receive the sent value. _Buffered_ channels accept a limited number of values without a corresponding receiver for those values.

Here we make a channel of strings buffering up to 2 values.

```go
messages := make(chan string, 2)
```

Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.

```go
messages <- "buffered"
messages <- "channel"
```

Later we can receive these two values as usual.

```go
fmt.Println(<-messages) // buffered
fmt.Println(<-messages) // channel
```

```sh
go run channel-buffering/main.go
# buffered
# channel
```
