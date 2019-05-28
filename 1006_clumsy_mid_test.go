package leetgo

import "testing"

func Test1006(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{100, 101},
		{10, 12},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 7},
	}

	for _, c := range cases {
		t.Logf(".....")
		got := clumsy(c.in)
		if got != c.want {
			t.Errorf("clumsy(%v)==%v, want %v", c.in, got, c.want)
		}

	}
}
