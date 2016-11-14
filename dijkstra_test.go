package main

import (
	"testing"
)

func Test_MakeEkiList(t *testing.T) {
	var expected []Eki
	g := &GlobalEkimei{EkimeiList:Ekimei_List}
	expected = make([]string, cap(Ekimei_List))

	actual := MakeEkiList(g)

	if cap(actual) != cap(expected) {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}