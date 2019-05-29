package leetgo

import (
	"testing"
)

func Test443(t *testing.T) {
	s := []byte{'x'}
	i := compress(s)
	t.Log(i)
	t.Log(s)
	t.Log(string(s))
}
