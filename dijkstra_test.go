package main

import (
	"testing"
	"math"
)

func Test_MakeEkiList(t *testing.T) {
	var expected []Eki
	g := &GlobalEkimei{EkimeiList:Ekimei_List}
	for _, eki := range g.EkimeiList {
		expected = append(expected, Eki{namae:eki.Kanji, saitan_kyori:math.Inf(+1)})
	}

	actual := MakeEkiList(g)

	if len(actual) != len(expected) {
		t.Errorf("got actual length %v instead of %v", len(actual), len(expected))
	}

	for i, eki := range actual {
		eki_expected := expected[i]
		if eki.namae != eki_expected.namae {
			t.Errorf("got actual %v instead of %v", eki_expected.namae, eki.namae)
		}
		if eki_expected.saitan_kyori != math.Inf(+1) {
			t.Error("saitankyori is supposed to be +inf")
		}
		if len(eki_expected.temae_list) != 0 {
			t.Error("temae list is supposed to be empty list")
		}
	}
}

func Test_Shokika(t *testing.T) {
	//var eki_list []Eki
	//kiten := "hoge"
	//
	//for _, eki := range Shokika(eki_list, kiten) {
	//	if eki.namae == kiten {
	//		if eki.saitan_kyori != 0 {
	//			t.Errorf("kiten %v's saitan kyori is %v instead of %v", kiten, eki.saitan_kyori, 0)
	//		} else if cap(eki.temae_list) != 1 || !containes(eki.temae_list, kiten) {
	//			t.Error("temae_list should only contain kiten itself.")
	//		}
	//	}
	//}

}