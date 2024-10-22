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

func DecodeString(s string) ([]byte, error) {
	sp := strings.Fields(s)
	if len(sp)%6 != 0 {
		panic("len(b)%6 != 0")
	}
	b := make([]byte, 0, len(sp)/6*8)
	for i := 0; i < len(sp); i += 6 {
		r, err := wordsToBytes(sp[i : i+6])
		if err != nil {
			return nil, err
		}
		b = append(b, r...)
	}
	return b, nil
}

func wordsToBytes(s []string) ([]byte, error) {
	if len(s) != 6 {
		panic("len(s) != 6")
	}
	result := make([]byte, 10)
	for i, w := range s {
		index, ok := wordsIndex[w]
		if !ok {
			return nil, ErrIndexWordNotFound
		}
		start := uint64(i) * bitsPerWord
		shift := (8 - ((start + bitsPerWord) % 8)) % 8
		y := index << shift
		cl := (y >> 16) & 0xFF
		cc := (y >> 8) & 0xFF
		cr := y & 0xFF
		if shift+bitsPerWord > 16 {
			result[start/8] |= byte(cl)
			result[start/8+1] |= byte(cc)
			result[start/8+2] |= byte(cr)
		} else if shift+bitsPerWord > 8 {
			result[start/8] |= byte(cc)
			result[start/8+1] |= byte(cr)
		} else {
			result[start/8] |= byte(cr)
		}
	}
	parity := uint64(0)
	for i := uint64(0); i < 64; i += 2 {
		parity += extractBits(result, i, 2)
	}
	if (parity & 3) != extractBits(result, 64, 2) {
		return nil, ErrInvalidCheckSum
	}
	return result[:8], nil
}

func bytesToWords(b []byte) []string {
	if len(b) != 8 {
		panic("len(b) != 8")
	}
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

	result := uint64(0)
	result = ((uint64(cl)<<8 | uint64(cc)) << 8) | uint64(cr)
	result = result >> (24 - (length + (start % 8)))
	result = result & (0xFFFF >> (16 - length))

	return result
}
