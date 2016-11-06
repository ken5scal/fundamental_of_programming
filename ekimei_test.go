package main

import (
	"testing"
)

func TestHyoji(t *testing.T) {
	sut := &Ekimei{Kana:"C", Kanji:"B", Shozoku:"A"}
	actual := sut.Hyoji()
	expected := "A, B(C)"
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}