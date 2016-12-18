package main

import (
	"testing"
	"math"
	"math/rand"
	"time"
	"sort"
	"fmt"
)

func Test_MakeEkiList(t *testing.T) {
	var expected []Eki
	g := &GlobalEkimei{EkimeiList:Ekimei_List}
	for _, eki := range g.EkimeiList {
		expected = append(expected, Eki{namae:eki.Kanji, saitan_kyori:math.Inf(+1)})
	}

	actual := MakeEkiList(g)

	if len(actual.eki_list) != len(expected) {
		t.Errorf("got actual length %v instead of %v", len(actual.eki_list), len(expected))
	}

	for i, eki := range actual.eki_list {
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
	sut := MakeEkiList(&GlobalEkimei{EkimeiList:Ekimei_List})
	(*sut).Shokika(expected_ekimei.Kanji)

	for i, actual_eki := range sut.eki_list {
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

func Test_Koushin1(t *testing.T) {
	ekiList := MakeEkiList(&GlobalEkimei{EkimeiList:Ekimei_List})
	ekiList.Shokika(Ekimei_List[0].Kanji) // 代々木上原を起点に。
	p := ekiList.eki_list[0]	// 代々木上原
	sut := ekiList.eki_list[1] // 代々木公園

	sut.Kousin1(p)
	if sut.saitan_kyori != 1.0 {
		t.Errorf("Got actual %v instead of %v", sut.saitan_kyori, 1.0)
	}

	sort.Strings(sut.temae_list)
	i := sort.SearchStrings(sut.temae_list, Ekimei_List[0].Kanji)
	j := sort.SearchStrings(sut.temae_list, Ekimei_List[1].Kanji)
	if i >= len(sut.temae_list) || sut.temae_list[i] != Ekimei_List[0].Kanji {
		t.Errorf("Supporsed to contain %v in temae_list", Ekimei_List[0].Kanji)
	}

	if j >= len(sut.temae_list) || sut.temae_list[j] != Ekimei_List[1].Kanji {
		t.Errorf("Supporsed to contain %v in temae_list", Ekimei_List[1].Kanji)
	}
}

func Test_koushin(t *testing.T) {
	ekiList := MakeEkiList(&GlobalEkimei{EkimeiList:Ekimei_List})
	ekiList.Shokika(ekiList.eki_list[0].namae)    // 代々木上原
	sut := &EkiList{eki_list: ekiList.eki_list[1:]} // 代々木公園
	Koushin(ekiList.eki_list[0], sut)

	// sut[0](代々木公園) should have following parameters
	expectedNamae := sut.eki_list[0].namae
	expectedSaitan := 1.0
	expectedTemaeList := []string{"代々木公園", "代々木上原"}

	if sut.eki_list[0].namae != expectedNamae {
		t.Errorf("got name %v instead of %v", sut.eki_list[0].namae, expectedNamae)
	}
	if sut.eki_list[0].saitan_kyori != expectedSaitan {
		t.Errorf("got distance %v instead of %v", sut.eki_list[0].saitan_kyori, expectedSaitan)
	}
	for i, acutalTemae := range sut.eki_list[0].temae_list {
		if acutalTemae != expectedTemaeList[i] {
			t.Errorf("got temae %v instead of %v", acutalTemae, expectedNamae[i])
			break
		}
	}
}

func TestEkiList_SaitanWoBunri(t *testing.T) {
	ekiList := MakeEkiList(&GlobalEkimei{EkimeiList:Ekimei_List})
	kiten := ekiList.eki_list[0].namae     // 代々木上原 is 起点
	ekiList.Shokika(kiten)
	sut := &EkiList{eki_list: ekiList.eki_list[1:]} // 代々木公園

	Koushin(ekiList.eki_list[0], sut)
	actualP, newV, _ := SaitanWoBunri(sut)
	fmt.Println(sut)
	fmt.Printf("Actual P: %v\n", actualP)
	fmt.Printf("New V: %v\n", newV)

	expectedNamae := "代々木公園"
	expectedSaitan := 1.0
	expectedTemaeList := []string{"代々木公園", "代々木上原"}

	if actualP.namae != expectedNamae {
		t.Errorf("got name %v instead of %v", actualP.namae, expectedNamae)
	}
	if actualP.saitan_kyori != expectedSaitan {
		t.Errorf("got distance %v instead of %v", actualP.saitan_kyori, expectedSaitan)
	}
	for i, acutalTemae := range actualP.temae_list {
		if acutalTemae != expectedTemaeList[i] {
			t.Errorf("got temae %v instead of %v", acutalTemae, expectedNamae[i])
			break
		}
	}
	if len(newV.eki_list) - 1 == len(sut.eki_list) {
		t.Errorf("got new V length %v instead of %v", len(newV.eki_list), len(sut.eki_list))
	}
}