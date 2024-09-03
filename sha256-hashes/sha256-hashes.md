# SHA256 Hashes

[_SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2) are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature. Here’s how to compute SHA256 hashes in Go.

Go implements several hash functions in various `crypto/*` packages.

```go
import (
    "crypto/sha256"
    "fmt"
)
```

Here we start with a new hash.

```go
h := sha256.New()
```

`Write` expects bytes. If you have a string s, use `[]byte(s)` to coerce it to bytes.

```go
h.Write([]byte(s))
```

This gets the finalized hash result as a byte slice. The argument to `Sum` can be used to append to an existing byte slice: it usually isn’t needed.

```go
bs := h.Sum(nil)

fmt.Println(s) // sha256 this string
fmt.Printf("%x\n", bs) // 1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb
```

Running the program computes the hash and prints it in a human-readable hex format.

```sh
go run sha256-hashes/main.go
# sha256 this string
# 1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb
```

You can compute other hashes using a similar pattern to the one shown above. For example, to compute SHA512 hashes import `crypto/sha512` and use `sha512.New()`.

Note that if you need cryptographically secure hashes, you should carefully research [hash strength](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
