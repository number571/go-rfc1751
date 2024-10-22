package gorfc1751

type RFC1751Error struct {
	str string
}

func (err *RFC1751Error) Error() string {
	return err.str
}

var (
	ErrIndexWordNotFound = &RFC1751Error{"index word not found"}
	ErrInvalidCheckSum   = &RFC1751Error{"invalid check sum"}
)
