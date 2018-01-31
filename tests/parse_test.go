package tests

import (
	"testing"
)

func TestParse(t *testing.T) {
	cases := []*TestCase{
		__testcase_1,
	}

	for _, c := range cases {
		runTestCase(t, c)
	}
}
