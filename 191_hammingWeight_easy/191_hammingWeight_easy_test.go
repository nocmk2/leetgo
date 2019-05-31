package leetgo

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cases := []struct {
		in   uint32
		want int
	}{
		{4, 1},
		{5, 2},
	}

	for _, c := range cases {
		res := hammingWeight(c.in)
		if res != c.want {
			t.Errorf("expect %v actual %v for input %v", c.want, res, c.in)
		}
	}
}
