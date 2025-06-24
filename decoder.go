package gorfc1751

import (
	"strings"
)

func DecodeString(s string) ([]byte, error) {
	sp := strings.Fields(strings.ToUpper(s))
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
	result := make([]byte, 10)
	for i, w := range s {
		index, ok := wordsIndex[w]
		if !ok {
			return nil, ErrIndexWordNotFound
		}
		start := uint64(i) * bitsPerWord // nolint: gosec
		shift := (8 - ((start + bitsPerWord) % 8)) % 8
		y := index << shift
		cl := (y >> 16) & 0xFF
		cc := (y >> 8) & 0xFF
		cr := y & 0xFF
		switch {
		case shift+bitsPerWord > 16:
			result[start/8] |= byte(cl)
			result[start/8+1] |= byte(cc)
			result[start/8+2] |= byte(cr)
		case shift+bitsPerWord > 8:
			result[start/8] |= byte(cc)
			result[start/8+1] |= byte(cr)
		default:
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
