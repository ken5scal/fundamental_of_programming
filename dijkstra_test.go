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

func Test_Shokika(t *testing.T) {
	var eki_list []Eki
	kiten := "hoge"

	for _, eki := range Shokika(eki_list, kiten) {
		if eki.namae == kiten {
			if eki.saitan_kyori != 0 {
				t.Errorf("kiten %v's saitan kyori is %v instead of %v", kiten, eki.saitan_kyori, 0)
			} else if cap(eki.temae_list) != 1 || !containes(eki.temae_list, kiten) {
				t.Error("temae_list should only contain kiten itself.")
			}
		}
	}

}