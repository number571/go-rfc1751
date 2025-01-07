package gorfc1751

import (
	"fmt"
	"testing"
)

func ExampleEncodeToString() {
	b := []byte{204, 172, 42, 237, 89, 16, 86, 190}
	fmt.Println(EncodeToString(b))
	// Output: RASH BUSH MILK LOOK BAD BRIM
}

func TestEncoderPanics(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Error("nothing panics")
			return
		}
	}()
	_ = EncodeToString([]byte{1, 2, 3, 4, 5, 6, 7}) // len!=8
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
