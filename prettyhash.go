package main

import (
	"bytes"
	"fmt"
	//"os"
    "reflect"
    "encoding/hex"
)

type buf = bytes.Buffer

func isBuffer(b buf) bool {
    if reflect.TypeOf(b).String() == "bytes.Buffer" {
        return true
    } else {
        return false
    }
}

func PrettyHash (b buf) string {
    if isBuffer(b) {
        if b.Len() > 8 {
            buf := hex.EncodeToString(b.Bytes())
            return buf[0:6] + ".." + buf[len(buf)-2:]
        } else {
            buf := hex.EncodeToString(b.Bytes())
            return buf
        }
    } else {
        return "Not a hash"
    }
}


func main() {
	var b buf
	b.Write([]byte("Hello World!"))

    fmt.Printf(PrettyHash(b))

}
