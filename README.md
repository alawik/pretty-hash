# pretty-hash
Output binary buffers as a nice shortened hex string

```go
PrettyHash([32]bytes) // => 396734..68
PrettyHash([4]bytes) // => 486f7764
PrettyHash("string") // => "Not a hash"
```

Inspired by https://github.com/pfrazee/pretty-hash

## License
MIT
