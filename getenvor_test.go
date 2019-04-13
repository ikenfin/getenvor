package getenvor

import (
	"os"
	"testing"
)

type testStruct struct {
	key          string
	setenv       string
	defaultValue string
	expected     string
}

var tests = []testStruct{
	{"test-prefilled", "test-value", "", "test-value"},
	{"test-prefilled", "test-value", "default", "test-value"},
	{"test-empty", "", "default", "default"},
}

func TestGetenvOr(t *testing.T) {
	for _, s := range tests {
		if len(s.setenv) > 0 {
			os.Setenv(s.key, s.setenv)
		}

		v := GetenvOr(s.key, s.defaultValue)

		if v != s.expected {
			t.Error(
				"For", s.key,
				"expected", s.expected,
				"got", v,
			)
		}
	}
}
