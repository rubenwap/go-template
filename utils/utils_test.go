package utils

import (
	"testing"
)

func TestBuildURLwithParams(t *testing.T) {
	for _, c := range testCases {
		got := utilExample(c.input)
		if got != c.expected {
			t.Fatalf("FAIL: %q, want %q, got %q", c.input, c.expected, got)
		}
		t.Logf("PASS: %q", c.description)
	}
}
