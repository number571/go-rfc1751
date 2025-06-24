package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	gorfc1751 "github.com/number571/go-rfc1751"
)

func main() {
	var (
		sizeParam = flag.Uint64("size", 256, "entropy size in bits")
		convParam = flag.Bool("conv", false, "convert mnemonic / raw-hex")
		rawParam  = flag.Bool("raw", false, "print raw entropy in hex format")
	)
	flag.Parse()

	if !*convParam {
		mnemonic, err := gorfc1751.NewMnemonic(rand.Reader, *sizeParam)
		if err != nil {
			panic(err)
		}
		if !*rawParam {
			fmt.Println(mnemonic)
			return
		}
		seedBytes, _ := gorfc1751.DecodeString(mnemonic)
		fmt.Println(hex.EncodeToString(seedBytes))
		return
	}

	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(line)

	if !*rawParam {
		seedBytes, err := gorfc1751.DecodeString(input)
		if err != nil {
			panic(err)
		}
		fmt.Println(hex.EncodeToString(seedBytes))
		return
	}

	seedBytes, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	mnemonic, err := gorfc1751.NewMnemonic(
		bytes.NewReader(seedBytes),
		uint64(len(seedBytes))*8,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(mnemonic)
}
