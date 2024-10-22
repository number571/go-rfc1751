package main

import (
	"crypto/rand"
	"flag"
	"fmt"

	gorfc1751 "github.com/number571/go-rfc1751"
)

func main() {
	sizeParam := flag.Uint64("size", 256, "entropy size in bits")
	flag.Parse()

	fmt.Println(gorfc1751.NewMnemonic(rand.Reader, *sizeParam))
}
