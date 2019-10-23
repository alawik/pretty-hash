package prettyhash

import (
	"fmt"
	"bytes"
	"testing"
	"crypto/sha256"
	"encoding/hex"
)

func TestIsBuffer(t *testing.T) {
	var b bytes.Buffer
	b.Write([]byte("Buffer Test"))
	if isBuffer(b) {
		fmt.Printf("Buffer Test: Successful\n\n")
	} else {
		fmt.Printf("Buffer Test: Fail\n\n")
	}
}

func TestPrettyHash(t *testing.T) {
	sum := sha256.Sum256([]byte("This Will Be A Hash"))
	var hash []byte = sum[:]

	var b bytes.Buffer
	b.Write([]byte(hash))

	buf := hex.EncodeToString(b.Bytes())

	fmt.Printf("Hash: %s\nPrettyHash: %s\n\n", buf, PrettyHash(b))
}
