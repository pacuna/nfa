package stack

import (
	"testing"
)

func TestStack_Append(t *testing.T) {
	s := New()
	s.Append('/')
	if s.operators[len(s.operators)-1] != '/'{
		t.Error("append not working")
	}
}

func TestStack_Empty(t *testing.T) {
	s := New()
	if s.Empty() != true {
		t.Error("empty not working")
	}

	s.Append('+')
	if s.Empty() != false {
		t.Error("empty not working")
	}
}

func TestStack_Pop(t *testing.T) {
	s := New()
	s.Append('+')
	s.Append('-')
	if s.Pop() != '-'{
		t.Error("last element wasn't -")
	}
	if s.operators[len(s.operators)-1] != '+'{
		t.Error("last element wasn't popped")
	}
}

func TestStack_Top(t *testing.T) {
	s := New()
	s.Append('+')
	if s.Top() != '+'{
		t.Error("top is not +")
	}
}
