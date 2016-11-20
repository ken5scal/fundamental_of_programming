package main

import (
	"math"
)

type Eki struct {
	namae        string
	saitan_kyori float64
	temae_list   []string
}

type EkiList struct {
	eki_list []Eki
}

// return []Eki's temae_list(should be empty list)
// Ex: MakeEkiList(&GlobalEkimei{EkimeiList: Ekimei_List}) -> return listEki
// 		-> return Eki List with name: Kanji, saitan_kyoei inf, temae_list: empty list
func MakeEkiList(g *GlobalEkimei) *EkiList {
	var eki_list []Eki

	for _, eki := range g.EkimeiList {
		eki_list = append(
			eki_list,
			Eki{
				namae:eki.Kanji,
				saitan_kyori:math.Inf(+1)})
	}
	return &EkiList{eki_list: eki_list}
}

// kiten should be Romaji!
// kiten's saitan_kyori should be 0
// kiten's temae_list should only contain kiten itself
func (e *EkiList) Shokika(kiten string) {
	for i := range e.eki_list {
		if e.eki_list[i].namae == kiten {
			e.eki_list[i].saitan_kyori = 0
			e.eki_list[i].temae_list = append(e.eki_list[i].temae_list, kiten)
		}
	}
}