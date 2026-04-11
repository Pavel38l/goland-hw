package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSpecSymbolsUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "bob\r2\t0", expected: "bob\r\r"},
		{input: "a\t3b", expected: "a\t\t\tb"},
		{input: "x\n0y", expected: "xy"},
		{input: "\n2\r2\t2", expected: "\n\n\r\r\t\t"},
		{input: "test\n1\r1\t1", expected: "test\n\r\t"},
		{input: " 3x", expected: "   x"},
		{input: "a 2b", expected: "a  b"},
		// для экранирования
		{input: "\n3\t2", expected: "\n\n\n\t\t"},
		{input: "a\n2\\23", expected: "a\n\n222"},
		{input: "test\t0\n5", expected: "test\n\n\n\n\n"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
