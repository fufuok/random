package utils

import (
	"testing"
)

var (
	testString = "  Fufu 中　文\u2728->?\n*\U0001F63A   "
	testBytes  = []byte(testString)
)

func TestS2B(t *testing.T) {
	t.Parallel()
	AssertEqual(t, true, S2B("") == nil)
	AssertEqual(t, testBytes, S2B(testString))
}

func TestB2S(t *testing.T) {
	t.Parallel()
	AssertEqual(t, true, B2S(nil) == "")
	AssertEqual(t, testString, B2S(testBytes))
}

// Ref: gofiber/utils
func TestAssertEqual(t *testing.T) {
	t.Parallel()
	AssertEqual(nil, []string{}, []string{})
	AssertEqual(t, []string{}, []string{})
}
