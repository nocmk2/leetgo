package leetgo

import (
	"reflect"
	"testing"
)

func Test344(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	t.Log(s)
	reverseString(s)
	t.Log(s)
	if !reflect.DeepEqual(s, []byte{'o', 'l', 'l', 'e', 'h'}) {
		t.Error()
	}
}
