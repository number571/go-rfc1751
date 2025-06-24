package gorfc1751

import (
	"errors"
	"io"
)

func NewMnemonic(r io.Reader, bitSize uint64) (string, error) {
	if bitSize == 0 || bitSize%64 != 0 {
		return "", ErrBitSize
	}
	buf := make([]byte, bitSize/8)
	if _, err := r.Read(buf); err != nil {
		return "", errors.Join(ErrReader, err)
	}
	return EncodeToString(buf), nil
}
