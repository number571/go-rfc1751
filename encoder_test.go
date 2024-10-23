package gorfc1751

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func ExampleEncodeToString() {
	b := []byte{204, 172, 42, 237, 89, 16, 86, 190}
	fmt.Println(EncodeToString(b))
	// Output: RASH BUSH MILK LOOK BAD BRIM
}

func ExampleDecodeString() {
	s := "RASH BUSH MILK LOOK BAD BRIM"
	fmt.Println(DecodeString(s))
	// Output: [204 172 42 237 89 16 86 190] <nil>
}

func TestErrors(t *testing.T) {
	t.Parallel()

	_, err := DecodeString("TIDE ITCH SLOW REIN RULE MOX") // incorrect: MOX
	if !errors.Is(err, ErrIndexWordNotFound) {
		t.Error("error is not ErrIndexWorkNotFound")
		return
	}
}

func TestEncoderPanics(t *testing.T) {
	t.Parallel()

	testPanicEncodeToString(t, []byte{1, 2, 3, 4, 5, 6, 7})
	testPanicDecodeString(t, "TIDE ITCH SLOW REIN RULE")
}

func testPanicDecodeString(t *testing.T, s string) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("nothing panics")
			return
		}
	}()
	_, _ = DecodeString(s) // len!=6
}

func testPanicEncodeToString(t *testing.T, b []byte) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("nothing panics")
			return
		}
	}()
	_ = EncodeToString(b) // len!=8
}

func TestDecodeString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		have string
		want []byte
	}{
		{
			name: "test-1", // example from: https://datatracker.ietf.org/doc/rfc1751/
			have: "TIDE ITCH SLOW REIN RULE MOT",
			want: []byte{0xEB, 0x33, 0xF7, 0x7E, 0xE7, 0x3D, 0x40, 0x53},
		},
		{
			name: "test-1-cases", // example from: https://datatracker.ietf.org/doc/rfc1751/
			have: "tiDE ITch sLoW ReIn RULE mot",
			want: []byte{0xEB, 0x33, 0xF7, 0x7E, 0xE7, 0x3D, 0x40, 0x53},
		},
		{
			name: "test-2", // example from: https://github.com/remram44/python-rfc1751
			have: "RASH BUSH MILK LOOK BAD BRIM",
			want: []byte{204, 172, 42, 237, 89, 16, 86, 190},
		},
		{
			name: "test-3", // example from: https://github.com/vmizg/rfc1751.js
			have: "AT TIC NIBS ODD JACK ABE",
			want: []byte{4, 8, 15, 16, 23, 42, 0, 0},
		},
		{
			name: "test-4", // example from: https://github.com/arokettu/php-rfc1751
			have: "AIM HEW BLUM FED MITE WARM",
			want: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		},
		{
			name: "test-4-extend", // example from: https://github.com/arokettu/php-rfc1751
			have: "AIM HEW BLUM FED MITE WARM AIM HEW BLUM FED MITE WARM",
			want: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		},
	}
	for _, tt := range tests {
		got, err := DecodeString(tt.have)
		if err != nil {
			t.Errorf("\ntest\t= %s\nerror\t= %s", tt.name, err.Error())
			continue
		}
		if !bytes.Equal(got, tt.want) {
			t.Errorf("\ntest\t= %s\nwant\t= %X\ngot\t= %X", tt.name, tt.want, got)
			continue
		}
	}
}

func TestEncodeToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		have []byte
		want string
	}{
		{
			name: "test-1", // example from: https://datatracker.ietf.org/doc/rfc1751/
			have: []byte{0xEB, 0x33, 0xF7, 0x7E, 0xE7, 0x3D, 0x40, 0x53},
			want: "TIDE ITCH SLOW REIN RULE MOT",
		},
		{
			name: "test-2", // example from: https://github.com/remram44/python-rfc1751
			have: []byte{204, 172, 42, 237, 89, 16, 86, 190},
			want: "RASH BUSH MILK LOOK BAD BRIM",
		},
		{
			name: "test-3", // example from: https://github.com/vmizg/rfc1751.js
			have: []byte{4, 8, 15, 16, 23, 42, 0, 0},
			want: "AT TIC NIBS ODD JACK ABE",
		},
		{
			name: "test-4", // example from: https://github.com/arokettu/php-rfc1751
			have: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			want: "AIM HEW BLUM FED MITE WARM",
		},
		{
			name: "test-4-extend", // example from: https://github.com/arokettu/php-rfc1751
			have: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			want: "AIM HEW BLUM FED MITE WARM AIM HEW BLUM FED MITE WARM",
		},
	}
	for _, tt := range tests {
		got := EncodeToString(tt.have)
		if got != tt.want {
			t.Errorf("\ntest\t= %s\nwant\t= %s\ngot\t= %s", tt.name, tt.want, got)
			continue
		}
	}
}
