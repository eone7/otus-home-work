package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"распаковка с числами":              {input: "a4bc2d5e", expected: "aaaabccddddde"},
		"распаковка без чисел":              {input: "abccd", expected: "abccd"},
		"проверка на пустую строку":         {input: "", expected: ""},
		"распаковка в случае с цифрой 0":    {input: "aaa0b", expected: "aab"},
		"распаковка из одного элемента":     {input: "a", expected: "a"},
		"распаковка сожержащая спец символ": {input: "⌘4bc2d", expected: "⌘⌘⌘⌘bccd"},
	}

	t.Parallel()
	for desc, tc := range tests {
		desc, tc := desc, tc
		t.Run(desc, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "4"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
