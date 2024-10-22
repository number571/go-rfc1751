package gorfc1751

import "io"

func NewMnemonic(r io.Reader, bitSize uint64) string {
	if bitSize == 0 || bitSize%64 != 0 {
		panic("bitSize == 0 || bitSize%64 != 0")
	}
	buf := make([]byte, bitSize/8)
	if _, err := r.Read(buf); err != nil {
		panic("read error: " + err.Error())
	}
	return EncodeToString(buf)
}
