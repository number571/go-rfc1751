package gorfc1751

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

type testReader struct{}

func (p *testReader) Read(b []byte) (int, error) {
	r := rand.New(rand.NewSource(1)) // nolint: gosec
	return r.Read(b)
}

type testReaderWithError struct{}

func (p *testReaderWithError) Read(_ []byte) (int, error) {
	return 0, errors.New("some error") // nolint: err113
}

func ExampleNewMnemonic() {
	r := rand.New(rand.NewSource(1)) // nolint: gosec
	fmt.Println(NewMnemonic(r, 128))
	// Output: BARK TROD AMY UP LUG KNOB GAS WHEN NEWT POT KEY MEAN
}

func TestMnemonicErrors(t *testing.T) {
	t.Parallel()

	testErrorMnemonicBitSize(t, 32)
	testErrorMnemonicBitSize(t, 65)
	testErrorMnemonicBitSize(t, 224)

	testErrorMnemonicReadError(t)
}

func testErrorMnemonicBitSize(t *testing.T, bitSize uint64) {
	if _, err := NewMnemonic(&testReader{}, bitSize); err == nil {
		t.Error("nothing errors")
	}
}

func testErrorMnemonicReadError(t *testing.T) {
	if _, err := NewMnemonic(&testReaderWithError{}, 64); err == nil {
		t.Error("nothing errors")
	}
}

func TestNewMnemonic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		have uint64
		want string
	}{
		{
			name: "test-1",
			have: 64,
			want: "BARK TROD AMY UP LUG KNOB",
		},
		{
			name: "test-2",
			have: 128,
			want: "BARK TROD AMY UP LUG KNOB GAS WHEN NEWT POT KEY MEAN",
		},
		{
			name: "test-3",
			have: 192,
			want: "BARK TROD AMY UP LUG KNOB GAS WHEN NEWT POT KEY MEAN HEAL RAM ROWS JET RIM MAE",
		},
		{
			name: "test-4",
			have: 256,
			want: "BARK TROD AMY UP LUG KNOB GAS WHEN NEWT POT KEY MEAN HEAL RAM ROWS JET RIM MAE EMMA SEAR AN JUKE NOB LIN",
		},
	}
	for _, tt := range tests {
		got, _ := NewMnemonic(&testReader{}, tt.have)
		if got != tt.want {
			t.Errorf("\ntest\t= %s\nwant\t= %s\ngot\t= %s", tt.name, tt.want, got)
			continue
		}
	}
}
