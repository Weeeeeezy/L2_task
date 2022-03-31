package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	ValidTests := []struct {
		input    string
		expected string
	}{
		{input: "g13j4", expected: "gggggggggggggjjjj"},
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abcd", expected: "abcd"},
		{input: "", expected: ""},
		{input: `qwe\4\5`, expected: "qwe45"},
		{input: `qwe\45`, expected: "qwe44444"},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
	}

	for _, data := range ValidTests {
		ans, err := UnpackString(data.input)
		assert.Equal(t, data.expected, ans)
		assert.NoError(t, err)
	}

	InvalidTests := []struct {
		input    string
		expected string
	}{
		{input: "45", expected: ""},
		{input: `5\4`, expected: ""},
		{input: `\\\`, expected: ""},
		{input: `\\_`, expected: ""},
	}
	for _, data := range InvalidTests {
		ans, err := UnpackString(data.input)
		assert.Equal(t, data.expected, ans)
		assert.Error(t, err)
	}
}
