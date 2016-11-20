package main

import (
	"testing"
	"math"
	"fmt"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(Ekimei_List))

	expected_ekimei := Ekimei_List[index]
	sut := Shokika(MakeEkiList(&GlobalEkimei{EkimeiList:Ekimei_List}), expected_ekimei.Kanji)

	for i, actual_eki := range sut {
		if i == index {
			if expected_ekimei.Kanji != actual_eki.namae {
				t.Errorf("got actual %v instead of %v", expected_ekimei.Kanji, actual_eki.namae)
			} else if 0 != actual_eki.saitan_kyori {
				t.Errorf("Saitan Kyori is not 0, but %v", actual_eki.saitan_kyori)
			} else if 1 != len(actual_eki.temae_list) {
				t.Errorf("Length of actual eki is not 1, but %v", len(actual_eki.temae_list))
			} else if expected_ekimei.Kanji != actual_eki.temae_list[0] {
				t.Errorf("got temae_list %v instead of %v", expected_ekimei.Kanji, actual_eki.temae_list[0])
			}
		}
	}
}