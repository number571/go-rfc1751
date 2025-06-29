package gorfc1751

const (
	errPrefix = "rfc1751: "
)

type RFC1751Error struct {
	str string
}

func (err *RFC1751Error) Error() string {
	return errPrefix + err.str
}

var (
	ErrIndexWordNotFound = &RFC1751Error{"index word not found"}
	ErrInvalidCheckSum   = &RFC1751Error{"invalid check sum"}
	ErrBitSize           = &RFC1751Error{"bitSize == 0 || bitSize%64 != 0"}
	ErrReader            = &RFC1751Error{"reader"}
)
