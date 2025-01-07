package gorfc1751

import (
	"strings"
)

const (
	bitsPerWord = 11
)

func EncodeToString(b []byte) string {
	if len(b)%8 != 0 {
		panic("len(b)%8 != 0")
	}
	s := make([]string, 0, len(b)/8*6)
	for i := 0; i < len(b); i += 8 {
		s = append(s, bytesToWords(b[i:i+8])...)
	}
	return strings.Join(s, " ")
}

func bytesToWords(b []byte) []string {
	parity := uint64(0)
	for i := uint64(0); i < 64; i += 2 {
		parity += extractBits(b, i, 2)
	}
	arr := make([]byte, len(b)+2)
	copy(arr, b)
	arr[len(b)] = byte((parity << 6) & 0xFF)
	words := make([]string, 0, 6)
	for i := uint64(0); i < 6; i++ {
		ebits := extractBits(arr, i*bitsPerWord, bitsPerWord)
		words = append(words, wordsList[ebits])
	}
	return words
}

func extractBits(b []byte, start, length uint64) uint64 {
	arr := make([]byte, len(b)+2)
	copy(arr, b)

	cl := arr[start/8]
	cc := arr[start/8+1]
	cr := arr[start/8+2]

	result := ((uint64(cl)<<8 | uint64(cc)) << 8) | uint64(cr)
	result >>= (24 - (length + (start % 8)))
	result &= (0xFFFF >> (16 - length))

	return result
}
