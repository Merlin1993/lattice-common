package basic

import (
	"testing"
)

func TestPeekAll(t *testing.T) {
	set := NewFIFOSet()
	set.Push(1)
	set.Push(2)
	set.Push(3)
	s1 := set.PeekAll()

	set2 := NewFIFOSet()
	set2.PushAll(s1)
	s2 := set2.PeekAll()
	for _, v := range s2 {
		t.Log(v)
	}

	s1cpy := set.Copy()
	for _, v := range s1cpy.PeekAll() {
		t.Log(v)
	}
}
