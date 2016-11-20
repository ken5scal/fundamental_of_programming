package main

import "math"

type Eki struct {
	namae        string
	saitan_kyori float64
	temae_list   []string
}

// return []Eki's temae_list(should be empty list)
// Ex: MakeEkiList(&GlobalEkimei{EkimeiList: Ekimei_List}) -> return listEki
// 		-> return Eki List with name: Kanji, saitan_kyoei inf, temae_list: empty list
func MakeEkiList(g *GlobalEkimei) []Eki {
	var eki_list []Eki

	for _, eki := range g.EkimeiList {
		eki_list = append(
			eki_list,
			Eki{
				namae:eki.Kanji,
				saitan_kyori:math.Inf(+1)})
	}
	return eki_list
}

// kiten should be Romaji!
// kiten's saitan_kyori should be 0
// kiten's temae_list should only contain kiten itself
func Shokika(eki_list []Eki, kiten string) []Eki {
	for _, eki := range eki_list {
		if eki.namae == kiten {
			eki.saitan_kyori = 0
			eki.temae_list = append(eki.temae_list, kiten)
		}
	}
	return eki_list
}