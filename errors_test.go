package gorfc1751

import "testing"

func TestError(t *testing.T) {
	t.Parallel()

	str := "value"
	err := &RFC1751Error{str}
	if err.Error() != errPrefix+str {
		t.Error("incorrect err.Error()")
		return
	}
}
