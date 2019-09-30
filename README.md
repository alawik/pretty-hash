# pretty-hash
Output binary buffers as a nice shortened hex string in Go language.

Inspired by [pfrazee/pretty-hash](https://github.com/pfrazee/pretty-hash)

## Example

```go
package main

import (
    "fmt"
    "bytes"
    "crypto/sha256"
    "github.com/alawik/pretty-hash"
)

func main() {
    sum := sha256.Sum256([]byte("some data."))

    var b bytes.Buffer
    b.Write(sum[:])

    fmt.Printf("%x\n", sum)
    fmt.Println(prettyhash.PrettyHash(b))
}
```

**Output**

```
be80822114cd5e8b0d944810ffb22e5f974536f150c7931f1bf8f2d402a22cc5
be8082..c5
```

## License
MIT
