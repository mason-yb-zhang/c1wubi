package wubi

import "testing"

func TestAcceptShortAndFull(t *testing.T) {
	e := Entry{Char: "你", Full: "wqiy", Short: "wq"}
	if !e.Accept("wq") || !e.Accept("WQ") || !e.Accept("wqiy") {
		t.Fatal("should accept short/full")
	}
	if e.Accept("w") || e.Accept("") {
		t.Fatal("should reject partial/empty")
	}
}

func TestSessionFlow(t *testing.T) {
	s := NewSession([]Entry{
		{Char: "一", Full: "ggll", Short: "ggg"},
		{Char: "地", Full: "fbn", Short: "f"},
	}, 0)
	s.Type('g')
	s.Type('g')
	s.Type('g')
	if !s.Submit() {
		t.Fatal("want advance")
	}
	if s.Correct != 1 || s.Index != 1 {
		t.Fatalf("c=%d i=%d", s.Correct, s.Index)
	}
	s.Type('x')
	s.Submit()
	if !s.ShowAns || s.Wrong != 1 {
		t.Fatal("want show answer")
	}
	s.Continue()
	if !s.Done() {
		t.Fatal("want done")
	}
	if s.Accuracy() != 50 {
		t.Fatalf("acc=%d", s.Accuracy())
	}
}

func TestLevelsNonEmpty(t *testing.T) {
	if len(All()) < 20 {
		t.Fatalf("table too small: %d", len(All()))
	}
}
