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

func TestGlobalEkimeiList_RomajiToKanji(t *testing.T) {
	sut := &GlobalEkimeiList{EkimeiList: Ekimei_List}
	actual := sut.RomajiToKanji("myogadani")
	expected := "茗荷谷"
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}

func TestRomajiToKanjiWithUnmatchedInput(t *testing.T) {
	sut := &GlobalEkimeiList{}
	actual := sut.RomajiToKanji("IdoNotExist")
	expected := ""
	if actual != expected {
		t.Error("Actual value is not Empty")
	}
}